package file

import (
	"api_gateway/pkg/auth"
	"api_gateway/pkg/config"
	"api_gateway/pkg/file/pb"
	"bufio"
	"context"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

const MAX_UPLOAD_SIZE = 5 << 20

type FileInfo struct {
	FileName string
	FileType string
}

type DownloadFileRequestBody struct {
	Id       int64  `json:"id"`
	NameFile string `json:"name_file"`
}

func RegisterRoutes(r *gin.Engine, c *config.Config, authSvc *auth.ServiceClient) *ServiceClient {
	//a := auth.InitAuthMiddleware(authSvc)

	svc := &ServiceClient{
		Client: InitFileServiceClient(c),
	}

	routes := r.Group("/upload")
	{
		routes.POST("/load/:id", svc.LoadFile)
		//routes.POST("/", svc.DownloadFile)
		routes.GET("/:idNameFile", svc.DownloadFile)
	}

	return svc
}

func (svc *ServiceClient) LoadFile(ctx *gin.Context) {
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 32)
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

	stream, err := svc.Client.UploadFileArticle(ctx)

	info := &pb.FileInfo{
		Id:       id,
		FileName: fileHeader.Filename,
		FileType: ".pdf",
	}

	for {
		n, err := reader.Read(buffer)
		if err == io.EOF {
			break
		}
		size += n

		req := &pb.UploadFileRequest{
			Info:     info,
			FileData: buffer[:n],
		}

		err = stream.Send(req)
	}

	_, err = stream.CloseAndRecv()

	ctx.Redirect(http.StatusFound, "http://192.168.17.247:9005/")
	//ctx.JSON(http.StatusOK, &res)
}

func (svc *ServiceClient) DownloadFile(ctx *gin.Context) {
	/*body := DownloadFileRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	id := body.Id
	nameFile := body.NameFile*/

	idNameFile := ctx.Param("idNameFile")

	split := strings.Split(idNameFile, "&")
	id, _ := strconv.ParseInt(split[0], 10, 32)
	nameFile := split[1]

	fileStreamResponse, err := svc.Client.DownloadFile(context.TODO(), &pb.DownloadFileRequest{
		Id:       id,
		FileName: nameFile,
	})
	if err != nil {
		log.Println("error downloading:", err)
		return
	}
	for {
		chunkResponse, err := fileStreamResponse.Recv()
		if err == io.EOF {
			log.Println("received all chunks")
			break
		}
		if err != nil {
			log.Println("err receiving chunk:", err)
			break
		}

		ctx.Data(http.StatusOK, "pdf", chunkResponse.FileData)
	}
}
