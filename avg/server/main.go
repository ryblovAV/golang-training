package main

import (
	pb "github.com/ryblovAV/golang-training/avg/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

var addr string = "0.0.0.0:50051"

func main() {
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("failed to listen: %v\n", err)
	}

	defer lis.Close()

	opts := []grpc.ServerOption{}
	server := grpc.NewServer(opts...)
	pb.RegisterAvgServiceServer(server, &Server{})

	if err := server.Serve(lis); err != nil {
		log.Fatalf("serve error: %v\n", err)
	}
}
