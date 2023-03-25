package main

import (
	pb "github.com/ryblovAV/golang-training/primes/proto"
	"log"
)

func (*Server) Primes(req *pb.PrimesRequest, s pb.PrimesService_PrimesServer) error {
	log.Printf("server: receive request %d", (*req).StartValue)
	calcPrimes(req.StartValue, func(i int32) {
		err := s.Send(&pb.PrimesResponse{Result: i})
		if err != nil {
			log.Fatalf("can't process new value: %s", err)
		}
	})
	return nil
}
