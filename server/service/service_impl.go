package service

import (
	"context"

	pb "github.com/sam-explorex/demo_math_grpc/proto"
)

func (s *MathServer) Add(ctx context.Context, req *pb.Request) (*pb.Response, error) {

	res := &pb.Response{
		Result:  float32(req.ValueA) + float32(req.ValueB),
		Message: "Add operation done succefully",
	}

	return res, nil
}

func (s *MathServer) Subtract(ctx context.Context, req *pb.Request) (*pb.Response, error) {

	res := &pb.Response{
		Result:  float32(req.ValueA) - float32(req.ValueB),
		Message: "Subtract operation done succefully",
	}

	return res, nil
}

func (s *MathServer) Multiply(ctx context.Context, req *pb.Request) (*pb.Response, error) {

	res := &pb.Response{
		Result:  float32(req.ValueA) * float32(req.ValueB),
		Message: "Multiply operation done succefully",
	}

	return res, nil
}

func (s *MathServer) Divide(ctx context.Context, req *pb.Request) (*pb.Response, error) {

	if req.ValueB == 0 {
		return &pb.Response{
			Result:  0,
			Message: "unsuccefull, value_b cannot be 0",
		}, nil
	}

	res := &pb.Response{
		Result:  float32(req.ValueA) / float32(req.ValueB),
		Message: "Divide operation done succefully",
	}

	return res, nil
}
