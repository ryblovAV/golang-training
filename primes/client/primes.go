package main

import (
	"context"
	pb "github.com/ryblovAV/golang-training/primes/proto"
	"io"
	"log"
)

func doPrimes(startValue int32, client pb.PrimesServiceClient) {
	log.Printf("start doPrimes")

	req := &pb.PrimesRequest{
		StartValue: startValue,
	}

	stream, err := client.Primes(context.Background(), req)

	if err != nil {
		log.Fatalf("Error when call Primes server: %s\n", err)
	}

	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			log.Printf("End stream\n")
			break
		}

		if err != nil {
			log.Fatalf("Error when read primes stream: %s\n", err)
		}

		log.Printf("primes: %s\n", msg.Result)
	}
}
