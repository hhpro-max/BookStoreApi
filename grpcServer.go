package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"BookStoreApi/bookstorepb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	bookstorepb.UnimplementedBookstoreServer
}

func (s *Server) GetBook(ctx context.Context, req *bookstorepb.BookRequest) (*bookstorepb.BookResponse, error) {
	isbn := req.GetIsbn()

	// Dummy book data for demonstration
	bookData := map[string]struct {
		Title  string
		Author string
	}{
		"978-0143128540": {"Sapiens: A Brief History of Humankind", "Yuval Noah Harari"},
		"978-0323854322": {"Introduction to Go Programming", "John Doe"},
	}

	if book, exists := bookData[isbn]; exists {
		res := &bookstorepb.BookResponse{
			Title:  book.Title,
			Author: book.Author,
		}
		return res, nil
	}

	return nil, fmt.Errorf("book with ISBN %s not found", isbn)
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen on port 50051: %v", err)
	}

	grpcServer := grpc.NewServer()

	bookstorepb.RegisterBookstoreServer(grpcServer, &Server{})

	reflection.Register(grpcServer)

	log.Println("gRPC server is running on port 50051...")

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}
