package service

import "github.com/thiagotrennepohl/alexa-skills/models"

//Saudation is responsible for all user greetings
type Saudation interface {
	SayTDCHello() models.AlexaResponse
}
