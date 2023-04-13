package service

import (
	"context"
	"log"

	pb "github.com/sam-explorex/demo_math_grpc/proto"
)

func (s *HelloServer) SayHello(ctx context.Context, req *pb.NoParam) (*pb.HelloResponse, error) {
	log.Println("running the SayHello Service")
	return &pb.HelloResponse{
		Message: "Hello",
	}, nil
}
