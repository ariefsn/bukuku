FROM golang:alpine

RUN apk add build-base

RUN go version

ENV PORT="3003"

RUN mkdir /app

WORKDIR /app

# COPY go.mod .

# COPY go.sum .

# RUN go mod download

ADD . /app

RUN go mod tidy

RUN go build -o main

CMD [ "/app/main" ]