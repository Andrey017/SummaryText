package routes

import (
	"api_gateway/pkg/article/pb"
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UpdateArticleRequestBody struct {
	Name       string `json:"name"`
	FileName   string `json:"file_name"`
	Note       string `json:"note"`
	DateCreate string `json:"date_create"`
}

func UpdateArticle(ctx *gin.Context, c pb.ArticleServiceClient) {
	body := UpdateArticleRequestBody{}
	userId, _ := getUserId(ctx)
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 32)

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.UpdateArticle(context.Background(), &pb.UpdateArticleRequest{
		Id:         id,
		Name:       body.Name,
		NameFile:   body.FileName,
		Note:       body.Note,
		DateCreate: body.DateCreate,
		UserId:     userId,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, &res)
}
