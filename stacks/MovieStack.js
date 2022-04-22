import * as sst from "@serverless-stack/resources";

export default class Moviestack extends sst.Stack {
  constructor(scope, id, props) {
    super(scope, id, props);

    // Create the DynamoDB table
    const moviesTable = new sst.Table(this, "movies", {
      fields: {
        Id: sst.TableFieldType.STRING,
      },
      primaryIndex: { partitionKey: "Id" },
    });

    const moviesRootUrl = "cmd/handlers/movies/"
    
    // Create a HTTP API
    const moviesApi = new sst.Api(this, "moviesApi", {
      defaultFunctionProps: {
        srcPath: "backend/mainmodule",
        environment: {
          movieTableName: moviesTable.dynamodbTable.tableName,
        },
      },
      routes: {
        "GET    /":             moviesRootUrl+"/list/list.go",
        "GET    /movies":       "cmd/handlers/movies/list/list.go",
        "POST   /movies":       "cmd/handlers/movies/create/create.go",
        "GET    /movies/{id}":  "cmd/handlers/movies/get/get.go",
        "PUT    /movies/{id}":  "cmd/handlers/movies/update/update.go",
        "DELETE /movies/{id}":  "cmd/handlers/movies/delete/delete.go",
      },
    });

    // Setup API permissions to Notes Table
    moviesApi.attachPermissions([moviesTable]);

    // Show the endpoint in the output
    this.addOutputs({
      moviesApiEndpoint: moviesApi.url,
      moviesTablename: moviesTable.dynamodbTable.tableName,
    });
  }
}
