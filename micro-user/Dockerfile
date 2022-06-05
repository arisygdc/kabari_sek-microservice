FROM golang:1.18.1-alpine3.14 AS stage
WORKDIR /app

COPY . .

RUN go build -o micro-user cmd/server.go

FROM alpine:3.15
WORKDIR /app

COPY --from=stage /app/micro-user .
COPY --from=stage /app/config.yaml .

ENTRYPOINT [ "./micro-user" ]