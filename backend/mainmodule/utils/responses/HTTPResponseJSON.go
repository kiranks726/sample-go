package responses

import (
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/rs/zerolog/log"
)

func SendHTTPResponseJSON(jsonData []byte, err error) (events.APIGatewayProxyResponse, error) {
	// return error if marshal error was passed into function
	if err != nil {
		log.Error().Err(err).Msg("Failed to UnmarshalMap result.Item")
		return events.APIGatewayProxyResponse{StatusCode: http.StatusInternalServerError}, nil
	}

	// return jsonData as response
	response := events.APIGatewayProxyResponse{Body: string(jsonData), StatusCode: http.StatusOK}
	response.Headers = make(map[string]string)
	response.Headers["Content-Type"] = "application/json"

	return response, nil
}
