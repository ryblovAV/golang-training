package main

import (
	pb "github.com/ryblovAV/golang-training/avg/proto"
	"io"
	"log"
)

func (*Server) CalcAvg(stream pb.AvgService_CalcAvgServer) error {

	var sum = 0.0
	var count int32 = 0
	var result = 0.0

	for {
		req, err := stream.Recv()

		if err == io.EOF {

			if count != 0 {
				result = sum / float64(count)
			}

			return stream.SendAndClose(&pb.AvgResponse{
				Result: result,
			})
		}

		if err != nil {
			log.Fatalf("error when get next value: %v\n", err)
		}

		log.Printf("receive: %f", req.Value)

		sum = sum + req.Value
		count = count + 1
	}
}
