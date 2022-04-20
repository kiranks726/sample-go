package main

import (
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/kiranks726/sample-go/src/apps/todos/config"
	"github.com/kiranks726/sample-go/src/apps/todos/services/todos"

	"encoding/json"
	"fmt"
)

func ListHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// Load app configuration
	config := config.Config{}.GetConfig()

	// Get list data from service
	tableName := config.Todos.TableName
	ts := todos.TodoService{TableName: tableName}

	result, err := ts.FindTodoItems()
	if err != nil {
		fmt.Println("Got error Finding Items result: ", err.Error())
		return events.APIGatewayProxyResponse{StatusCode: 500}, nil
	}

	// Format results and return as API Gateway Response
	itemArrayJson, err := json.Marshal(result)
	if err != nil {
		fmt.Println("Got error marshalling result: ", err.Error())
		return events.APIGatewayProxyResponse{StatusCode: 500}, nil
	}

	response := events.APIGatewayProxyResponse{Body: string(itemArrayJson), StatusCode: http.StatusOK}
	response.Headers = make(map[string]string)
	response.Headers["Content-Type"] = "application/json"
	return response, nil
}

func main() {
	lambda.Start(ListHandler)
}
