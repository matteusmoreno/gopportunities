FROM golang:1.24.2

# Define o diretório de trabalho
WORKDIR /go/src/app

# Copia os arquivos do projeto
COPY go.mod go.sum ./
RUN go mod download

# Copia o restante dos arquivos depois do download (evita rebuilds desnecessários)
COPY . .

# Expõe a porta da aplicação
EXPOSE 8090

# Compila o projeto
RUN go build -o main cmd/main.go

# Executa o binário
CMD ["./main"]
