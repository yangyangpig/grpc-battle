package helloWorld

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	pb "grpc-battle/protocol/helloWorld"
	"testing"
)

func TestClient_SayHello(t *testing.T) {
	c, err := newClient(grpc.WithInsecure(),grpc.WithBlock())
	if err != nil {
		t.Fatalf("%v", err)
	}
	resp, err := c.SayHello(context.Background(), &pb.SayHelloRequest{Ping: "hello world"})
	if err != nil {
		t.Errorf("say hello world error (%+v)", err)
		return
	}
	fmt.Println(resp)

}

func TestClient_ServerStreamList(t *testing.T) {
	c, err := newClient(grpc.WithInsecure(),grpc.WithBlock())
	if err != nil {
		t.Fatalf("%v", err)
	}
	err = c.ServerStreamList(context.Background(), &pb.StreamRequest{Pt:&pb.StreamPoint{Name:"grpc client request please repose stream", Value:0}})
	if err != nil {
		t.Errorf("server stream list error (%+v)", err)
		return
	}
}

func TestClient_ClientStreamRecord(t *testing.T) {
	c, err := newClient(grpc.WithInsecure(),grpc.WithBlock())
	if err != nil {
		t.Fatalf("%v", err)
	}

	err = c.ClientStreamRecord(context.Background())
	if err != nil {
		t.Errorf("client stream list error (%+v)", err)
		return
	}
}

func TestClient_BidirectionalRoute(t *testing.T) {
	//c, err := newClient(grpc.WithInsecure(),grpc.WithBlock())
	//if err != nil {
	//	t.Fatalf("%v", err)
	//}
	ts := fmt.Sprintf("%.8x", 1232)
	fmt.Println(ts)

	return
	//err = c.BidirectionalRoute(context.Background())
	//if err != nil {
	//	t.Errorf("bidirection stream list error (%+v)", err)
	//	return
	//}
}


