package main

// https://medium.com/zeals-tech-blog/grpc-go-server-web-client-853daf40cdef
import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/improbable-eng/grpc-web/go/grpcweb"

	pb "go.server.web/v1/proto"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

// Ensure that reverseServer implements the proto.ReverseServiceServer interface
var _ pb.ReverseServiceServer = &reverseServer{}

// server is used to implement proto.ReverseServiceServer.
type reverseServer struct {
	pb.UnimplementedReverseServiceServer
}

// ReverseString implements proto.ReverseServiceServer
func (s *reverseServer) ReverseString(ctx context.Context, req *pb.ReverseRequest) (*pb.ReverseReply, error) {
	reversed := reverseString(req.GetData())
	return &pb.ReverseReply{Reversed: reversed}, nil
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterReverseServiceServer(s, &reverseServer{})
	log.Printf("grpc server listening at %v", lis.Addr())
	go func() {
		log.Fatalf("failed to serve: %v", s.Serve(lis))
	}()

	// gRPC web code
	grpcWebServer := grpcweb.WrapServer(
		s,
		// Enable CORS
		grpcweb.WithOriginFunc(func(origin string) bool { return true }),
	)

	srv := &http.Server{
		Handler: grpcWebServer,
		Addr:    fmt.Sprintf("localhost:%d", *port+1),
	}

	log.Printf("http server listening at %v", srv.Addr)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
