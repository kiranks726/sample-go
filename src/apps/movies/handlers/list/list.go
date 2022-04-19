package main

import (
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	Config "github.com/kiranks726/sample-go/src/apps/movies/config"
	"github.com/kiranks726/sample-go/src/apps/movies/services/movies"
	"github.com/rs/zerolog/log"

	"encoding/json"
)

func ListHandler(request *events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// Load app configuration
	config := Config.Config{}.GetConfig()

	// Get list data from service
	tableName := config.Movies.TableName
	s := movies.MovieService{TableName: tableName}
	result, err := s.FindList()

	if err != nil {
		log.Error().Err(err).Msg("Got error Finding Items result")
		return events.APIGatewayProxyResponse{StatusCode: http.StatusInternalServerError}, nil
	}

	// Format results and return as API Gateway Response
	itemArrayJSON, err := json.Marshal(result)
	if err != nil {
		log.Error().Err(err).Msg("Got error marshalling result")
		return events.APIGatewayProxyResponse{StatusCode: http.StatusInternalServerError}, nil
	}

	response := events.APIGatewayProxyResponse{Body: string(itemArrayJSON), StatusCode: http.StatusOK}
	response.Headers = make(map[string]string)
	response.Headers["Content-Type"] = "application/json"

	return response, nil
}

func main() {
	lambda.Start(ListHandler)
}
