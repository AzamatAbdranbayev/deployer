FROM golang:1.17 AS builder
WORKDIR /app
COPY go.mod .
COPY go.sum .
COPY . .
RUN GO111MODULE="on" CGO_ENABLED=0 GOOS=linux GOARCH=amd64  go build -o app ./cmd

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/app .
COPY --from=builder /app/internal/config /app/internal/config
COPY --from=builder /usr/local/go/lib/time/zoneinfo.zip /
RUN apk update && apk add git
EXPOSE 3001
CMD ["./app"]