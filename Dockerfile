# Testing
# FROM docker:1.20.2-alpine3.17 AS testing
# WORKDIR /app
# COPY ["go.mod", "go.sum", "./"]
# RUN go mod download -x
# COPY . .
# RUN go test ./... -v -cover

# Stage 1: Builder
FROM golang:1.20-alpine3.17 as builder

RUN apk add --no-cache git upx

WORKDIR /app

COPY ["go.mod", "go.sum", "./"]

RUN go mod download -x

COPY . .

RUN go build -ldflags="-s -w" -o app ./cmd/main.go

# Stage 2: Producction
FROM alpine:3.17

RUN apk update && apk --no-cache add ca-certificates

WORKDIR /app

COPY --from=builder /app/app .

EXPOSE 3000

ENTRYPOINT [ "./app" ]
