package repository

type DockerRepository interface {
	RestartContainer(containerName string) error
	Connect(endpoint string) error
	ListContainers() ([]string, error)
	StopContainer(containerName string) error
	StartContainer(containerName string) error
}
