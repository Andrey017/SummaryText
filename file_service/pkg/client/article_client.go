package client

import (
	"context"
	"file_service/pkg/pb"
	"fmt"

	"google.golang.org/grpc"
)

type ArticleClient struct {
	Client pb.ArticleServiceClient
}

func InitArticleServiceClient(url string) ArticleClient {
	cc, err := grpc.Dial(url, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connection:", err)
	}

	c := ArticleClient{
		Client: pb.NewArticleServiceClient(cc),
	}

	return c
}

func (c *ArticleClient) UpdateArticleFileName(id int64, nameFile string) (*pb.UpdateArticleFileNameResponse, error) {
	req := &pb.UpdateArticleFileNameRequest{
		Id:       id,
		NameFile: nameFile,
	}

	return c.Client.UpdateArticleFileName(context.Background(), req)
}
