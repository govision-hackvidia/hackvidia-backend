package service

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/govision-hackvidia/hackvidia-backend/protobuf"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type MllmClient struct {
	Client protobuf.MllmServiceClient
}

func NewMllmClient() *MllmClient {
	var opt []grpc.DialOption
	opt = append(opt, grpc.WithTransportCredentials(insecure.NewCredentials()))
	mllm_host := os.Getenv("MLLM_HOST")
	if mllm_host == "" {
		log.Fatal("MLLM_HOST not set")
	}
	mllm_port := os.Getenv("MLLM_PORT")
	if mllm_port == "" {
		log.Fatal("MLLM_PORT not set")
	}
	conn, err := grpc.NewClient(fmt.Sprintf("%s:%s", mllm_host, mllm_port), opt...)
	if err != nil {
		log.Println("error on gRPC client: ", err)
	}
	client := protobuf.NewMllmServiceClient(conn)
	return &MllmClient{
		Client: client,
	}
}

func (c *MllmClient) Chat(text string, image []byte) string {
	resp, err := c.Client.Chat(context.Background(), &protobuf.MllmRequest{
		Text:  text,
		Image: image,
	})
	if err != nil {
		log.Println("error on Chat: ", err)
	}

	return resp.GetText()
}
