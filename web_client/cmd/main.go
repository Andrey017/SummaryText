package main

import (
	webclient "web_client"
	"web_client/config"
	"web_client/pkg/handler"

	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	config, errLoadconfig := config.LoadConfig()

	if errLoadconfig != nil {
		logrus.Fatalf("Failed load config 1: %s", errLoadconfig.Error())
	}

	handlers := handler.NewHandler()

	srv := new(webclient.Server)
	err := srv.Run(config.Port, handlers.InitRoutes())

	if err != nil {
		logrus.Fatalf("Failed start http server: %s", err.Error())
	}
}
