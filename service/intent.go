package service

import "github.com/thiagotrennepohl/alexa-skills/models"

type Intent interface {
	Response(requestBody models.AlexaIncomingPayload) (models.AlexaResponse, error)
}
