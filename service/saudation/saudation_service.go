package saudation

import (
	"github.com/thiagotrennepohl/alexa-skills/models"
	"github.com/thiagotrennepohl/alexa-skills/service"
)

type saudation struct {
}

//NewSaudationService creates a new saudation implementation of the Saudation interface
func NewSaudationService() service.Intent {
	return &saudation{}
}

func (svc *saudation) Response() (models.AlexaResponse, error) {
	response := models.AlexaResponse{}
	response.Response.OutputSpeech.Type = "SSML"
	response.Response.OutputSpeech.SSML = TDCSaudationMessage
	return response, nil
}
