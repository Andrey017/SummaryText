package routes

import (
	"api_gateway/pkg/article/pb"
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func FindOneArticle(ctx *gin.Context, c pb.ArticleServiceClient) {
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 32)
	userId, _ := getUserId(ctx)

	res, err := c.FindOneArticle(context.Background(), &pb.FindOneArticleRequest{
		Id:     id,
		UserId: userId,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, &res)
}
