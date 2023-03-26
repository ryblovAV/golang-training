package main

import (
	pb "github.com/ryblovAV/golang-training/avg/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

var addr string = "0.0.0.0:50051"

func main() {

	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	defer conn.Close()

	if err != nil {
		log.Fatalf("Can't connect %s", addr)
	}

	client := pb.NewAvgServiceClient(conn)
	calcAvg(client)
}
