FROM golang:1.23-alpine AS builder

WORKDIR /app
COPY . .
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o bootstrap -ldflags="-s -w" cmd/server/main.go

FROM arm64v8/alpine:latest

WORKDIR /app

COPY --from=builder /app/bootstrap /app/bootstrap

# Set the timezone and install CA certificates
# RUN apk --no-cache add ca-certificates tzdata

EXPOSE 3000
EXPOSE 6060

ENTRYPOINT ["/app/bootstrap"]
