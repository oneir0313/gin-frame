FROM golang:1.17.2 AS builder

WORKDIR /api
COPY . .

ENV GO111MODULE=on

WORKDIR /api/cmd
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o api

FROM alpine:latest
RUN apk update && \
    apk upgrade && \
    apk add --no-cache tzdata  && \
    apk add --no-cache ca-certificates && \
    apk add --no-cache curl && \
    rm -rf /var/cache/apk/*
ARG env
WORKDIR /api
COPY --from=builder /api/cmd/api /api/
COPY --from=builder /api/configs/config.yml /api/configs/config.yml

CMD ./api