package main

import (
	"context"
	"fmt"
	"log"
	"net"

	L "go_workspace/cleanc/data"
	ma "go_workspace/cleanc/ma/proto"
	mb "go_workspace/cleanc/mb/protoc"

	"google.golang.org/grpc"
)

type server struct{}

var id string
var aid string
var val string
var val1 string

func (*server) Name(ctx context.Context, req *ma.Request) (*ma.Response, error) {
	id = "000"
	firstName := req.GetName()
	if firstName == "excaliber" {
		id = "1234"
	}
	val1 = clientb()
	result := L.Data(id) + val1
	res := &ma.Response{
		Result: result,
	}
	return res, nil
}

func main() {
	fmt.Println("Server of Microservice A")

	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	ma.RegisterAddServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func clientb() string {
	fmt.Println("Client of Microservice B")
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()

	c := mb.NewDetailServiceClient(conn)
	if id == "1324" {
		aid = "LTP1"
	}
	req := &mb.IdRequest{
		Id: aid,
	}
	res, err := c.Details(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling RPC: %v", err)
	}

	log.Printf("Response from Server of B: %v", res.Result)
	val = res.Result
	return val
}
