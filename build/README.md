# `/build`

Packaging and Continuous Integration.

## `/build/ci`
Put your CI (github-actions, azure-pipeline, travis, jenkins) configurations and scripts in the `/build/ci` directory. Note that some of the CI tools (e.g., Travis CI) are very picky about the location of their config files. Try putting the config files in the `/build/ci` directory linking them to the location where the CI tools expect them when possible (don't worry if it's not and if keeping those files in the root directory makes your life easier :-)).


## `build/package`
Put your cloud (AMI), container (Docker), OS (deb, rpm, pkg) package configurations and scripts in the `/build/package` directory.


## `build/orchestration`
IaaS, PaaS, system and container orchestration deployment configurations and templates (docker-compose, kubernetes/helm, mesos, terraform, bosh).


## NOTE: `.build` (dot build)

This is a directory for build meta files, not the actual build. Actual build files are generally stored in the hidden `.build` directory. 


## Examples:

* https://github.com/cockroachdb/cockroach/tree/master/build
