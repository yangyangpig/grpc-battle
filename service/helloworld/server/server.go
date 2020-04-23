package server

import (
	"context"
	pb "grpc-battle/protocol/helloWorld"
)
type HelloWorld struct {

}

func NewHelloWorld() *HelloWorld {
	return &HelloWorld{}
}

func (h *HelloWorld) SayHello(ctx context.Context, in *pb.SayHelloRequest) (*pb.SayHelloResponse, error) {
	return &pb.SayHelloResponse{
		Pong: "hello world",
	}, nil
}
