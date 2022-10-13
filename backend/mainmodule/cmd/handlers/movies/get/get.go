package main

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	
)

type ServiceMetadata struct {
	Name    string `json:"name"`
	Date    string `json:"date"`
	Version string `json:"version"`
	RunId   string `json:"runId"`
	Branch  string `json:"branch"`
}

func Handler(request events.APIGatewayV2HTTPRequest) (events.APIGatewayProxyResponse, error) {
	infoBytes, err := os.ReadFile("workflow_info.json")
	if err != nil {
		return Response(err, "Error reading workflow info file")
	}

	metadata := ServiceMetadata{}
	if err := json.Unmarshal(infoBytes, &metadata); err != nil {
		return Response(err, "Invalid workflow json")
	}

	resultJson, err := json.Marshal(metadata)
	if err != nil {
		return Response(err, "Error encoding workflow json")
	}

	return events.APIGatewayProxyResponse{
		Body:       string(resultJson),
		StatusCode: http.StatusOK,
		Headers:    map[string]string{"Content-Type": "application/json"},
	}, nil
}

func Response(err error, msg string) (events.APIGatewayProxyResponse, error) {
	metadata := ServiceMetadata{
		Name: "ctx-parts-metadata-service",
	}
	resultJson, _ := json.Marshal(metadata)
	return events.APIGatewayProxyResponse{
		Body:       string(resultJson),
		StatusCode: http.StatusOK,
		Headers:    map[string]string{"Content-Type": "application/json"},
	}, nil
}

func main() {
	lambda.Start(Handler)
}
