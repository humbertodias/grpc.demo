package main

import (
	"context"
	"fmt"
	proto "github.com/humbertodias/grpc.demo/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

type reverseServer struct{}

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

	srv := grpc.NewServer()
	proto.RegisterReverseServiceServer(srv, &reverseServer{})

	fmt.Println("Server started on port 50051")
	if err := srv.Serve(listen); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
