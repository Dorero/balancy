FROM golang:1.20.4-alpine

WORKDIR /app

COPY . .

RUN apk add make && apk add postgresql-client

CMD make run