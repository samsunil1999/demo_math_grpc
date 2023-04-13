package service

import (
	"context"

	pb "github.com/sam-explorex/demo_math_grpc/proto"
)

func (s *MathServer) Add(ctx context.Context, req *pb.Request) (*pb.Response, error) {

	res := &pb.Response{
		Result:  req.ValueA + req.ValueB,
		Message: "Add operation done succefully",
	}

	return res, nil
}

func (s *MathServer) Subtract(ctx context.Context, req *pb.Request) (*pb.Response, error) {

	res := &pb.Response{
		Result:  req.ValueA - req.ValueB,
		Message: "Subtract operation done succefully",
	}

	return res, nil
}

func (s *MathServer) Multiply(ctx context.Context, req *pb.Request) (*pb.Response, error) {

	res := &pb.Response{
		Result:  req.ValueA * req.ValueB,
		Message: "Multiply operation done succefully",
	}

	return res, nil
}

func (s *MathServer) Divide(ctx context.Context, req *pb.Request) (*pb.Response, error) {

	res := &pb.Response{
		Result:  req.ValueA / req.ValueB,
		Message: "Divide operation done succefully",
	}

	return res, nil
}
