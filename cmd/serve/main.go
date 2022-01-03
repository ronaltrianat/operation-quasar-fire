package main

import (
	"github.com/labstack/echo/v4"

	"github.com/ronaltrianat/operation-quasar-fire/src/adapters/http"
	"github.com/ronaltrianat/operation-quasar-fire/src/core/services"
	"github.com/ronaltrianat/operation-quasar-fire/src/core/usecases"
)

func main() {

	messageUseCase := usecases.NewMessageUseCase()
	trilaterationUseCase := usecases.NewTrilaterationUseCase()
	topSecretService := services.NewTopSecretService(trilaterationUseCase, messageUseCase)
	httpHandler := http.NewHTTPHandler(topSecretService)

	e := echo.New()
	e.POST("/top-secret", httpHandler.TopSecret)

	e.Logger.Fatal(e.Start(":1323"))
}
