package main

import (
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"encoding/json"

	Config "mainmodule/internal/apps/movies/config"
	"mainmodule/internal/apps/movies/services/movies"
)

func ListHandler(request *events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// EXAMPLE LIVE LAMBDA CHANGE!!
	// log.Info().Msg("hello world")
	// Load app configuration
	config := Config.Config{}.GetConfig()

	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	// Get list data from service!
	tableName := config.Stacks.Movies.Tablename
	s := movies.MovieService{TableName: tableName}
	result, err := s.FindList()

	if err != nil {
		log.Error().Err(err).Msg("Got error Finding Items result")
		return events.APIGatewayProxyResponse{StatusCode: http.StatusInternalServerError, Body: "Got error Finding Items result"}, nil
	}

	// Format results and return as API Gateway Response
	itemArrayJSON, err := json.Marshal(result)
	if err != nil {
		log.Error().Err(err).Msg("Got error marshalling result")
		return events.APIGatewayProxyResponse{StatusCode: http.StatusInternalServerError}, nil
	}

	// Send Response
	response := events.APIGatewayProxyResponse{Body: string(itemArrayJSON), StatusCode: http.StatusOK}
	response.Headers = make(map[string]string)
	response.Headers["Content-Type"] = "application/json"

	return response, nil
}

func main() {
	lambda.Start(ListHandler)
}
