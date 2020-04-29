package server

import (
	"context"
	pb "grpc-battle/protocol/helloWorld"
	"io"
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

// server stream
func (h *HelloWorld) List(r *pb.StreamRequest, stream pb.HelloWorld_ListServer) error {
	for n := 0; n <= 6; n++ {
		err := stream.Send(&pb.StreamResponse{
			Pt: &pb.StreamPoint{
				Name:  r.Pt.Name,
				Value: r.Pt.Value + int32(n),
			},
		})

		if err != nil {
			return err
		}
	}
	return nil
}

// client stream
func (h *HelloWorld) Record(stream pb.HelloWorld_RecordServer) error {
	// loop read data from client stream
	for {
		r, err := stream.Recv()
		if err == io.EOF {
			break

		}
		if err != nil {
			return err
		}
		log.Printf("stream.Recv pt.name: %s, pt.value: %d", r.Pt.Name, r.Pt.Value)

	}
	return stream.SendAndClose(&pb.StreamResponse{Pt: &pb.StreamPoint{
		Name:  "grpc stream Server:Record",
		Value: 1,
	}})
}

// bidirectional stream
func (h *HelloWorld) Route(stream pb.HelloWorld_RouteServer) error {
	n := 0
	for {
		err := stream.Send(&pb.StreamResponse{
			Pt: &pb.StreamPoint{
				Name:  "grpc stream client: Route",
				Value: int32(n),
			},
		})
		if err != nil {
			return err
		}

		r, err := stream.Recv()
		if err == io.EOF {
			return nil
		}

		if err != nil {
			return err
		}
		n++
		log.Printf("stream.Recv pt.name: %s, pt.value: %d", r.Pt.Name, r.Pt.Value)
	}

}
