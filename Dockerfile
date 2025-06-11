# Etapa de build
FROM golang:1.23.1

# Diretório de trabalho dentro do container
WORKDIR /app

# Copia os arquivos para dentro do container
COPY . .

# Instala dependências do Go
RUN go mod tidy

# Compila o binário (opcional, dependendo do seu projeto)
# RUN go build -o main .

# (opcional) etapa final: use uma imagem mais enxuta
# FROM alpine:latest
# COPY --from=builder /app/main /app/main
# CMD ["/app/main"]
