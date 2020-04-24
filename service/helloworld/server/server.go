package server

import (
	"context"
	pb "grpc-battle/protocol/helloWorld"
	"log"
)
type HelloWorld struct {

}

func NewHelloWorld() *HelloWorld {
	log.Println("new hello world server init")
	return &HelloWorld{}
}

func (h *HelloWorld) SayHello(ctx context.Context, in *pb.SayHelloRequest) (*pb.SayHelloResponse, error) {
	log.Println("say hello pong hello world")
	return &pb.SayHelloResponse{
		Pong: "hello world",
	}, nil
}
