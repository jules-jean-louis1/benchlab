# Étape 1 : Build
FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
# Compilation statique de l'application
RUN CGO_ENABLED=0 GOOS=linux go build -o grpc-server .

# Étape 2 : Image finale très légère
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/grpc-server .
# On expose le port gRPC (par défaut 50051 en général)
EXPOSE 50051
CMD ["./grpc-server"]