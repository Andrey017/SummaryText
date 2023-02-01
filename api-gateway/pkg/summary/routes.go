package summary

import (
	"api_gateway/pkg/auth"
	"api_gateway/pkg/config"
	"api_gateway/pkg/summary/pb"
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateSummaryTextRequestBody struct {
	Name             string `json:"name"`
	FileName         string `json:"file_name"`
	Type             string `json:"type"`
	Note             string `json:"note"`
	ArticleId        int64  `json:"article_id"`
	CountWord        int64  `json:"countWord"`
	CountSentense    int64  `json:"countSentense"`
	MinCountSentense int64  `json:"minCountSentense"`
	MaxCountSentense int64  `json:"maxCountSentense"`
}

func RegisterRoutes(r *gin.Engine, c *config.Config, authSvc *auth.ServiceClient) *ServiceClient {
	svc := &ServiceClient{
		Client: InitSummaryTextServiceClient(c),
	}

	routes := r.Group("/summaryservice")
	{
		routes.POST("/rank", svc.CreateSummaryTextRank)
		routes.POST("/lsa", svc.CreateSummaryLSA)
		routes.POST("/t5", svc.CreateSummaryT5)
	}

	return svc
}

func (svc *ServiceClient) CreateSummaryTextRank(ctx *gin.Context) {
	body := CreateSummaryTextRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	info := &pb.Summary{
		Name:      body.Name,
		NameFile:  body.FileName,
		Type:      body.Type,
		Note:      body.Note,
		ArticleId: body.ArticleId,
		CountWord: body.CountWord,
	}

	fmt.Println(info)

	res, err := svc.Client.SummaryTextRank(context.Background(), &pb.SummaryTextRequest{
		Info: info,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, &res)
}

func (svc *ServiceClient) CreateSummaryLSA(ctx *gin.Context) {
	body := CreateSummaryTextRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	info := &pb.Summary{
		Name:          body.Name,
		NameFile:      body.FileName,
		Type:          body.Type,
		Note:          body.Note,
		ArticleId:     body.ArticleId,
		CountSentense: body.CountSentense,
	}

	res, err := svc.Client.SummaryLSA(context.Background(), &pb.SummaryLSARequest{
		Info: info,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, &res)
}

func (svc *ServiceClient) CreateSummaryT5(ctx *gin.Context) {
	body := CreateSummaryTextRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	info := &pb.Summary{
		Name:             body.Name,
		NameFile:         body.FileName,
		Type:             body.Type,
		Note:             body.Note,
		ArticleId:        body.ArticleId,
		MinCountSentense: body.MinCountSentense,
		MaxCountSentense: body.MaxCountSentense,
	}

	res, err := svc.Client.SummaryT5(context.Background(), &pb.SummaryT5Request{
		Info: info,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, &res)
}
