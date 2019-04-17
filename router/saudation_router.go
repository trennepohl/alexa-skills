package router

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/thiagotrennepohl/alexa-skills/models"
	"github.com/thiagotrennepohl/alexa-skills/service"
)

type alexaRouter struct {
	engine  *echo.Echo
	intents map[string]service.Intent
}

func NewSaudationRouter(engine *echo.Echo, intents map[string]service.Intent) {
	router := alexaRouter{engine: engine, intents: intents}
	router.engine.POST("/v1/incomming", router.processIncoming)
}

func (a *alexaRouter) processIncoming(ctx echo.Context) error {
	payload := &models.AlexaIncomingPayload{}
	err := ctx.Bind(&payload)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, models.JSON{"message": err.Error()})
	}

	fmt.Println(payload.Request.Intent.Name)

	if _, ok := a.intents[payload.Request.Intent.Name]; !ok {
		return ctx.JSON(http.StatusNotFound, models.JSON{"message": "Intent not found"})
	}

	alexaResponse, err := a.intents[payload.Request.Intent.Name].Response(*payload)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, models.JSON{"message": err.Error()})
	}

	return ctx.JSON(http.StatusOK, &alexaResponse)
}
