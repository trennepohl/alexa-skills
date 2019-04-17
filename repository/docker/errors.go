package docker

type ErrCannotRestartClient struct {
	Message string
}

type ErrCannotCreateClient struct {
	Message string
}

type ErrCannotListContainers struct {
	Message string
}

type ErrNoContainersFound struct {
	Message string
}

type ErrCannotRemoveContainer struct {
	Message string
}

type ErrCannotPullImage struct {
	Message string
}

type ErrCannotCreateContainer struct {
	Message string
}

type ErrCannotStartContainer struct {
	Message string
}

func (e *ErrCannotRestartClient) Error() string {
	return e.Message
}

func (e *ErrNoContainersFound) Error() string {
	return e.Message
}

func (e *ErrCannotListContainers) Error() string {
	return e.Message
}

func (e *ErrCannotCreateClient) Error() string {
	return e.Message
}

func (e *ErrCannotRemoveContainer) Error() string {
	return e.Message
}

func (e *ErrCannotPullImage) Error() string {
	return e.Message
}

func (e *ErrCannotCreateContainer) Error() string {
	return e.Message
}

func (e *ErrCannotStartContainer) Error() string {
	return e.Message
}
