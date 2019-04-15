package service

import "github.com/thiagotrennepohl/alexa-containers/models"

//Saudation is responsible for all user greetings
type Saudation interface {
	SayTDCHello() models.AlexaResponse
}
