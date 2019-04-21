package docker

import (
	"fmt"
	"strings"

	gdocker "github.com/fsouza/go-dockerclient"
	"github.com/thiagotrennepohl/alexa-skills/repository"
	"gitlab.agilepromoter.com/iat/apoc/models"
)

type dockerRepository struct {
	client *gdocker.Client
}

func NewDockerRepository() repository.DockerRepository {
	return &dockerRepository{}
}

func (docker *dockerRepository) OpenTLSDockerConnection(APIAddress string, CACertPath, keyPath, certPath string) (gdocker.Client, error) {
	client, err := gdocker.NewTLSClient(
		"tcp://"+APIAddress,
		certPath,
		keyPath,
		CACertPath,
	)
	if err != nil {
		return *client, &models.ErrCannotOpenDockerConnection{Message: err.Error()}
	}

	client.TLSConfig.InsecureSkipVerify = true

	return *client, err
}

func (docker *dockerRepository) Connect(endpoint string) error {
	return docker.OpenDockerConnection(endpoint)
}

func (docker *dockerRepository) OpenDockerConnection(APIAddress string) error {
	client, err := gdocker.NewClient(
		"tcp://" + APIAddress,
	)
	if err != nil {
		return &models.ErrCannotOpenDockerConnection{Message: err.Error()}
	}

	docker.client = client
	return nil
}

func (docker *dockerRepository) RestartContainer(containerName string) error {
	containerID, err := docker.GetContainerID(containerName)
	if err != nil {
		return &models.ErrCannotRestartClient{Message: err.Error()}
	}

	err = docker.client.RestartContainer(containerID, dockerRestartTimeout)
	if err != nil {
		return &models.ErrCannotRestartClient{Message: err.Error()}
	}

	return err
}

// func (docker *dockerRepository) RemoveContainer(conn gdocker.Client, containerName string) error {
// 	containerID, err := docker.GetContainerID(conn, containerName)
// 	if err != nil {
// 		return &models.ErrCannotRemoveContainer{Message: cannotDeleteContainer + err.Error()}
// 	}
// 	opts := gdocker.RemoveContainerOptions{
// 		ID:    containerID,
// 		Force: true,
// 	}
// 	err = conn.RemoveContainer(opts)
// 	if err != nil {
// 		return &models.ErrCannotRemoveContainer{Message: cannotDeleteContainer + err.Error()}
// 	}
// 	return err
// }

func (docker *dockerRepository) ListContainers() ([]string, error) {
	containerNames := []string{}
	runningContainers, err := docker.client.ListContainers(gdocker.ListContainersOptions{All: true})
	if err != nil {
		return containerNames, nil
	}
	for _, container := range runningContainers {
		for _, name := range container.Names {
			containerNames = append(containerNames, docker.removeSlashes(name))
		}
	}
	return containerNames, nil
}

func (docker *dockerRepository) StartContainer(containerName string) error {
	containerID, err := docker.GetContainerID(containerName)
	if err != nil {
		return err
	}
	err = docker.client.StartContainer(containerID, nil)
	if err != nil {
		return &models.ErrCannotStartContainer{cannotStartContainer + err.Error()}
	}
	fmt.Println("foi")
	return nil
}

func (docker *dockerRepository) StopContainer(containerName string) error {
	containerID, err := docker.GetContainerID(containerName)
	if err != nil {
		return err
	}
	err = docker.client.StopContainer(containerID, 10)
	if err != nil {
		return &models.ErrCannotStartContainer{cannotStartContainer + err.Error()}
	}
	return nil
}

func (docker *dockerRepository) GetContainerID(containerName string) (string, error) {
	var containerID string

	listOptions := gdocker.ListContainersOptions{
		All:     true,
		Filters: map[string][]string{dockerFilterOptionName: []string{containerName}},
	}

	containers, err := docker.client.ListContainers(listOptions)
	if err != nil {
		return "", &models.ErrCannotListContainers{err.Error()}
	}

	if ok := docker.isContainerListEmpty(containers); ok {
		return "", &models.ErrNoContainersFound{Message: containerNotFound}
	}

	//A container can habe multiple names, also the filters can return multiple containers
	//E.G:
	// Container 1 -> Name: production_sup
	// Container 2 -> Name: production_suplements
	// Container 3 -> Name: production_super
	// We must check if the required container is one of them.
	for _, container := range containers {
		for _, name := range container.Names {
			name = docker.removeSlashes(name)
			if name == containerName {
				containerID = container.ID
				break
			}
		}
	}

	return containerID, err
}

func (docker *dockerRepository) isContainerListEmpty(containers []gdocker.APIContainers) bool {
	return !(len(containers) != 0)
}

func (docker *dockerRepository) removeSlashes(containerName string) string {
	containerName = strings.Replace(containerName, "/", "", -1)
	return containerName
}
