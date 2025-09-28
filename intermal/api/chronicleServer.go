package api

import (
	"Lab1/intermal/app/handler"
	"Lab1/intermal/app/repository"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func StartServer() {
	log.Println("Starting server")

	chronicleModel, err := repository.NewChronicleModel()
	if err != nil {
		logrus.Error("ошибка инициализации репозитория")
	}

	chronicleController := handler.NewChronicleController(chronicleModel)

	r := gin.Default()
	// добавляем наш html/шаблон
	r.LoadHTMLGlob("templates/*")
	r.Static("/static", "./resources")
	// слева название папки, в которую выгрузится наша статика
	// справа путь к папке, в которой лежит статика

	r.GET("/chronicles", chronicleController.GetChronicleResources)
	r.GET("/chronicle/:id", chronicleController.GetChronicleResource)
	r.GET("/chronicle-research/:id", chronicleController.GetChronicleApplication)

	r.Run()
	log.Println("Server down")
}
