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

type HazalertClient struct {
	Client protobuf.HazalertServiceClient
}

func NewHazalertClient() *HazalertClient {
	var opt []grpc.DialOption
	opt = append(opt, grpc.WithTransportCredentials(insecure.NewCredentials()))

	hazalert_host := os.Getenv("HAZALERT_HOST")
	if hazalert_host == "" {
		log.Fatal("HAZALERT_HOST not set")
	}

	hazalert_port := os.Getenv("HAZALERT_PORT")
	if hazalert_port == "" {
		log.Fatal("HAZALERT_PORT not set")
	}
	conn, err := grpc.NewClient(fmt.Sprintf("%s:%s", hazalert_host, hazalert_port), opt...)
	if err != nil {
		log.Println("error on gRPC client: ", err)
	}
	client := protobuf.NewHazalertServiceClient(conn)
	return &HazalertClient{
		Client: client,
	}
}

func (c *HazalertClient) Predict(image []byte) float32 {
	resp, err := c.Client.Predict(context.Background(), &protobuf.HazalertRequest{
		Image: image,
	})
	if err != nil {
		log.Println("error on Predict: ", err)
	}
	return resp.GetNearestDistance()
}
