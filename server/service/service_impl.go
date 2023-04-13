package service

import (
	"context"

	"github.com/sam-explorex/demo_math_grpc/entities"
	pb "github.com/sam-explorex/demo_math_grpc/proto"
	"github.com/sam-explorex/demo_math_grpc/repository"
)

func (s *MathServer) Add(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	result := float32(req.ValueA) + float32(req.ValueB)

	log := entities.Logs{
		ValueA:    req.ValueA,
		ValueB:    req.ValueB,
		Operation: "add",
		Result:    float64(result),
	}

	_, err := repository.CreateLogs(log)
	if err != nil {
		return &pb.Response{}, err
	}

	res := &pb.Response{
		Result:  result,
		Message: "Add operation done succefully",
	}

	return res, nil
}

func (s *MathServer) Subtract(ctx context.Context, req *pb.Request) (*pb.Response, error) {

	result := float32(req.ValueA) - float32(req.ValueB)

	log := entities.Logs{
		ValueA:    req.ValueA,
		ValueB:    req.ValueB,
		Operation: "subtract",
		Result:    float64(result),
	}

	_, err := repository.CreateLogs(log)
	if err != nil {
		return &pb.Response{}, err
	}

	res := &pb.Response{
		Result:  result,
		Message: "Subtract operation done succefully",
	}

	return res, nil
}

func (s *MathServer) Multiply(ctx context.Context, req *pb.Request) (*pb.Response, error) {

	result := float32(req.ValueA) * float32(req.ValueB)

	log := entities.Logs{
		ValueA:    req.ValueA,
		ValueB:    req.ValueB,
		Operation: "multiply",
		Result:    float64(result),
	}

	_, err := repository.CreateLogs(log)
	if err != nil {
		return &pb.Response{}, err
	}

	res := &pb.Response{
		Result:  result,
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

	result := float32(req.ValueA) / float32(req.ValueB)

	log := entities.Logs{
		ValueA:    req.ValueA,
		ValueB:    req.ValueB,
		Operation: "divide",
		Result:    float64(result),
	}

	_, err := repository.CreateLogs(log)
	if err != nil {
		return &pb.Response{}, err
	}

	res := &pb.Response{
		Result:  result,
		Message: "Divide operation done succefully",
	}

	return res, nil
}
