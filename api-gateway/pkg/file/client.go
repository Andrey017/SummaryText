package file

import (
	"api_gateway/pkg/config"
	"api_gateway/pkg/file/pb"
	"fmt"

	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.FileServiceClient
}

func InitFileServiceClient(c *config.Config) pb.FileServiceClient {
	cc, err := grpc.Dial(c.FileSvcUrl, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return pb.NewFileServiceClient(cc)
}
