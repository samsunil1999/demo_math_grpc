package pkg

import pb "github.com/sam-explorex/demo_math_grpc/proto"

type HelloServer struct {
	pb.UnimplementedGreetServiceServer
}
