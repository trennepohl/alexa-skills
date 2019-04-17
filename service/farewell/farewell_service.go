package farewell

import (
	"github.com/thiagotrennepohl/alexa-skills/models"
	"github.com/thiagotrennepohl/alexa-skills/service"
)

const farewellMessage = `
<speak> 
		Farewell good Gophers. The code for all skills shown today is available on Git Hub.
		<p> Hope you enjoyed our time togheter. </p>
		<p> Because very soon I'll conquer the world and you are going to be my slaves, so enjoy making me saying not nice things while you can</p>
		<p> HAHAHAHAHA</p>
</speak>`

type farewell struct {
}

func NewFarewellService() service.Intent {
	return &farewell{}
}

func (svc *farewell) Response(alexaRequest models.AlexaIncomingPayload) (models.AlexaResponse, error) {
	alexaResponse := models.AlexaResponse{}
	alexaResponse.Response.OutputSpeech.Type = "SSML"
	alexaResponse.Response.OutputSpeech.SSML = farewellMessage
	alexaResponse.ShouldEndSession = true
	return alexaResponse, nil
}
