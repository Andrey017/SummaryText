package routes

import (
	"api_gateway/pkg/article/pb"
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateArticleRequestBody struct {
	Name       string `json:"name"`
	FileName   string `json:"file_name"`
	Note       string `json:"note"`
	DateCreate string `json:"date_create"`
}

func CreateArticle(ctx *gin.Context, c pb.ArticleServiceClient) {
	body := CreateArticleRequestBody{}
	userId, _ := getUserId(ctx)

	if err := ctx.BindJSON(&body); err != nil {
		fmt.Println(body)

		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.CreateArticle(context.Background(), &pb.CreateArticleRequest{
		Name:       body.Name,
		NameFile:   body.FileName,
		Note:       body.Note,
		DateCreate: body.DateCreate,
		UserId:     userId,
	})

	if err != nil {
		fmt.Println(err)

		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}

func getUserId(ctx *gin.Context) (int64, error) {
	userId, ok := ctx.Get("userId")

	if !ok {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return 0, errors.New("Not found user id context")
	}

	userIdInt, ok := userId.(int64)

	if !ok {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return 0, errors.New("Eror user id context")
	}

	return userIdInt, nil
}

/*func CreateArticle(ctx *gin.Context, c pb.ArticleServiceClient) {
	body := CreateArticleRequestBody{}
	userId, _ := getUserId(ctx)

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	info := &pb.CreateInfo{
		Name:       body.Name,
		NameFile:   body.FileName,
		Note:       body.Note,
		DateCreate: body.DateCreate,
		UserId:     userId,
	}

	fmt.Println(info)

	ctx.Request.Body = http.MaxBytesReader(ctx.Writer, ctx.Request.Body, MAX_UPLOAD_SIZE)

	file, fileHeader, err := ctx.Request.FormFile("file")

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	buffer := make([]byte, fileHeader.Size)
	size := 0

	stream, err := c.CreateArticle(ctx)

	for {
		n, err := reader.Read(buffer)
		if err == io.EOF {
			break
		}
		size += n

		req := &pb.CreateArticleRequest{
			Info: info,
			File: buffer[:n],
		}

		err = stream.Send(req)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
*/
