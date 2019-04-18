package docker

import (
	"fmt"

	"github.com/thiagotrennepohl/alexa-skills/models"
	"github.com/thiagotrennepohl/alexa-skills/repository"
	"github.com/thiagotrennepohl/alexa-skills/service"
)

type dockerService struct {
	dockerRepo repository.DockerRepository
	skills     map[string]func(string) (models.AlexaResponse, error)
}

func NewDockerService(dockerRepository repository.DockerRepository) service.Intent {
	dockerSvc := &dockerService{
		dockerRepo: dockerRepository,
	}

	dockerSvc.skills = map[string]func(string) (models.AlexaResponse, error){
		"restartContainer": dockerSvc.restartContainer,
	}

	err := dockerSvc.dockerRepo.Connect("localhost:2376")
	if err != nil {
		panic(err)
	}

	return dockerSvc
}

func (d *dockerService) Response(alexaRequest models.AlexaIncomingPayload) (models.AlexaResponse, error) {

	alexaResponse := models.AlexaResponse{}
	intentName := alexaRequest.Request.Intent.Name
	containerName := alexaRequest.Request.Intent.Slots["container_name"].(map[string]interface{})["value"].(string)

	if _, ok := d.skills[intentName]; !ok {
		fmt.Println("deu ruim aqui")
		alexaResponse.SetPlainTextErrorResponse(fmt.Sprintf("Sorry I couldn't find an action called %s", intentName))
		return alexaResponse, fmt.Errorf("Skill %s not found", intentName)
	}

	return d.skills[alexaRequest.Request.Intent.Name](containerName)
}

func (d *dockerService) restartContainer(containerName string) (models.AlexaResponse, error) {
	alexaResponse := models.AlexaResponse{}
	err := d.dockerRepo.RestartContainer(containerName)
	if err != nil {
		fmt.Println(err.Error())
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
