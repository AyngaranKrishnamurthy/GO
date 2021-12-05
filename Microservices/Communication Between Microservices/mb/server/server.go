package main

import (
	"context"
	"fmt"
	"log"
	"net"

	proto "go_workspace/cleanc/mb/protoc"

	"google.golang.org/grpc"
)

type server struct{}

var aid string
var val string
var cred string

func (*server) Details(ctx context.Context, req *proto.IdRequest) (*proto.IdResponse, error) {
	fmt.Println("Server 2 invoked successfully")
	val := req.GetId()
	if val == "LTP1" {
		cred = "||| Description : Founder of Solutionz  | Machine Learning Enthusiast | Risk Management Analyst | Motivational Speaker "
	}
	result := cred
	res := &proto.IdResponse{
		Result: result,
	}
	return res, nil
}

func main() {
	fmt.Println("Server of Microservice B")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	proto.RegisterDetailServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
