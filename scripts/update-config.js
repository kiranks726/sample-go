const AWS = require("aws-sdk");
const fs = require("fs");

async function main(options) {
    console.log(`Updating config for ${options.project} ...`);
    let cf = new AWS.CloudFormation();
    let resp = await cf.describeStacks({}).promise();
    let outputs = resp.Stacks
        .filter(stack => stack.StackName.startsWith(options.project))
        .filter(stack => stack.Outputs && stack.Outputs.length)
        .map(stack => {
            let info = stack.StackId.split(":");
            return Object.assign(
                {
                    Name: stack.StackName.substring(options.project.length + 1),
                    Region: info[3],
                    AccountId: info[4]
                },
                stack.Outputs.reduce((map, output) => {
                    map[output.OutputKey] = output.OutputValue; return map;
                }, {})
            );
        })
        .reduce((map, output) => {
            map[output.Name] = output;
            return map;
        }, {});

    let config = JSON.parse(fs.readFileSync("config/default.json"));
    let deploy = JSON.parse(fs.readFileSync("config/deploy.json"));
    config = Object.assign(config, deploy, {stacks: outputs});
    fs.writeFileSync(`config/local-${options.stage}.json`, JSON.stringify(config, null, 2));
}

if (process.argv.length < 4) {
    console.error("Usage: update-config (stage) (project)")
} else {
    main({
        stage: process.argv[2],
        project: process.argv[3]
    });
}
