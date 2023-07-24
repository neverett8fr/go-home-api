# syntax=docker/dockerfile:1

## Build
FROM golang:1.16-buster AS build

WORKDIR /home-service

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY config/*.yaml ./

COPY . .
COPY *.go ./

RUN go build -o /home-service

## Deploy
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /home-service ./

EXPOSE 8080

USER nonroot:nonroot


ENTRYPOINT ["/home-service"]