package summary

import (
	"api_gateway/pkg/config"
	"api_gateway/pkg/summary/pb"
	"fmt"

	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.SummaryTextServiceClient
}

func InitSummaryTextServiceClient(c *config.Config) pb.SummaryTextServiceClient {
	cc, err := grpc.Dial(c.SummaryTextSvcUrl, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return pb.NewSummaryTextServiceClient(cc)
}
