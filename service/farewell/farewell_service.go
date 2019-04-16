package farewell

import (
	"github.com/thiagotrennepohl/alexa-skills/models"
	"github.com/thiagotrennepohl/alexa-skills/service"
)

const farewellMessage = `
<speak> 
		Adeus <lang xml:lang="en-US"> Gophers!</lang> O código desta talk estará no<lang xml:lang="en-US"> Git</lang>.
		<p> Espero que vocês tenham gostado </p>
</speak>`

type farewell struct {
}

func NewFarewellService() service.Intent {
	return &farewell{}
}

func (svc *farewell) Response() (models.AlexaResponse, error) {
	alexaResponse := models.AlexaResponse{}
	alexaResponse.Response.OutputSpeech.Type = "SSML"
	alexaResponse.Response.OutputSpeech.SSML = farewellMessage
	alexaResponse.ShouldEndSession = true
	return alexaResponse, nil
}
