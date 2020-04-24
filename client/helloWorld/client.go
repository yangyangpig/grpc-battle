package helloWorld

import (
	"context"
	"errors"
	"google.golang.org/grpc"
	pb "grpc-battle/protocol/helloWorld"
	"log"
	"time"
)

const target  = "127.0.0.1:8080"
type Client struct {
	sub pb.HelloWorldClient
}

func NewClient(dopts ...grpc.DialOption) (*Client, error) {
	return NewClient(dopts...)
}

func newClient(dopts ...grpc.DialOption) (*Client, error) {
	c, err := grpc.Dial(target, dopts...)
	if err != nil {
		log.Printf("dial remove server fail %v", err)
		return nil, errors.New("dial remove server fail")
	}
	return &Client{
		sub:pb.NewHelloWorldClient(c),
	}, nil
}

func (c *Client) SayHello(ctx context.Context, in *pb.SayHelloRequest, opts ...grpc.CallOption) (*pb.SayHelloResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 3)
	defer cancel()

	res, err := c.sub.SayHello(ctx, in, opts...)

	return res, err
}
