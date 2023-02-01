package article

import (
	"api_gateway/pkg/article/routes"
	"api_gateway/pkg/auth"
	"api_gateway/pkg/config"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, c *config.Config, authSvc *auth.ServiceClient) *ServiceClient {
	a := auth.InitAuthMiddleware(authSvc)

	svc := &ServiceClient{
		Client: InitArticleServiceClient(c),
	}

	routes := r.Group("/article")
	{
		routes.Use(a.AuthRequired)

		routes.POST("/", svc.CreateArticle)
		routes.GET("/", svc.FindAllArticle)
		routes.GET("/:id", svc.FindOneArticle)
		routes.GET("/by/:name", svc.FindArticleBy)
		routes.PUT("/:id", svc.UpdateArticle)
		routes.DELETE("/:id", svc.DeleteArticle)

		routes.POST("/load", svc.LoadFile)

		summary := routes.Group("/summary")
		{
			summary.Use(a.AuthRequired)

			summary.POST("/", svc.CreateSummary)
			summary.GET("/:article_id", svc.FindAllSummary)
			summary.POST("/one", svc.FindOneSummary)
			summary.PUT("/:id", svc.UpdateSummary)
			summary.DELETE("/:id", svc.DeleteSummary)
		}
	}

	return svc
}

const MAX_UPLOAD_SIZE = 32 << 20

func (svc *ServiceClient) LoadFile(ctx *gin.Context) {
	/*srcFile, _ := os.Open("test.txt")
	defer srcFile.Close()

	fmt.Println(srcFile)

	destFile, _ := os.Create("test_copy.txt") // creates if file doesn't exist
	defer destFile.Close()

	_, _ = io.Copy(destFile, srcFile) // check first var for number of bytes copied*/

	//err = destFile.Sync()

	/*ctx.Request.Body = http.MaxBytesReader(ctx.Writer, ctx.Request.Body, MAX_UPLOAD_SIZE)

	file, fileHeader, err := ctx.Request.FormFile("file")

	fmt.Println(err)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	defer file.Close()

	buffer := make([]byte, fileHeader.Size)

	file.Read(buffer)
	//fileType := http.DetectContentType(buffer)

	dst, err := os.Create(filepath.Join("pkg/article", fileHeader.Filename))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	defer dst.Close()

	if _, err := io.Copy(dst, file); err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(err)*/

}

//Article
func (svc *ServiceClient) CreateArticle(ctx *gin.Context) {
	routes.CreateArticle(ctx, svc.Client)
}

func (svc *ServiceClient) FindAllArticle(ctx *gin.Context) {
	routes.FindAllArticle(ctx, svc.Client)
}

func (svc *ServiceClient) FindOneArticle(ctx *gin.Context) {
	routes.FindOneArticle(ctx, svc.Client)
}

func (svc *ServiceClient) FindArticleBy(ctx *gin.Context) {
	routes.FindArticleBy(ctx, svc.Client)
}

func (svc *ServiceClient) UpdateArticle(ctx *gin.Context) {
	routes.UpdateArticle(ctx, svc.Client)
}

func (svc *ServiceClient) DeleteArticle(ctx *gin.Context) {
	routes.DeleteArticle(ctx, svc.Client)
}

//Summary
func (svc *ServiceClient) CreateSummary(ctx *gin.Context) {
	routes.CreateSummary(ctx, svc.Client)
}

func (svc *ServiceClient) FindAllSummary(ctx *gin.Context) {
	routes.FindAllSummary(ctx, svc.Client)
}

func (svc *ServiceClient) FindOneSummary(ctx *gin.Context) {
	routes.FindOneSummary(ctx, svc.Client)
}

func (svc *ServiceClient) UpdateSummary(ctx *gin.Context) {
	routes.UpdateSummary(ctx, svc.Client)
}

func (svc *ServiceClient) DeleteSummary(ctx *gin.Context) {
	routes.DeleteSummary(ctx, svc.Client)
}
