package server

import (
	"context"
	pb "github.com/rafimuhammad01/learn-go-grpc/proto/currency"
)

type GrpcServer struct {
	pb.UnimplementedCurrencyServer
}


func NewGrpcServer() *GrpcServer {
	return &GrpcServer{}

}


func (s *GrpcServer) GetCurrency(ctx context.Context,req *pb.GetCurrencyRequest) (*pb.GetCurrencyResponse, error)  {
	return &pb.GetCurrencyResponse{
		Rate : 0.5,
	}, nil
}
