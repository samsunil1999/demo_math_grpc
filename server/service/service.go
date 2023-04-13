package service

import pb "github.com/sam-explorex/demo_math_grpc/proto"

type MathServer struct {
	pb.UnimplementedMathServiceServer
}
