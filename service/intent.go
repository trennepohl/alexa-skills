package service

import "github.com/thiagotrennepohl/alexa-containers/models"

type Intent interface {
	Response() (models.AlexaResponse, error)
}
