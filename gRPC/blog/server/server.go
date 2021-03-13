package main

import (
	"../blogpb"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
)

type server struct{}

type blogItem struct {
	ID       string `json:"id"`
	AuthorID string `json:"author_id"`
	Content  string `json:"content"`
	Title    string `json:"title"`
}

var db *sql.DB

//ReadBlog(context.Context, *ReadBlogRequest) (*ReadBlogResponse, error)
//UpdateBlog(context.Context, *UpdateBlogRequest) (*UpdateBlogResponse, error)
//DeleteBlog(context.Context, *DeleteBlogRequest) (*DeleteBlogResponse, error)
//ListBlog(*ListBlogRequest, BlogService_ListBlogServer) error

func main() {
	// In case of crash, get the filename and line number
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	//Mysql connection
	fmt.Println("Starting mysql server")
	var dbErr error
	db, dbErr = sql.Open("mysql", "root:example@tcp(127.0.0.1:3306)/TestDB")
	if dbErr != nil {
		log.Fatalf("Error while starting mysql: %v", dbErr)
	}

	//Go listener
	fmt.Println("Starting blog server")
	lis, connErr := net.Listen("tcp", "127.0.0.1:8050")
	if connErr != nil {
		log.Fatalf("Failed to Listen: %v", connErr)
	}

	//Grpc server
	opts := []grpc.ServerOption{}
	s := grpc.NewServer(opts...)
	blogpb.RegisterBlogServiceServer(s, &server{})
	go func() {
		fmt.Println("Starting server")
		err := s.Serve(lis)
		if err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()

	//Wait for ctrl+c to exit
	exitServe := make(chan os.Signal, 1)
	signal.Notify(exitServe, os.Interrupt)
	//block till signal is received
	<-exitServe

	println("Stopping the Serve")
	s.Stop()
	println("Closing the Listener")
	lis.Close()
	println("Closing mysql connection")
	db.Close()
}
