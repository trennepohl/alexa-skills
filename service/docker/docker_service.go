package docker

import (
	"fmt"

	"github.com/thiagotrennepohl/alexa-skills/models"
	"github.com/thiagotrennepohl/alexa-skills/repository"
	"github.com/thiagotrennepohl/alexa-skills/service"
)

type dockerService struct {
	dockerRepo repository.DockerRepository
}

func NewDockerService(dockerRepository repository.DockerRepository) service.Intent {
	return &dockerService{
		dockerRepo: dockerRepository,
	}
}

func (d *dockerService) Response(alexaRequest models.AlexaIncomingPayload) (models.AlexaResponse, error) {
	d.dockerRepo.Connect("localhost:2376")
	alexaResponse := models.AlexaResponse{}
	// fmt.Printf("%+v\n", alexaRequest.Request.Intent.Slots)
	// alexaRequest.Request.Intent.Slots["clientname"].(map[string]interface{})["value"].(string)
	containerName := alexaRequest.Request.Intent.Slots["container_name"].(map[string]interface{})["value"].(string)
	fmt.Println(containerName)
	err := d.dockerRepo.RestartContainer(alexaRequest.Request.Intent.Slots["container_name"].(map[string]interface{})["value"].(string))
	if err != nil {
		alexaResponse.Response.OutputSpeech.Text = "Failed"
		alexaResponse.Response.OutputSpeech.Type = "PlainText"
		alexaResponse.ShouldEndSession = true
	}
	alexaResponse.Response.OutputSpeech.Text = "Restarting banana container"
	alexaResponse.Response.OutputSpeech.Type = "PlainText"
	alexaResponse.ShouldEndSession = true
	return alexaResponse, nil
}
