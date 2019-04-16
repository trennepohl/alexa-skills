package service

import "github.com/thiagotrennepohl/alexa-skills/models"

type Intent interface {
	Response() (models.AlexaResponse, error)
}
