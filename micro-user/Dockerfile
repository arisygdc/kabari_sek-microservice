ARG path=chat-in-app_microservices/micro-user

FROM golang:1.18.1-alpine3.14 AS stage

WORKDIR /app/${path}

COPY . .

RUN apk add build-base
RUN go build -o micro-user cmd/server.go

FROM alpine:3.15
WORKDIR /app

COPY --from=stage /app/${path}/micro-user .
COPY --from=stage /app/${path}/config.yaml .

ENTRYPOINT [ "./micro-user" ]