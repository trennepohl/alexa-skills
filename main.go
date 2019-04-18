package main

import (
	"github.com/labstack/echo"
	dockerRepo "github.com/thiagotrennepohl/alexa-skills/repository/docker"
	alexaHTTPRouter "github.com/thiagotrennepohl/alexa-skills/router"
	"github.com/thiagotrennepohl/alexa-skills/service"
	"github.com/thiagotrennepohl/alexa-skills/service/docker"
	"github.com/thiagotrennepohl/alexa-skills/service/farewell"
	"github.com/thiagotrennepohl/alexa-skills/service/saudation"
)

func main() {
	router := echo.New()

	dockerRepository := dockerRepo.NewDockerRepository()

	intents := make(map[string]service.Intent)

	dockerService := docker.NewDockerService(dockerRepository)
	saudationService := saudation.NewSaudationService()
	farewellService := farewell.NewFarewellService()

	intents = map[string]service.Intent{
		"saudation":        saudationService,
		"farewell":         farewellService,
		"restartContainer": dockerService,
		"listContainers":   dockerService,
	}

	alexaHTTPRouter.NewSaudationRouter(router, intents)
	router.Start(":8080")
}
