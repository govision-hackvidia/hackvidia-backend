package service

import (
	"context"
	"log"

	"github.com/govision-hackvidia/hackvidia-backend/protobuf"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GrpcClient struct {
	Client protobuf.MllmServiceClient
}

func NewGrpcClient() *GrpcClient {
	var opt []grpc.DialOption
	opt = append(opt, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.NewClient("localhost:50051", opt...)
	if err != nil {
		log.Println("error on gRPC client: ", err)
	}
	client := protobuf.NewMllmServiceClient(conn)
	return &GrpcClient{
		Client: client,
	}
}

func (c *GrpcClient) Chat() {
	resp, err := c.Client.Chat(context.Background(), &protobuf.MllmRequest{
		Text: "ABC",
	})
	if err != nil {
		log.Println("error on Chat: ", err)
	}

	log.Println(resp)
	// c.Client.Chat()
}
