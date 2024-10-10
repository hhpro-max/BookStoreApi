package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"BookStoreApi/bookstorepb"
	"google.golang.org/grpc"
)

func main() {

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	client := bookstorepb.NewBookstoreClient(conn)

	req := &bookstorepb.BookRequest{Isbn: "978-0143128540"}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.GetBook(ctx, req)
	if err != nil {
		log.Fatalf("Error while calling GetBook: %v", err)
	}
	fmt.Printf("Book Title: %s\n", res.GetTitle())
	fmt.Printf("Book Author: %s\n", res.GetAuthor())
}
