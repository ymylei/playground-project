# playground-project
Services intended to act as testing and playground areas for various tools

## Rough Plan

Service Side:
- Rest API using `gin`
- Some sort of open source messaging lib to send

Backend Side:
- Some sort of open source messaging lib (F) to receive
- Redis for state sharing between the two sides?
- PostgresDB locally

Test Client Side:
- CronJob to issue random requests to the service?
- Some sort of pseudo web or mobile app to send/receive data from the service
- ??? Anything else interesting

Try to integrate:
- OAuth2 between the test client and service side minimally
- OpenTelemetry Based Traces/Metrics
- JSON logging via zerolog
- CI (Actions, golangci-lint, standard test framework, and go-vcr where appropriate, dependabot)
- K8s manifests to run locally on Docker K8s