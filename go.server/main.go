package main

import (
	"log"
	"net"

	pb "github.com/humbertodias/grpc.demo/proto"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

type server struct{}

func (s *server) ReverseString(ctx context.Context, in *pb.ReverseRequest) (*pb.ReverseReply, error) {
	// your reverse string implementation here
	return &pb.ReverseReply{Reversed: reversed}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterReverseServiceServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}
