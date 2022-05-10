package main

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/go-resty/resty/v2"
)

type EventInput struct {
	Email string `json:"email"`
}

func HandleRequest(ctx context.Context, inp EventInput) (string, error) {
	reqUrl := fmt.Sprintf("%s/user/%s", os.Getenv("ALPHA_URL"), inp.Email)

	client := resty.New()
	res, err := client.R().EnableTrace().Get(reqUrl)

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Hello %s!", res.Body()), nil
}

func main() {
	lambda.Start(HandleRequest)
}
