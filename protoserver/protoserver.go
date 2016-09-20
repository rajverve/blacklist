package protoserver

import (
	"fmt"
	pb "github.com/rajverve/protobuf"
	"golang.org/x/net/context"
)

type ProtoServer struct {
}

func (s ProtoServer) GetBlacklistInfo(context context.Context, a *pb.AdRequest) (b *pb.BlacklistResponse, err error) {
	fmt.Printf("Received Blacklist Request %v\n", a)

	if a.DeviceId == "abcdef" {
		return &pb.BlacklistResponse{Blacklist: false}, nil
	}

	return &pb.BlacklistResponse{Blacklist: true}, nil

}
