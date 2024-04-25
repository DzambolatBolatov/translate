package main

import (
	"dictionary/internal/service"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func main() {
	// создаем логгер
	logger := log.New("dict")

	// подключаемся к базе
	db, err := PostgresConnection()
	if err != nil {
		logger.Fatal(err)
	}

	svc := service.NewService(db, logger)

	router := echo.New()
	// создаем группу api
	api := router.Group("api")

	// прописываем пути
	api.GET("/word/:id", svc.GetWordById)
	api.POST("/words", svc.CreateWords)
	api.PUT("/word", svc.UpdateWordsById)
	api.DELETE("/word", svc.DeleteWordById)

	api.POST("/report", svc.CreateReport)
	api.GET("/report/:id", svc.GetReport)
	api.PUT("/report/:id", svc.UpdateReport)
	api.DELETE("/report/:id", svc.DeleteReport)

	api.GET("/search/ru?title", svc.SearchRu)

	// запускаем сервер, чтобы слушал 8000 порт
	router.Logger.Fatal(router.Start(":8000"))
}
