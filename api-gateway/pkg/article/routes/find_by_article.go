package routes

import (
	"api_gateway/pkg/article/pb"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FindArticleBy(ctx *gin.Context, c pb.ArticleServiceClient) {
	name := ctx.Param("name")
	userId, _ := getUserId(ctx)

	res, err := c.FindArticleByName(context.Background(), &pb.FindArticleByNameRequest{
		Name:   name,
		UserId: userId,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, &res)
}
