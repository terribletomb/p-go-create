# ---------- Etapa 1: Compilación ----------
    ARG SERVICE=./cmd/create
    FROM golang:1.24.1-alpine AS builder
    
    WORKDIR /app
    
    ENV CGO_ENABLED=0

    # Copiar dependencias y descargar módulos
    COPY go.mod go.sum ./
    RUN go mod download
    
    # Copiar el resto del código fuente
    COPY . .
    
    # Compilar el binario con soporte CGO (por defecto en Alpine)
    RUN go build -o /app/service $SERVICE

    # ---------- Etapa 2: Imagen final optimizada ----------
    FROM alpine:latest

    # Label de la imagen
    LABEL maintainer="danysoftdev" \
          version="0.1" \
          description="Imagen optimizada para la aplicación Go"
    
    RUN apk add --no-cache ca-certificates

    # Crear un usuario no-root para seguridad
    RUN adduser -D gouser
    
    # Definir directorio de trabajo
    WORKDIR /home/gouser/app
    
    # Copiar el binario desde la etapa anterior
    COPY --from=builder /app/service /service

    # Cambiar al usuario seguro
    USER gouser

    EXPOSE 8081

    # Ejecutar el binario
    ENTRYPOINT ["/service"]
