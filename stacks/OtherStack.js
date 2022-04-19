import * as sst from "@serverless-stack/resources";

export default class OtherStack extends sst.Stack {
  constructor(scope, id, props) {
    super(scope, id, props);

    // Create the DynamoDB table
    const othersTable = new sst.Table(this, "Others", {
      fields: {
        Id: sst.TableFieldType.STRING,
        Title: sst.TableFieldType.STRING,
        Details: sst.TableFieldType.STRING,
      },
      primaryIndex: { partitionKey: "Id" },
    });



    // Create an API
    // const othersApi = new sst.Api(this, "OthersApi", {
    //   defaultFunctionProps: {
    //     // Pass in the table name to our API
    //     environment: {
    //       othersTableName: othersTable.dynamodbTable.tableName,
    //     },
    //   },
    //   routes: {
    //     "GET /": "src",
    //     "GET    /others":        "src/apps/others/handlers/list/list.go",
    //     "POST   /others":        "src/apps/others/handlers/create/create.go",
    //     "GET    /others/{id}":   "src/apps/others/handlers/get/get.go",
    //     "PUT    /others/{id}":   "src/apps/others/handlers/update/update.go",
    //     "DELETE /others/{id}":   "src/apps/others/handlers/delete/delete.go",
    //   },

    // });

    // Setup Notes API permissions to Notes Table
    // othersApi.attachPermissions([othersTable]);


    // Show the endpoint in the output
    this.addOutputs({
      othersTable: "Deployed"
      // OtherApiEndpoint: othersApi.url,
    });
  }
}
