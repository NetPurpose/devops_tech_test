# Instructions

Your goal here is to build some tooling to allow a new developer to relatively quickly an easilly run the multi-component dummy application contained within this repo.

The application consists of:

a client ui  (`/frontend`)
a backend server (`/publisher`)
rabbitmq
an asynchronous worker (`/consumer`)


The expected flow of messages is that the client ui connects to the server via a websocket and can then send messages to the server. The server will then pass those messages on to the asynchronous worker via rabbitmq.

The worker will then respond with a very unimaginative reply, which will also be sent to rabbitmq and back to the client ui via the backend server.


There are a few points to consider:

The microservices need to be told where each other is
This is for local testing


Some technologies which may be of interest (not all of them need to be used!):

- docker-compose
- minikube
- skaffold
- ansible
- puppet
- chef
- docker-swarm
