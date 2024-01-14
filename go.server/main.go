package main

import (
	"context"
	"fmt"
	"go.server/v1/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

// Ensure that reverseServer implements the proto.ReverseServiceServer interface
var _ proto.ReverseServiceServer = &reverseServer{}

type reverseServer struct {
	proto.UnimplementedReverseServiceServer
}

func (s *reverseServer) ReverseString(ctx context.Context, req *proto.ReverseRequest) (*proto.ReverseReply, error) {
	reversed := reverseString(req.GetData())
	return &proto.ReverseReply{Reversed: reversed}, nil
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// inicia servidor gRPC
	var opts []grpc.ServerOption
	srv := grpc.NewServer(opts...)

	//srv := grpc.NewServer()
	proto.RegisterReverseServiceServer(srv, &reverseServer{})
	reflection.Register(srv)

	fmt.Println("Server started on port 50051")
	if err := srv.Serve(listen); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
