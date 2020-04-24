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
