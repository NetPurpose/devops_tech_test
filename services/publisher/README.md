# Web server

This container provides the ingress for messages. It exposes three endpoints


1. GET /ping -> healthcheck, will respond with a pong message if this server is up, and the status of it's connection to RabbitMQ
1. POST /send -> post a message to the "chat bot"
1. GET /receive -> receive a message from the chatbot, if any. an empty queue informs as such


## ENVVARS
`AMQP_URL` -> an amqp connection string e.g. `amqp://<username>:<password>@<host>:<port>/`
`OUTBOUND` -> Name of RabbitMQ channel to send messages out on (should match to consumer `INBOUND`)
`INBOUND` -> Name of RabbitMQ channel to listen to replies on (should match consumer `OUTBOUND`)
