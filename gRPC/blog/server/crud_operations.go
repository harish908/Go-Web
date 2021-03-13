package main

import (
	"../blogpb"
	"context"
	"fmt"
	"log"
)

func (*server) CreateBlog(ctx context.Context, req *blogpb.CreateBlogRequest) (*blogpb.CreateBlogResponse, error) {
	blog := req.GetBlog()
	//data := blogItem{
	//	AuthorID: blog.GetAuthorId(),
	//	Content: blog.GetContent(),
	//	Title: blog.GetTitle(),
	//}

	//Insert data
	res, err := db.Exec("INSERT INTO Blog VALUES('212', 'Content', 'title')")
	if err != nil {
		log.Fatalf("Blog data not inserted: %v", err)
	}

	lastId, err := res.LastInsertId()
	if err != nil {
		log.Fatalf("Error while getting id: %v", err)
	}
	fmt.Printf("Last Inserted data: %v", lastId)

	//update data

	response := &blogpb.CreateBlogResponse{
		Blog: &blogpb.Blog{
			Id:       string(lastId),
			AuthorId: blog.GetAuthorId(),
			Content:  blog.GetContent(),
			Title:    blog.GetTitle(),
		},
	}

	return response, nil
}
