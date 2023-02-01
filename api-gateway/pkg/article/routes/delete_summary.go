package routes

import (
	"api_gateway/pkg/article/pb"
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DeleteSummary(ctx *gin.Context, c pb.ArticleServiceClient) {
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 32)

	res, err := c.DeleteSummary(context.Background(), &pb.DeleteSummaryRequest{
		Id: id,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, &res)
}
