package main

import (
	pb "../gen/proto"
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
)

type testApiServer struct {
	pb.UnimplementedTestApiServer
}

func (s *testApiServer)  GetUser(ctx context.Context,req *pb.UserRequest) (*pb.UserResponse, error) {
	return &pb.UserResponse{}, nil
}


func (s *testApiServer)  Echo(ctx context.Context,req *pb.ResponseRequest) (*pb.ResponseRequest, error) {
 return req, nil
}


func main() {

	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

  grpcServer :=	grpc.NewServer()

  pb.RegisterTestApiServer(grpcServer, &testApiServer{})

  err = grpcServer.Serve(listener)
	if err != nil {
		log.Println(err)
	}
}
