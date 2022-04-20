// TodoService to manage request for TodoItem data
package todos

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

// TodoService struct to contains features
// related to TodoItem data management
type TodoService struct {
	Name      string
	TableName string
}

func (ps TodoService) DoTest() []string {
	list := []string{"hello", "ninja", "monkey"}
	return list
}

// FindTodoItems to get TodoItems
// Return an array of TodoItems
func (ps TodoService) FindTodoItems() ([]TodoItem, error) {

	// Creating session for client
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create DynamoDB client
	svc := dynamodb.New(sess)

	// Build the query input parameters
	params := &dynamodb.ScanInput{
		TableName: &ps.TableName,
	}

	// Scan table
	result, err := svc.Scan(params)

	// Checking for errors, return error
	if err != nil {
		fmt.Println("Query API call failed: ", err.Error())
		return []TodoItem{}, err
	}

	// Setup resut data
	var itemArray []TodoItem
	for _, i := range result.Items {
		item := TodoItem{}

		err := dynamodbattribute.UnmarshalMap(i, &item)

		if err != nil {
			fmt.Println("DynamoDB UnmarshalError")
			fmt.Println(err)
		}

		itemArray = append(itemArray, item)
	}
	return itemArray, err
}
