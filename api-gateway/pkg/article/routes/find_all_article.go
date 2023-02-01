package routes

import (
	"api_gateway/pkg/article/pb"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FindAllArticle(ctx *gin.Context, c pb.ArticleServiceClient) {
	userId, _ := getUserId(ctx)

	res, err := c.FindAllArticle(context.Background(), &pb.FindAllArticleRequest{
		UserId: userId,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, &res)
}
