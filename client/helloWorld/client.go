package helloWorld

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc"
	pb "grpc-battle/protocol/helloWorld"
	"io"
	"log"
	"time"
)

const target = "127.0.0.1:8080"

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
		sub: pb.NewHelloWorldClient(c),
	}, nil
}

func (c *Client) SayHello(ctx context.Context, in *pb.SayHelloRequest, opts ...grpc.CallOption) (*pb.SayHelloResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	res, err := c.sub.SayHello(ctx, in, opts...)

	return res, err
}

func (c *Client) ServerStreamList(ctx context.Context, in *pb.StreamRequest) error {
	serverStream, err := c.sub.List(ctx, in)
	if err != nil {
		return err
	}
	// loop read data from server stream
	for {
		resp, err := serverStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		log.Printf("resp: pj.name: %s, pt.value: %d", resp.Pt.Name, resp.Pt.Value)
	}
	return nil
}

func (c *Client) ClientStreamRecord(ctx context.Context) error {
	stream, err := c.sub.Record(ctx)
	if err != nil {
		return err
	}

	for n := 0; n <= 10; n++ {
		err := stream.Send(&pb.StreamRequest{
			Pt: &pb.StreamPoint{
				Name:  fmt.Sprintf("v-%d", n),
				Value: int32(n),
			},
		})
		if err != nil {
			return err
		}
	}

	resp, err := stream.CloseAndRecv()
	if err != nil {
		return nil
	}
	log.Printf("resp: pj.name: %s, pt.value: %d", resp.Pt.Name, resp.Pt.Value)
	return nil

}

func (c *Client)BidirectionalRoute(ctx context.Context) error {
	streamClient, err := c.sub.Route(ctx)
	if err != nil {
		return err
	}

	for n:=0; n <= 10; n++ {
		err = streamClient.Send(&pb.StreamRequest{Pt:&pb.StreamPoint{
			Name:                 fmt.Sprintf("v-%d", n),
			Value:                int32(n),
		}})

		if err != nil {
			return err
		}

		// wait the server response
		resp, err := streamClient.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		log.Printf("resp: pj.name: %s, pt.value: %d", resp.Pt.Name, resp.Pt.Value)
	}
	streamClient.CloseSend()
	return nil
}


