

# Boilerplate of Golang micro-service with Kafka
The example of the GoLang project based on the microservices architecture which uses the Kafka Message Broker for communication

## Messaging Patterns

### Implement the Request-Reply pattern
In many cases, service need to handle real-time synchronous requests from client or other microservices to support application's user experience. With Request-Reply pattern, each will send request and wait for response from other services via different channels

![Request-Reply Pattern](./assets/RequestReplyPatternFigure.gif)








## References
- [Enterprise Integration Patterns: Request-Reply](https://www.enterpriseintegrationpatterns.com/RequestReply.html)