package main

import (
	"fmt"
	"github.com/rajverve/blacklist/protoserver"
	pb "github.com/rajverve/protobuf"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 3333))
	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
		return
	}

	fmt.Println("Starting Blacklist server...")
	grpcServer := grpc.NewServer()
	pb.RegisterBlacklistServer(grpcServer, protoserver.ProtoServer{})
	grpcServer.Serve(lis)
}
