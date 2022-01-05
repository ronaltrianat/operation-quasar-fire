package main

import (
	"github.com/labstack/echo/v4"

	"github.com/ronaltrianat/operation-quasar-fire/src/adapters/http"
	"github.com/ronaltrianat/operation-quasar-fire/src/adapters/repositories"
	"github.com/ronaltrianat/operation-quasar-fire/src/core/services"
	"github.com/ronaltrianat/operation-quasar-fire/src/core/usecases"
)

func main() {

	messageUseCase := usecases.NewMessageUseCase()
	trilaterationUseCase := usecases.NewTrilaterationUseCase()
	repository := repositories.NewInMemoryDB()

	topSecretService := services.NewTopSecretService(trilaterationUseCase, messageUseCase, repository)

	httpHandler := http.NewHTTPHandler(topSecretService)

	e := echo.New()
	e.POST("/topsecret", httpHandler.TopSecret)
	e.POST("/topsecret_split/:name", httpHandler.SaveTopSecretSplit)
	e.GET("/topsecret_split", httpHandler.GetTopSecretSplit)
	e.GET("/", httpHandler.Health)

	e.Logger.Fatal(e.Start(":5000"))
}
