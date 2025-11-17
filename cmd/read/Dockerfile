# Build the read service
FROM golang:1.24.1-alpine AS builder
WORKDIR /src
ENV CGO_ENABLED=0
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o /app/read ./cmd/read

FROM alpine:3.18
RUN apk add --no-cache ca-certificates
COPY --from=builder /app/read /read
EXPOSE 8082
ENTRYPOINT ["/read"]
