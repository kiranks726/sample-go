package main

import (
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/rs/zerolog/log"

	Config "mainmodule/internal/apps/movies/config"
	"mainmodule/internal/apps/movies/services/movies"
)

type DeleteResponse struct {
	ID            string `json:"id"`
	StatusMessage string `json:"status"`
}

func Handler(request *events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// Getting id from path parameters
	itemID := request.PathParameters["id"]

	// Load app configuration
	config := Config.Config{}.GetConfig()

	// GetItem request
	s := movies.MovieService{TableName: config.Movies.TableName}
	err := s.DeleteOne(itemID)

	// Checking for errors, return error
	if err != nil {
		log.Error().Err(err).Msg("Got error Finding Items result")
		return events.APIGatewayProxyResponse{StatusCode: http.StatusBadRequest, Body: err.Error()}, nil
	}

	responseJSON, _ := json.Marshal(DeleteResponse{ID: itemID, StatusMessage: "DELETED"})

	return events.APIGatewayProxyResponse{StatusCode: http.StatusOK, Body: string(responseJSON)}, nil
}

func main() {
	lambda.Start(Handler)
}
