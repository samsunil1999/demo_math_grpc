package main

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"

	"github.com/sam-explorex/demo_math_grpc/database"
	pb "github.com/sam-explorex/demo_math_grpc/proto"
	"github.com/sam-explorex/demo_math_grpc/server/service"
)

const (
	port     = ":8080"
	httpPort = ":8081"
)

func main() {
	database.ConnectDatabase()
	database.MigrateDatabase()

	// go routine to run REST api server simultaneously
	go func() {
		// mux
		mux := runtime.NewServeMux()

		// register
		pb.RegisterMathServiceHandlerServer(context.Background(), mux, &service.MathServer{})

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

	pb.RegisterMathServiceServer(grpcServer, &service.MathServer{})
	log.Printf("server started at %v", lis.Addr())
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("failed to start %v", err)
	}
}
