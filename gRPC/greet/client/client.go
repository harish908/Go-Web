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
	"time"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:80", grpc.WithInsecure())
	if err != nil {
		fmt.Println("Server error", err)
	}

	c := greetpb.NewGreetServiceClient(conn)
	//unary(c)
	//serverStreaming(c)
	//clientStreaming(c)
	//biDirectionalStreaming(c)
	//unaryErrorHandling(c)
	unaryDeadline(c, 1*time.Second)
}

func unary(c greetpb.GreetServiceClient) {
	log.Print("Starting unary GRPC")
	req := greetpb.GreetRequest{
		Greet: &greetpb.Greeting{
			FirstName: "Harish",
			LastName:  "M",
		},
	}
	res, err := c.Greet(context.Background(), &req)
	if err != nil {
		log.Fatalf("Error for unary GRPC: %v", err)
	}
	log.Printf("unary result: %v", res.Result)
}

func serverStreaming(c greetpb.GreetServiceClient) {
	log.Print("Starting Server streaming GRPC")
	req := greetpb.GreetManyTimesRequest{
		Greet: &greetpb.Greeting{
			FirstName: "Harish",
			LastName:  "M",
		},
	}

	res, err := c.GreetManyTimes(context.Background(), &req)
	if err != nil {
		log.Fatalf("Error for Server streaming GRPC: %v", err)
	}
	for {
		msg, err := res.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error while reading stream: %v", err)
		}
		log.Printf("Response from server: %v", msg.GetResult())
	}
}

func clientStreaming(c greetpb.GreetServiceClient) {
	log.Print("Starting Client streaming GRPC")

	requests := []*greetpb.LongGreetRequest{
		{
			Greet: &greetpb.Greeting{
				FirstName: "Harish",
			},
		},
		{
			Greet: &greetpb.Greeting{
				FirstName: "Sam",
			},
		},
	}

	stream, err := c.LongGreet(context.Background())
	if err != nil {
		log.Fatalf("Error while connecting LongGreet: %v", err)
	}
	for _, req := range requests {
		log.Printf("Sending request: %v", req)
		stream.Send(req)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while returning LongGreet: %v", err)
	}
	log.Printf("LongGreet Response: %v", res.GetResult())
}

func biDirectionalStreaming(c greetpb.GreetServiceClient) {
	log.Print("Starting biDirectional streaming GRPC")
	requests := []*greetpb.GreetEveryoneRequest{
		{
			Greet: &greetpb.Greeting{
				FirstName: "Harish",
			},
		},
		{
			Greet: &greetpb.Greeting{
				FirstName: "Sam",
			},
		},
	}

	stream, err := c.GreetEveryone(context.Background())
	if err != nil {
		log.Fatalf("Error while connecting GreetEveryone: %v", err)
	}

	wait := make(chan struct{})
	//Sending messages
	go func() {
		for _, req := range requests {
			log.Printf("Sending request: %v", req)
			stream.Send(req)
			time.Sleep(500 * time.Millisecond)
		}
		stream.CloseSend()
	}()
	//Receiving messages
	go func() {
		res, err := stream.Recv()
		if err == io.EOF {
			close(wait)
		}
		if err != nil {
			log.Fatalf("Error while receiving message: %v", err)
		}
		log.Printf("Received data: %v", res.GetResult())
		time.Sleep(500 * time.Millisecond) //Since we use goroutines we must wait to display output
	}()

	<-wait
}

func unaryErrorHandling(c greetpb.GreetServiceClient) {
	log.Print("Starting unary Error handling GRPC")
	//Correct Call
	errorCall(c, 12)
	//Wrong Call
	errorCall(c, -2)
}

func errorCall(c greetpb.GreetServiceClient, n int32) {
	req := &greetpb.SquareRootRequest{
		Number: n,
	}
	res, err := c.SquareRoot(context.Background(), req)
	if err != nil {
		resError, ok := status.FromError(err)
		if ok {
			//actual error from grpc(user error)
			log.Printf("Error message from server: %v\n", resError.Message())
			log.Printf("Error Code from server: %v\n", resError.Code())

			if resError.Code() == codes.InvalidArgument {
				log.Fatalf("Sent Negative Number")
				return
			}
		} else {
			log.Fatalf("Big error calling: %v", err)
			return
		}
	}
	log.Printf("Result of Square root %v : %v\n", n, res.GetNumberRoot())
}

func unaryDeadline(c greetpb.GreetServiceClient, seconds time.Duration) {
	log.Print("Starting unaryDeadline GRPC")
	req := greetpb.GreetWithDeadLineRequest{
		Greet: &greetpb.Greeting{
			FirstName: "Harish",
			LastName:  "M",
		},
	}
	ctx, cancel := context.WithTimeout(context.Background(), seconds)
	defer cancel()

	res, err := c.GreetWithDeadLine(ctx, &req)
	if err != nil {
		statusErr, ok := status.FromError(err)
		if ok {
			if statusErr.Code() == codes.DeadlineExceeded {
				fmt.Println("Deadline limit!")
			} else {
				fmt.Printf("unexcepted error: %v", statusErr)
			}
		} else {
			log.Fatalf("Error for unary GRPC: %v", err)
		}
		return
	}
	log.Printf("unary result: %v", res.Result)
}
