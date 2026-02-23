# GO Application Testing

A basic Golang application to use for testing various DevOps tooling used within applications.

Goals:
 - VictoriaMetrics for application
 - Tracing
 - ECS Deployment
 - Terraform deployment
 - Security
 - E2E testing
 - CI/CD


## Observability 

The observability stack currently uses the following tools: 

- VictoriaMetrics (Metrics collection)
- Grafana (Visualisation)
- Grafana Tempo (Tracing)
- Loki (Logging)
- Profiling (Pyroscope)

# Setup / Running

The Makefile should have most of the commands required for building and running the application.

## Image Build

`Make image-build` will build the image file with the latest tag.

## Docker run

`Make docker-up` will run the application and other components required to run the application as well.

## Access

The ingress for the application can be accessed on `http://localhost:8080/` 

The observability stack can be accessed from the Grafana frontend on `http://localhost:3000/`