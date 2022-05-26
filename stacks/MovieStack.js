import * as sst from "@serverless-stack/resources";

export default class Moviestack extends sst.Stack {
  constructor(scope, id, props) {
    super(scope, id, props);

    // Create the DynamoDB table
    const moviesTable = new sst.Table(this, "newmovies", {
      fields: {
        Id: sst.TableFieldType.STRING,
      },
      primaryIndex: { partitionKey: "Id" },
    });

    // Route Constants
    const MOVIES_API_ENDPOINT = " /movies"
    const MOVIES_ID_PATH = "/{id}"
    
    // Create a HTTP API
    const moviesApi = new sst.Api(this, "moviesApi", {
      defaultFunctionProps: {
        environment: {
          movieTableName: moviesTable.dynamodbTable.tableName,
        },
      },
      routes: {
        ["GET"    + MOVIES_API_ENDPOINT]:           this.movieRoute("list"),
        ["POST"   + MOVIES_API_ENDPOINT]:           this.movieRoute("create"),
        ["GET"    + MOVIES_API_ENDPOINT + MOVIES_ID_PATH]: this.movieRoute("get"),
        ["PUT"    + MOVIES_API_ENDPOINT + MOVIES_ID_PATH]: this.movieRoute("update"),
        ["DELETE" + MOVIES_API_ENDPOINT + MOVIES_ID_PATH]: this.movieRoute("delete"),

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

  // Additional Methods
  resourceName(name) {
    return `${this.stage}-ctx-kitchensink-${name}`
  }
  
  movieRoute(name) {
    return {
      function: {
        srcPath: "backend/mainmodule",
        functionName: this.resourceName(`movies-${name}`),
        handler: "cmd/handlers/movies/" + name + "/" + name + ".go",
      }
    }
  }
}
