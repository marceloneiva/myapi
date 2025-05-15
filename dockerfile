# Etapa de build
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Copia apenas os arquivos essenciais de dependência
COPY go.mod go.sum ./
# Atualiza os pacotes e instala o Git usando apk
RUN apk add --no-cache git
RUN go mod download

# Copia o restante do código-fonte
COPY . .

# Compila a aplicação
RUN go build -o myapi ./cmd/myapi

# Etapa final: imagem leve para produção
FROM alpine:latest

WORKDIR /root/

# Copia o binário gerado para a imagem final
COPY --from=builder /app/myapi .

# Instala certificados (necessário para conexões HTTPS)
RUN apk add --no-cache ca-certificates

EXPOSE 8080

# Comando padrão da aplicação
CMD ["./myapi"]

