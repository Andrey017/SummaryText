package article

import (
	"api_gateway/pkg/article/pb"
	"api_gateway/pkg/config"
	"fmt"

	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.ArticleServiceClient
}

func InitArticleServiceClient(c *config.Config) pb.ArticleServiceClient {
	cc, err := grpc.Dial(c.ArticleSvcUrl, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return pb.NewArticleServiceClient(cc)
}
