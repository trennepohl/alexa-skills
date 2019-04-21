package docker

import (
	"bytes"
	"fmt"
	"html/template"

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
		"stopContainer":    dockerSvc.stopContainer,
		"startContainer":   dockerSvc.startContainer,
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
	Containers := models.Containers{}
	cNames, err := d.dockerRepo.ListContainers()
	if err != nil {
		fmt.Println(err)
		alexaResponse.SetPlainTextErrorResponse(err.Error())
		return alexaResponse, err
	}

	Containers.ContainerNames = cNames
	cNames = nil

	if len(Containers.ContainerNames) < 1 {
		alexaResponse.SetPlainTextErrorResponse("no containers running")
		return alexaResponse, nil
	}

	tmpl, err := template.New("response.tmpl").ParseFiles("templates/response.tmpl")
	if err != nil {
		panic(err)
	}

	var response bytes.Buffer
	err = tmpl.Execute(&response, Containers)
	if err != nil {
		panic(err)
	}

	fmt.Println(response.String())

	alexaResponse.SetSSMLResponse(response.String())
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

func (d *dockerService) stopContainer(alexaRequest *models.AlexaIncomingPayload) (models.AlexaResponse, error) {
	alexaResponse := models.AlexaResponse{}
	if _, ok := alexaRequest.Request.Intent.Slots["container_name"]; !ok {
		alexaResponse.SetPlainTextErrorResponse("Sorry the container name is not being specified")
		return alexaResponse, fmt.Errorf("%s", "Sorry the container name is not being specified")
	}
	containerName := alexaRequest.Request.Intent.Slots["container_name"].(map[string]interface{})["value"].(string)
	err := d.dockerRepo.StopContainer(containerName)
	if err != nil {
		alexaResponse.SetPlainTextErrorResponse(err.Error())
		return alexaResponse, err
	}
	alexaResponse.SetPlainTextResponse(fmt.Sprintf("%s has been stopped successfully", containerName))
	return alexaResponse, err
}

func (d *dockerService) startContainer(alexaRequest *models.AlexaIncomingPayload) (models.AlexaResponse, error) {
	alexaResponse := models.AlexaResponse{}
	if _, ok := alexaRequest.Request.Intent.Slots["container_name"]; !ok {
		fmt.Println("caiu")
		alexaResponse.SetPlainTextErrorResponse("Sorry the container name is not being specified")
		return alexaResponse, fmt.Errorf("%s", "Sorry the container name is not being specified")
	}
	containerName := alexaRequest.Request.Intent.Slots["container_name"].(map[string]interface{})["value"].(string)
	err := d.dockerRepo.StartContainer(containerName)
	if err != nil {
		fmt.Println(err.Error())
		alexaResponse.SetPlainTextErrorResponse(err.Error())
		return alexaResponse, err
	}
	alexaResponse.SetPlainTextResponse(fmt.Sprintf("%s has been started successfully", containerName))
	return alexaResponse, err
}
