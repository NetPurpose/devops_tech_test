FROM golang:latest as build-env
WORKDIR /np-app

COPY go.* ./
RUN go mod download
COPY . ./

RUN go build -mod=readonly -o listener

FROM debian:buster

WORKDIR /

COPY --from=build-env /np-app/listener /listener

CMD ["/listener"]


