package main

import (
	"../blogpb"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
)

func main() {
	fmt.Println("Blog Client")
	conn, err := grpc.Dial("127.0.0.1:8050", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Failed to dial: %v", err)
	}
	defer conn.Close()

	c := blogpb.NewBlogServiceClient(conn)

	// Create Blog
	fmt.Println("Creating a blog")
	blog := &blogpb.Blog{
		AuthorId: "Sooraj",
		Title:    "My Third Blog",
		Content:  "Third Blog Contents",
	}

	createdBlog, err := c.CreateBlog(context.Background(), &blogpb.CreateBlogRequest{
		Blog: blog,
	})
	if err != nil {
		log.Fatalf("Unexpected Error: %v", err)
	}
	fmt.Printf("Created blog: %v", createdBlog)
}
