ARG path=chat-in-app_microservices/api-gateway

FROM golang:1.18-alpine3.15 AS stage

WORKDIR /app/${path}
COPY . .

RUN go build -o api-gateway ./cmd/server.go

FROM alpine:3.15

WORKDIR /app

COPY --from=stage /app/${path}/api-gateway .
COPY --from=stage /app/${path}/config.yaml .

ENTRYPOINT [ "./api-gateway" ]