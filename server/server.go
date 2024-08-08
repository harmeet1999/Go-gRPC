package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "grpc/proto"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedHelloServiceServer
}

func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HellResponse, error) {
	fmt.Println("Received: " + req.Name)
	return &pb.HellResponse{Message: "Hello, " + req.Name}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer() //Engine
	pb.RegisterHelloServiceServer(grpcServer, &server{})

	log.Println("gRPC server is running on port :9000")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}
