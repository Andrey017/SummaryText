package handler

import (
	"github.com/gin-gonic/gin"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) InitRoutes() *gin.Engine {
	routes := gin.New()

	routes.LoadHTMLGlob("templates/*")

	routes.GET("/", getHome) //HomePage

	authRoutes := routes.Group("/u")
	{
		authRoutes.GET("/login", showLoginPage) //Show login page
		authRoutes.POST("/login", performLogin)

		authRoutes.GET("/reg", showRegPage) //Show reg page
		authRoutes.POST("/reg", performReg)
	}

	api := routes.Group("/api")
	{
		api.GET("/create", showCreateArticle)
		api.POST("/create", createArticleInfo)

		api.GET("/create/loadFile/:id", showLoadFileArticle)

		article := api.Group("/article")
		{
			article.GET("/:id", showArticle)

			article.GET("/delete/:id", deleteArticle)

			article.GET("/update/:id", showUpdateArticle)
			article.POST("/update/:id", updateArticleInfo)
		}

		summary := api.Group("/summary")
		{
			summary.GET("/create/:id", showCreateSummary)
			summary.POST("/create/:id", createSummary)

			summary.GET("/delete/:id", deleteSummary)
		}
	}

	return routes
}
