package main

import (
	"context"
	"fmt"
	"log"

	proto "go_workspace/cleanc/mb/protoc"

	"google.golang.org/grpc"
)

func main() {

	fmt.Println("Client of Microservice B")
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()

	c := proto.NewDetailServiceClient(conn)

	req := &proto.IdRequest{
		Id: "13246170",
	}
	res, err := c.Details(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling RPC: %v", err)
	}
	log.Printf("Response from Server of B: %v", res.Result)
}
