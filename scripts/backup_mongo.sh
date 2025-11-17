#!/usr/bin/env bash
set -euo pipefail
TIMESTAMP=$(date +%Y%m%d-%H%M)
BACKUP_DIR=./backups
mkdir -p "$BACKUP_DIR"

if command -v mongodump >/dev/null 2>&1; then
  echo "Ejecutando mongodump local..."
  mongodump --uri="${MONGO_URI}" --archive="${BACKUP_DIR}/backup-${TIMESTAMP}.gz" --gzip
  echo "Backup creado: ${BACKUP_DIR}/backup-${TIMESTAMP}.gz"
else
  echo "mongodump no encontrado localmente. Intentando ejecutar dentro del contenedor 'mongo'..."
  if docker ps --format '{{.Names}}' | grep -q '^mongo$'; then
    docker exec mongo mongodump --uri="${MONGO_URI}" --archive="/backups/backup-${TIMESTAMP}.gz" --gzip
    docker cp mongo:/backups/backup-${TIMESTAMP}.gz "${BACKUP_DIR}/backup-${TIMESTAMP}.gz"
    docker exec mongo rm -f "/backups/backup-${TIMESTAMP}.gz" || true
    echo "Backup creado: ${BACKUP_DIR}/backup-${TIMESTAMP}.gz"
  else
    echo "Contenedor 'mongo' no está corriendo y mongodump no está disponible. Abortando."
    exit 1
  fi
fi

### File: services/template/Dockerfile
# Multi-stage build genérico para copiar a cada servicio.
FROM golang:1.21-alpine AS builder
WORKDIR /src
ENV CGO_ENABLED=0
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o /app/service ./...

FROM alpine:3.18
RUN apk add --no-cache ca-certificates
COPY --from=builder /app/service /service
EXPOSE 8081
ENTRYPOINT ["/service"]