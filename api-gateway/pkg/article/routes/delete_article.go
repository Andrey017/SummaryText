package routes

import (
	"api_gateway/pkg/article/pb"
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DeleteArticle(ctx *gin.Context, c pb.ArticleServiceClient) {
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 32)

	res, err := c.DeleteArticle(context.Background(), &pb.DeleteArticleRequest{
		Id: id,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, &res)
}
