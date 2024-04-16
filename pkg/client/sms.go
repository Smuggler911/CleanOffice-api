package client

import (
	"CleanOffice/config"
	"fmt"
	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
)

func SendSmsToClient(clientPhone string, body string) {

	env, _ := config.LoadConfig()
	accountsid := env.TwilioAccountSid
	authToken := env.TwilioAuthToke
	client := twilio.NewRestClientWithParams(twilio.ClientParams{Username: accountsid, Password: authToken})

	params := &openapi.CreateMessageParams{}

	senderPhone := env.TwilioPhoneNumber
	params.SetTo(clientPhone)
	params.SetFrom(senderPhone)
	params.SetBody(body)

	_, err := client.Api.CreateMessage(params)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("SMS sent successfully!")
	}
}
