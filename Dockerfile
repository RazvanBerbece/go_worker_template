# syntax=docker/dockerfile:1

# Base Image
FROM golang:latest

WORKDIR /app

# Install project dependencies
RUN apt-get update
COPY go.mod go.sum ./
RUN go mod download

COPY . ./
COPY ./.env ./.env

RUN CGO_ENABLED=0 GOOS=linux go build -o build/worker/main ./src/cmd/main.go 

CMD sh -c "./build/worker/main"