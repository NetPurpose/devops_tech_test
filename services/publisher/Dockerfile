FROM golang:latest as build-env
WORKDIR /np-app

COPY go.* ./
RUN go mod download
COPY . ./

RUN go build -mod=readonly -o server

FROM debian:buster

EXPOSE 8080
WORKDIR /

COPY --from=build-env /np-app/server /server

CMD ["/server"]


