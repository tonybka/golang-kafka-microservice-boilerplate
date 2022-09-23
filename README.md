

# Boilerplate of Golang micro-service with Kafka
The GoLang boilerplate that implements best practices for great foundation of event-sourcing Microservice with Kafka

## Features
- [x] **Command Line Options (CLI)**: Provides options to interact with golang application

## Design Patterns & Principles
### Clean Architecture
<img src="assets/CleanArchitectureFigure.png" width="350" height="270" />

Decouples Golang web service to independent layers, the changes from one layer won't affect codes of the others.
### Messaging Patterns

#### Implement the Request-Reply pattern
In many cases, service needs to handle real-time synchronous requests from client or other microservices to support application's user experience. With Request-Reply pattern, each service will send request and wait for response from other services via different channels, in this case are different Kafka topics

![Request-Reply Pattern](./assets/RequestReplyPatternFigure.gif)


## Highlighted Dependencies
- Build CLI with Cobra: [github.com/spf13/cobra](github.com/spf13/cobra)
## References
- [Enterprise Integration Patterns: Request-Reply](https://www.enterpriseintegrationpatterns.com/RequestReply.html)