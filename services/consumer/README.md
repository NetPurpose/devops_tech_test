# Asynchronous worker

This service consumes messages from it's `INBOUND` channel and generates a reply based on them, sending them back via the `OUTBOUND` channel

## ENVVARS
`AMQP_URL` -> an amqp connection string e.g. `amqp://<username>:<password>@<host>:<port>/`
`OUTBOUND` -> Name of RabbitMQ channel to send messages out on (should match to publisher `INBOUND`)
`INBOUND` -> Name of RabbitMQ channel to listen to replies on (should match publisher `OUTBOUND`)
