package routes

import (
	"api_gateway/pkg/article/pb"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateSummaryRequestBody struct {
	Name      string `json:"name"`
	FileName  string `json:"file_name"`
	Type      string `json:"type"`
	Note      string `json:"note"`
	ArticleId int64  `json:"article_id"`
}

func CreateSummary(ctx *gin.Context, c pb.ArticleServiceClient) {
	body := CreateSummaryRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.CreateSummary(context.Background(), &pb.CreateSummaryRequest{
		Name:      body.Name,
		NameFile:  body.FileName,
		Type:      body.Type,
		Note:      body.Note,
		ArticleId: body.ArticleId,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
