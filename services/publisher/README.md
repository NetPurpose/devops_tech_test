# Web stack server

This container provides the ingress for messages. It exposes three endpoints


1. GET /ping -> healthcheck, will respond with a pong message if this server is up
1. POST /send -> post a message to the "chat bot"
1. GET /receive -> receive a message from the chatbot, if any
