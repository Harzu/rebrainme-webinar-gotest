## Repo for golang test webinar

use `make help` for info about run

```shell
help                 List all available targets with help
init                 Init project
proto-gen            Generate protobuf schema
generate             Golang codegen
build                Build docker containers
dev-migration-up     Up all migrations to local database
dev-migration-down   Down all migrations to local database
run                  Run services in develop mode
stop                 Stop develop environment
test-short           Run only short(unit) tests
test-long            Run all tests(unit/integration) with docker
```