import * as sst from "@serverless-stack/resources";
import * as lambda from "aws-cdk-lib/aws-lambda";

export default class Moviestack extends sst.Stack {
  constructor(scope, id, props) {
    super(scope, id, props);

    // Create the DynamoDB table
    const moviesTable = new sst.Table(this, "movies", {
      fields: {
        Id: "string",
      },
      primaryIndex: { partitionKey: "Id" },
    });

    // Route Constants
    const MOVIES_API_ENDPOINT = " /movies"
    const MOVIES_ID_PATH = "/{id}"
    
    // Create a HTTP API
    const moviesApi = new sst.Api(this, "moviesApi", {
      defaults: {
        function: {
          environment: {
            STAGE: this.stage,
          }
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
    moviesApi.attachPermissions([
      moviesTable,
      "appconfig:GetLatestConfiguration",
      "appconfig:StartConfigurationSession",
    ]);

    // Show the endpoint in the output
    this.addOutputs({
      MoviesApiEndpoint: moviesApi.url,
      MoviesTablename: moviesTable.tableName,
    });
  }

  // Additional Methods
  resourceName(name) {
    return `${this.stage}-ctx-kitchensink-${name}`
  }
  
  movieRoute(name) {
    const layerArnX86 = "arn:aws:lambda:us-east-1:027255383542:layer:AWS-AppConfig-Extension:68"
    return {
      function: {
        srcPath: "backend/mainmodule",
        functionName: this.resourceName(`movies-${name}`),
        handler: "cmd/handlers/movies/" + name + "/" + name + ".go",
        layers: [
          lambda.LayerVersion.fromLayerVersionArn(this, name + "app-config-layer", layerArnX86),
        ]
      }
    }
  }
}
