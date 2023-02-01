package routes

import (
	"api_gateway/pkg/article/pb"
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UpdateSummaryRequestBody struct {
	Name      string `json:"name"`
	FileName  string `json:"file_name"`
	Type      string `json:"type"`
	Note      string `json:"note"`
	ArticleId int64  `json:"article_id"`
}

func UpdateSummary(ctx *gin.Context, c pb.ArticleServiceClient) {
	body := UpdateSummaryRequestBody{}

	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 32)

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.UpdateSummary(context.Background(), &pb.UpdateSummaryRequest{
		Id:        id,
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

	ctx.JSON(http.StatusOK, &res)
}
