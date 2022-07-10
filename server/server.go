package main

import (
	"context"
	"log"
	"net"

	pb "github.com/ajjj001/sl-grpc/gen/proto"

	"google.golang.org/grpc"
)

type testApiServer struct {
	pb.UnimplementedTestApiServer
}

func (s *testApiServer) GetBook(ctx context.Context, req *pb.BookRequest) (*pb.BookResponse, error) {
	return &pb.BookResponse{}, nil
}

func (s *testApiServer) Echo(ctx context.Context, req *pb.ResponseRequest) (*pb.ResponseRequest, error) {
	return req, nil
}

func main() {

	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterTestApiServer(grpcServer, &testApiServer{})

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Println(err)
	}
}
