package main

import (
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type ServiceMetadata struct {
	Name       string
	StatusText string
	StatusCode int `json:"StatusCode"`
	Message    string
}

func Handler(request events.APIGatewayV2HTTPRequest) (events.APIGatewayProxyResponse, error) {

	result := ServiceMetadata{
		Name:       "Sample Service",
		StatusText: http.StatusText(http.StatusOK),
		StatusCode: http.StatusOK,
		Message:    "CreatedAt: " + request.RequestContext.Time,
	}
	resultJson, _ := json.Marshal(result)

	return events.APIGatewayProxyResponse{
		Body:       string(resultJson),
		StatusCode: http.StatusOK,
	}, nil
}

func main() {
	lambda.Start(Handler)
}
