FROM golang:1.22.5 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o /access ./cmd

COPY ./locale.yaml /app/locale.yaml

WORKDIR /

CMD ["/access"]