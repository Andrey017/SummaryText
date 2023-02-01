package routes

import (
	"api_gateway/pkg/article/pb"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type FindOneSummaryRequest struct {
	Name      string `json:"name"`
	ArticleId int64  `json:"article_id"`
}

func FindOneSummary(ctx *gin.Context, c pb.ArticleServiceClient) {
	body := FindOneSummaryRequest{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.FindOneSummary(context.Background(), &pb.FindOneSummaryRequest{
		Name:      body.Name,
		ArticleId: body.ArticleId,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, &res)
}
