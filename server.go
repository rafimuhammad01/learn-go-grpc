package main

import (
	"fmt"
	"github.com/rafimuhammad01/learn-go-grpc/server"

	"go.elastic.co/apm/module/apmgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"

	pb "github.com/rafimuhammad01/learn-go-grpc/proto/currency"
)

func Serve() {
	port := "9001"
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v \n", err)
	}
	//opts := recovery.Options()
	s := grpc.NewServer(grpc.UnaryInterceptor(
		apmgrpc.NewUnaryServerInterceptor(apmgrpc.WithRecovery()),
	))
	pb.RegisterCurrencyServer(s, server.NewGrpcServer())
	fmt.Printf("⇨ grpc server started on :%s \n", port)
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func main()  {
	Serve()
}