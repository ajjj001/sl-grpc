package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/ajjj001/sl-grpc/gen/proto"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	client := pb.NewTestApiClient(conn)

	resp, err := client.Echo(context.Background(), &pb.ResponseRequest{Msg: "Hello World"})
	if err != nil {
		log.Println(err)
	}

	fmt.Println(resp)
	fmt.Println(resp.Msg)
}
