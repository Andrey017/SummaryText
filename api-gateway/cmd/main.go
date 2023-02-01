package main

import (
	"api_gateway/pkg/article"
	"api_gateway/pkg/auth"
	"api_gateway/pkg/config"
	"api_gateway/pkg/file"
	"api_gateway/pkg/summary"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	r := gin.Default()

	authSvc := *auth.RegisterRoutes(r, &c)
	article.RegisterRoutes(r, &c, &authSvc)
	file.RegisterRoutes(r, &c, &authSvc)
	summary.RegisterRoutes(r, &c, &authSvc)

	r.Run(c.Port)
}
