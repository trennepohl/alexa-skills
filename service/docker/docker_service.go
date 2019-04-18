package docker

import (
	"fmt"

	"github.com/thiagotrennepohl/alexa-skills/models"
	"github.com/thiagotrennepohl/alexa-skills/repository"
	"github.com/thiagotrennepohl/alexa-skills/service"
)

type dockerService struct {
	dockerRepo repository.DockerRepository
	skills     map[string]func(*models.AlexaIncomingPayload) (models.AlexaResponse, error)
}

func NewDockerService(dockerRepository repository.DockerRepository) service.Intent {
	dockerSvc := &dockerService{
		dockerRepo: dockerRepository,
	}

	dockerSvc.skills = map[string]func(*models.AlexaIncomingPayload) (models.AlexaResponse, error){
		"restartContainer": dockerSvc.restartContainer,
		"listContainers":   dockerSvc.listContainers,
	}

	err := dockerSvc.dockerRepo.Connect("localhost:2376")
	if err != nil {
		panic(err)
	}

	return dockerSvc
}

func (d *dockerService) Response(alexaRequest models.AlexaIncomingPayload) (models.AlexaResponse, error) {
	return d.skills[alexaRequest.Request.Intent.Name](&alexaRequest)
}

func (d *dockerService) listContainers(alexaRequest *models.AlexaIncomingPayload) (models.AlexaResponse, error) {
	alexaResponse := models.AlexaResponse{}
	containerNames, err := d.dockerRepo.ListContainers()
	if err != nil {
		fmt.Println(err)
		alexaResponse.SetPlainTextErrorResponse(err.Error())
		return alexaResponse, err
	}

	var message string

	if len(containerNames) >= 1 {
		message = "Here is the list of all running containers"
	}

	for _, name := range containerNames {
		message += name + ","
	}
	alexaResponse.SetPlainTextResponse(message)
	return alexaResponse, err
}

func (d *dockerService) restartContainer(alexaRequest *models.AlexaIncomingPayload) (models.AlexaResponse, error) {
	alexaResponse := models.AlexaResponse{}
	if _, ok := alexaRequest.Request.Intent.Slots["container_name"]; !ok {
		alexaResponse.SetPlainTextErrorResponse("Sorry the container name is not being specified")
		return alexaResponse, fmt.Errorf("%s", "Sorry the container name is not being specified")
	}
	containerName := alexaRequest.Request.Intent.Slots["container_name"].(map[string]interface{})["value"].(string)
	err := d.dockerRepo.RestartContainer(containerName)
	if err != nil {
		alexaResponse.SetPlainTextErrorResponse(err.Error())
		return alexaResponse, err
	}
	alexaResponse.SetPlainTextResponse(fmt.Sprintf("%s has been restarted successfully", containerName))
	return alexaResponse, err
}

// func (d *dockerService) listContainers(containerName string) (models.AlexaResponse, error) {
// 	alexaResponse := models.AlexaResponse{}
// 	err := d.dockerRepo.RestartContainer(containerName)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		alexaResponse.SetPlainTextErrorResponse(err.Error())
// 		return alexaResponse, err
// 	}

// 	alexaResponse.SetPlainTextResponse(fmt.Sprintf("%s has been restarted successfully", containerName))

// 	return alexaResponse, err
// }
