package routes

import (
	"api_gateway/pkg/article/pb"
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func FindAllSummary(ctx *gin.Context, c pb.ArticleServiceClient) {
	articleId, _ := strconv.ParseInt(ctx.Param("article_id"), 10, 32)

	res, err := c.FindAllSummary(context.Background(), &pb.FindAllSummaryRequest{
		ArticleId: articleId,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, &res)
}
