package main

import (
	"../greetpb"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"log"
	"math"
	"net"
	"strconv"
	"time"
)

type server struct{}

func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	log.Printf("Greet function invoked: %v", req)
	firstName := req.GetGreet().GetFirstName()
	result := "Hello " + firstName
	res := &greetpb.GreetResponse{
		Result: result,
	}
	return res, nil
}

func (*server) GreetManyTimes(req *greetpb.GreetManyTimesRequest, stream greetpb.GreetService_GreetManyTimesServer) error {
	log.Printf("GreetManyTimes function invoked: %v", req)
	firstName := req.GetGreet().GetFirstName()
	for i := 0; i < 10; i++ {
		result := "Hello " + firstName + " Number:" + strconv.Itoa(i)
		res := &greetpb.GreetManyTimesResponse{
			Result: result,
		}
		stream.Send(res)
		time.Sleep(500 * time.Millisecond)
	}
	return nil
}

func (*server) LongGreet(stream greetpb.GreetService_LongGreetServer) error {
	log.Printf("LongGreet function invoked")
	result := ""
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&greetpb.LongGreetResponse{
				Result: result,
			})
		}
		if err != nil {
			log.Fatalf("LongGreet throws error: %v", err)
		}
		firstName := req.GetGreet().GetFirstName()
		result += " Hello: " + firstName
	}
	return nil
}

func (*server) GreetEveryone(stream greetpb.GreetService_GreetEveryoneServer) error {
	log.Printf("GreetEveryone function invoked")
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("LongGreet throws error: %v", err)
			return err
		}
		firstName := req.GetGreet().GetFirstName()
		result := " Hello:" + firstName
		sendErr := stream.Send(&greetpb.GreetEveryoneResponse{
			Result: result,
		})
		if sendErr != nil {
			log.Fatalf("Error while sending data: %v", sendErr)
			return sendErr
		}
	}
}

func (*server) SquareRoot(ctx context.Context, req *greetpb.SquareRootRequest) (*greetpb.SquareRootResponse, error) {
	log.Printf("SquareRoot function invoked: %v", req)
	number := req.GetNumber()
	if number < 0 {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Invalid number received: %v", number))
	}
	return &greetpb.SquareRootResponse{
		NumberRoot: math.Sqrt(float64(number)),
	}, nil
}

func (*server) GreetWithDeadLine(ctx context.Context, req *greetpb.GreetWithDeadLineRequest) (*greetpb.GreetWithDeadLineResponse, error) {
	log.Printf("GreetWithDeadLine function invoked: %v", req)

	for i := 0; i < 3; i++ {
		if ctx.Err() == context.Canceled {
			//Client cancelled request
			fmt.Println("Client cancelled request")
			return nil, status.Error(codes.Canceled, "Client cancelled request")
		}
		time.Sleep(1 * time.Second)
	}

	firstName := req.GetGreet().GetFirstName()
	result := "Hello " + firstName
	res := &greetpb.GreetWithDeadLineResponse{
		Result: result,
	}
	return res, nil
}

func main() {
	lis, err := net.Listen("tcp", "127.0.0.1:80")
	if err != nil {
		log.Fatalf("Server error: %v", err)
	}
	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})

	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("Server error: %v", err)
	}

}
