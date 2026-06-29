# ---- Build stage ----
FROM golang:1.25-alpine AS builder

WORKDIR /app

# Baixa as dependências primeiro para aproveitar o cache de camadas.
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Binário estático, sem dependências de C, para rodar numa imagem mínima.
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /app/server .

# ---- Runtime stage ----
FROM alpine:3.20

# Certificados para conexões TLS e um usuário não-root por segurança.
RUN apk add --no-cache ca-certificates && \
    adduser -D -u 10001 appuser

WORKDIR /app
COPY --from=builder /app/server /app/server

USER appuser

EXPOSE 8080

ENTRYPOINT ["/app/server"]
