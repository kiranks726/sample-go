import * as sst from "@serverless-stack/resources";

export default class TodoStack extends sst.Stack {
  constructor(scope, id, props) {
    super(scope, id, props);

    // Create the DynamoDB table
    const todosTable = new sst.Table(this, "Todos", {
      fields: {
        Id: sst.TableFieldType.STRING,
        Title: sst.TableFieldType.STRING,
        Details: sst.TableFieldType.STRING,
      },
      primaryIndex: { partitionKey: "Id" },
    });



    // Create a Todos HTTP API
    const todosApi = new sst.Api(this, "TodosApi", {
      defaultFunctionProps: {
        // Pass in the table name to our API
        environment: {
          todoTableName: todosTable.dynamodbTable.tableName,
        },
      },
      routes: {
        "GET /": "src",
        "GET    /todos":        "src/apps/todos/handlers/list/list.go",
        "POST   /todos":        "src/apps/todos/handlers/create/create.go",
        "GET    /todos/{id}":   "src/apps/todos/handlers/get/get.go",
        "PUT    /todos/{id}":   "src/apps/todos/handlers/update/update.go",
        "DELETE /todos/{id}":   "src/apps/todos/handlers/delete/delete.go",
      },

    });

    // Setup Notes API permissions to Notes Table
    todosApi.attachPermissions([todosTable]);


    // Show the endpoint in the output
    this.addOutputs({
      TodosApiEndpoint: todosApi.url,
      TodosTablename: todosTable.dynamodbTable.tableName
    });
  }
}
