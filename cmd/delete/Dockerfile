# Build the delete service
FROM golang:1.24.1-alpine AS builder
WORKDIR /src
ENV CGO_ENABLED=0
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o /app/delete ./cmd/delete

FROM alpine:3.18
RUN apk add --no-cache ca-certificates
COPY --from=builder /app/delete /delete
EXPOSE 8084
RUN apt-get update && apt-get install -y curl

ENTRYPOINT ["/delete"]
