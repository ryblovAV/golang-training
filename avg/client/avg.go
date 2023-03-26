package main

import (
	"context"
	pb "github.com/ryblovAV/golang-training/avg/proto"
	"log"
	"time"
)

func calcAvg(client pb.AvgServiceClient) {

	stream, err := client.CalcAvg(context.Background())

	if err != nil {
		log.Fatalf("Error while calc avg, error: %v", err)
	}

	for i := 0; i < 10; i++ {
		req := pb.AvgRequest{Value: float64(i)}
		stream.Send(&req)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error while receiving from calc avg: %v", err)
	}

	log.Printf("calc avg: %v", res.Result)
}
