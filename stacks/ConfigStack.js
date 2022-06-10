import * as sst from "@serverless-stack/resources";
import * as appconfig from "aws-cdk-lib/aws-appconfig";
import * as fs from "fs";
export default class Configstack extends sst.Stack {
    constructor(scope, id, props) {
        super(scope, id, props);

        const cfnApplication = new appconfig.CfnApplication(this, "CtxKitchenSinkApplication", {
            name: `${this.stage}-ctx-kitchensink-go-application`,
            description: "CTX Kitchen Sink Stack Configuration",
          });

          const cfnEnvironment = new appconfig.CfnEnvironment(this, "CtxKitchenSinkEnvironment", {
            applicationId: cfnApplication.ref,
            name: this.stage,
            description: `Environment for ${this.stage}`,
            // tags: [{ key: projectKey, value: project, }],
          });

          const cfnConfigurationProfile = new appconfig.CfnConfigurationProfile(this, "CtxKitchenSinkConfigurationProfile", {
            applicationId: cfnApplication.ref,
            locationUri: "hosted",
            name: "stacks-config",
            description: "CTX stacks configuration",
            // tags: [{ key: projectKey, value: project, }],
          });

          const cfnDeploymentStrategy = new appconfig.CfnDeploymentStrategy(this, "CtxKitchenSinkDeploymentStrategy", {
            deploymentDurationInMinutes: 0,
            growthFactor: 100,
            name: "all-at-once",
            replicateTo: "NONE",
          
            // the properties below are optional
            description: "CTX configuration deployment strategy",
            finalBakeTimeInMinutes: 1,
            growthType: "LINEAR",
            // tags: [{ key: projectKey, value: project, }],
          });

          const content = JSON.parse(fs.readFileSync(`./config/local-${this.stage}.json`).toString());
          const cfnHostedConfigurationVersion = new appconfig.CfnHostedConfigurationVersion(this, "CtxKitchenSinkHostedConfigurationVersion", {
            applicationId: cfnApplication.ref,
            configurationProfileId: cfnConfigurationProfile.ref,
            content: JSON.stringify(content),
            contentType: "application/json",
          
            // the properties below are optional
            description: "CTX configuration version",
            latestVersionNumber: 1,
          });

          new appconfig.CfnDeployment(this, "CtxDeployment", {
            applicationId: cfnApplication.ref,
            configurationProfileId: cfnConfigurationProfile.ref,
            configurationVersion: cfnHostedConfigurationVersion.ref,
            deploymentStrategyId: cfnDeploymentStrategy.ref,
            environmentId: cfnEnvironment.ref,
          
            // the properties below are optional
            description: "Ctx Config Deployment",
            // tags: [{ key: projectKey, value: project, }],
          });

          this.addOutputs({
            ConfigApplication: cfnApplication.name,
            ConfigProfile: cfnConfigurationProfile.name,
          });      
    }

}