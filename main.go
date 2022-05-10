package main

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/go-resty/resty/v2"
	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
)

type EventInput struct {
	RequesterEmail string `json:"email"`
	ToNumber       string `json:"to"`
}

func HandleRequest(ctx context.Context, inp EventInput) (string, error) {
	accountSid := os.Getenv("TWILIO_SID")
	authToken := os.Getenv("TWILIO_TOKEN")

	reqUrl := fmt.Sprintf("%s/user/%s", os.Getenv("ALPHA_URL"), inp.RequesterEmail)

	client := resty.New()
	res, err := client.R().EnableTrace().Get(reqUrl)

	if err != nil {
		return "", err
	}

	twilioClient := twilio.NewRestClientWithParams(twilio.ClientParams{AccountSid: accountSid, Password: authToken})
	twilioParams := &openapi.CreateMessageParams{}
	twilioParams.SetTo(inp.ToNumber)
	twilioParams.SetBody("Hello from Go!")
	twilioParams.SetFrom(os.Getenv("TWILIO_NUMBER"))

	twilioRes, err := twilioClient.ApiV2010.CreateMessage(twilioParams)

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Sent Message. Got %s from alpha! Got %s from twilio!", res.Body(), *twilioRes.Body), nil
}

func main() {
	lambda.Start(HandleRequest)
}
