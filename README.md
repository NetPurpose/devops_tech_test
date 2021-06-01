# Instructions

The goal of this technical test is to take the dummy application (two go-lang based microservices) and "deploy" them in a local stack.

The aim of the local stack is to allow the hypothetical developer to easily spin up a local environment for testing.
Bonus points for also demonstrating ability to instrument things such as:-
* monitoring tools (e.g. prometheus) 
* advanced load balancers (traefik, Kong etc)
* a readiness check
* anything else you think might be useful/ helpful for developers working on this "application"

The application consists of:

1. A web server `/services/publisher`
1. An asynchronous worker `/services/consumer`

(more details on each can be found in their folders)

The two applications will talk to each other via a rabbitmq server. Once successfully configured it should be possible to send an http `POST` request to the webserver at `/send` with a JSON payload of `{"body": <some message>}` and afterwards receive a message from
`/receive` via a `GET` request which will have content with your original message embedded in it.



Some technologies which may be of interest (not all of them need to be used!):

- docker-compose
- minikube
- skaffold
- ansible
- puppet
- chef
- docker-swarm
