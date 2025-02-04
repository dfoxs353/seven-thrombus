#build
FROM golang:alpine AS builder

WORKDIR /go/src/app

RUN apk update --no-cache

COPY . .

RUN go install github.com/swaggo/swag/cmd/swag@v1.8.12

RUN go mod tidy

RUN go generate

RUN go build -ldflags="-s -w" -o /go/bin/app main.go 

# start
FROM alpine:latest

ENV TZ=Europe/Moscow

RUN apk update --no-cache \
    && apk add --no-cache \
        tzdata

COPY --from=builder /go/src/app/config.yml /config.yml
COPY --from=builder /go/bin/app /app

ENTRYPOINT /app

EXPOSE 8080