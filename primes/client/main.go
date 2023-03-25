package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"

	pb "github.com/ryblovAV/golang-training/primes/proto"
)

var addr string = "0.0.0.0:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Can't connect %s", addr)
	}

	defer conn.Close()

	client := pb.NewPrimesServiceClient(conn)

	doPrimes(120, client)
}
