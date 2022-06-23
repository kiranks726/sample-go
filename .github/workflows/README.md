# Overview

CTX pipelines for serverless is a continuous integration and continuous delivery (CI/CD) platform that automates build, test, and deployment pipeline of CTX serverless applications and infrastructure

## Usage

- Service generated from service template should be pushed to main branch as initial commit and tagged as  version 0.0.0. This is baseline tag.
  
- Open a pull request from any feature branch into `main` branch to trigger the CI(build.yml) pipeline workflow.

- Once pull request is merged to `main` branch, it triggers deployment to dev(deploy_dev.yml) pipeline workflow. This has two variants
  - If version is unchanged in VERSION.ini file, then `validation of version` fails, but `deployment to   dev` still succeeds which allows developers to iterate on dev until the need to deploy to QA environment.
  - Incrementing version in VERSION.ini file will result in `creation of tag in git` provided both `deployment to dev` and `validation of version succeeds`.

- Go to `Actions` tab in the github repository and click on the `Deployment` Wokflow. Now click on `Run workflow` and choose the `Tags` to use the Workflow from and then choose the `Environment` to deploy the application. This triggers the (deploy_qa_prod.yml) pipeline workflow which deploys from the chosen tag.

>Note: To know about structure, approach, functionalities of this pipeline, refer [CTX pipelines for serverless applications](https://dev.azure.com/adi-ctx/Operations/_wiki/wikis/Operations.wiki/137/CTX-pipelines-for-serverless-applications)
