FROM golang:1.20.8-bookworm AS builder

WORKDIR /app

# go.mod go.sum etc
COPY go.* ./

RUN go mod download

# source code
COPY *.go ./

RUN go build -o /todowheel-backend

EXPOSE 8080

ENTRYPOINT ["/todowheel-backend"]
