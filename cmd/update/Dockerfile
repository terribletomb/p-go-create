# Build the update service
FROM golang:1.24.1-alpine AS builder
WORKDIR /src
ENV CGO_ENABLED=0
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o /app/update ./cmd/update

FROM alpine:3.18
RUN apk add --no-cache ca-certificates
COPY --from=builder /app/update /update
EXPOSE 8083
RUN apt-get update && apt-get install -y curl

ENTRYPOINT ["/update"]
