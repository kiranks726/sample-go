# Sample-Go

## Overview

Sample Services in Go. This project is meant to showcase common patterns and best practices for application development in Go.

## Features

Features of this project

* [x] Scaffold of AWS Lambda in Go Pattern
* [x] REST API
* [x] Linting
* [x] Centralized Configuration
* [x] Project Structure
* [x] Go Modules / Packages
* [x] Error handling 
* [x] Logging (verify in AWS), abstract zerolog
* [ ] Test framework
* [ ] Concurrency patterns, use import data command
* [ ] Package Versioning
* [ ] Add Dependendency Injection for main handler pattern (Service, Config)
* [ ] Define non-Lambda Web Framework like Gin or Fiber service pattern
* [ ] Host Web Framework as container in AWS 
* [ ] REST: Open API Schema validation
* [ ] Service structure with Controller
* [ ] gRPC API
* [ ] gRPC Schema validation
* [ ] WebSocket API
* [ ] Pipelines



# Setup

This project was bootstrapped with [Create Serverless Stack](https://docs.serverless-stack.com/packages/create-serverless-stack).

1. Start by installing the dependencies.

    ```bash
    $ npm install
    ```

2. Setup proper AWS profile either in `~/.aws/credentials`. This app is setup for `593186514662_AWSAdministratorAccess` profile name. If this changes make sure to update in the first line of the `Makefile`

3. To start with local development type `make start`. This will deploy a development stack in AWS and connect local machine to the stack for development. You can edit the files in `src` live. NOTE: The first deployment will take a while, but then it is pretty fast.
4. To deploy a build to AWS, run `make deploy`. This will deploy the stack to the AWS account prefixing the stack resources with the "dev" stage name. NOTE: The stage name can be changed if needed. The "dev" name helps to separate the "local" and "dev" resources in the target AWS account.
4. Type `make` or `make help` for additioanl details on commands to get started.


## Commands

### Make Commands

Use `make` for project specific commands.

<span style="color:firebrick">ATTENTION: The make commands use the AWS profile `593186514662_AWSAdministratorAccess` for proper provisioning. This should be updated to the AWS profile that is used. Change in the first line of the `Makefile`</span>

#### `make help`

This will output a list of commands similar to the `npm` commands below, but lead with the `make` command.



## Documentation

Learn more about the Serverless Stack.
- [Docs](https://docs.serverless-stack.com)
- [@serverless-stack/cli](https://docs.serverless-stack.com/packages/cli)
- [@serverless-stack/resources](https://docs.serverless-stack.com/packages/resources)
Error Handling
- [Err with if (error)](https://www.bacancytechnology.com/blog/golang-error-handling) - it feels wrong buyt it is on purpose
- https://earthly.dev/blog/golang-errors/


## Community

[Follow us on Twitter](https://twitter.com/ServerlessStack) or [post on our forums](https://discourse.serverless-stack.com).


## Reference

* [How to create a REST API in Golang with serverless](https://serverless-stack.com/examples/how-to-create-a-rest-api-in-golang-with-serverless.html) 
* [Golang Monorepo Patterns](https://earthly.dev/blog/golang-monorepo/)
