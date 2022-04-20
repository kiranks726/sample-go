import * as sst from "@serverless-stack/resources";

export default class Moviestack extends sst.Stack {
  constructor(scope, id, props) {
    super(scope, id, props);

    // Create the DynamoDB table
    const moviesTable = new sst.Table(this, "movies", {
      fields: {
        Id: sst.TableFieldType.STRING
      },
      primaryIndex: { partitionKey: "Id" },
    });



    // Create a movies HTTP API
    const moviesApi = new sst.Api(this, "moviesApi", {
      defaultFunctionProps: {
        // Pass in the table name to our API
        environment: {
          movieTableName: moviesTable.dynamodbTable.tableName,
        },
      },
      routes: {
        "GET /": "src",
        "GET    /movies":        "src/apps/movies/handlers/list/list.go",
        "POST   /movies":        "src/apps/movies/handlers/create/create.go",
        "GET    /movies/{id}":   "src/apps/movies/handlers/get/get.go",
        "PUT    /movies/{id}":   "src/apps/movies/handlers/update/update.go",
        "DELETE /movies/{id}":   "src/apps/movies/handlers/delete/delete.go",
      },

    });

    // Setup Notes API permissions to Notes Table
    moviesApi.attachPermissions([moviesTable]);


    // Show the endpoint in the output
    this.addOutputs({
      moviesApiEndpoint: moviesApi.url,
      moviesTablename: moviesTable.dynamodbTable.tableName
    });
  }
}
