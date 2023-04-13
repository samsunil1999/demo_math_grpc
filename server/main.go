package main

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/sam-explorex/demo_math_grpc/server/pkg"
	"google.golang.org/grpc"

	pb "github.com/sam-explorex/demo_math_grpc/proto"
)

const (
	port     = ":8080"
	httpPort = ":8081"
)

func main() {
	// go routine to run REST api server simultaneously
	go func() {
		// mux
		mux := runtime.NewServeMux()

		// register
		pb.RegisterGreetServiceHandlerServer(context.Background(), mux, &pkg.HelloServer{})

		// http server
		err := http.ListenAndServe("localhost"+httpPort, mux)
		if err != nil {
			log.Fatalf("failed to start http server: %v", err)
		}
		log.Fatalln("hettp server running at port ", httpPort)

	}()

	// listen to port 8080
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to start the server %v", err)
	}

	// creating the grpc server
	grpcServer := grpc.NewServer()

	pb.RegisterGreetServiceServer(grpcServer, &pkg.HelloServer{})
	log.Printf("server started at %v", lis.Addr())
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("failed to start %v", err)
	}
}
