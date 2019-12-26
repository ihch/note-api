FROM golang:1.13.5-alpine3.10

WORKDIR /api
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY src ./src
