package main

import (
	pb "github.com/ryblovAV/golang-training/primes/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

var addr string = "0.0.0.0:50051"

func main() {
	log.Printf("starting sever: addr=%s\n", addr)

	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("failed to listen: %s", err)
	}

	defer lis.Close()

	opts := []grpc.ServerOption{}
	s := grpc.NewServer(opts...)
	pb.RegisterPrimesServiceServer(s, &Server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("server error: %s\n", err)
	}
}
