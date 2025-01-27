FROM golang:1.23-alpine AS builder

WORKDIR /app
COPY . .
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o bootstrap -ldflags="-s -w" cmd/server/main.go

FROM arm64v8/alpine:latest

WORKDIR /app

COPY --from=builder /app/bootstrap /app/bootstrap

RUN ulimit -n 1048576

ENTRYPOINT ["/app/bootstrap"]
