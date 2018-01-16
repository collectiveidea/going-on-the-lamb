package main

import (
	"errors"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"log"
)

var (
	// ErrNameNotProvided is thrown when a name is not provided
	ErrNameNotProvided = errors.New("No name was provided in the HTTP Body")
)

// Handler is your Lambda function handler
func Handler(request events.APIGatewayProxyRequest) (response events.APIGatewayProxyResponse, err error) {
	log.Printf("Processing Lambda request %s\n", request.RequestContext.RequestID)

	if len(request.Body) < 1 {
		response = events.APIGatewayProxyResponse{
			Body: "",
		}
		err = ErrNameNotProvided

		return
	}

	response = events.APIGatewayProxyResponse{
		Body:       "Hello " + request.Body,
		StatusCode: 200,
	}

	return
}

func main() {
	lambda.Start(Handler)
}
