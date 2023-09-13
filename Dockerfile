FROM golang:1.20.8-bookworm AS builder

WORKDIR /app

COPY . .

RUN go mod tidy
# RUN go mod download
RUN go build -o /todowheel-backend

EXPOSE 8080

ENTRYPOINT ["/todowheel-backend"]
