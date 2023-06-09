# Imagem base com o Go instalado
FROM golang:latest

# Define o diretório de trabalho dentro do contêiner
WORKDIR /application

# Copia o código fonte para o diretório de trabalho
COPY . .

# Compila o código fonte
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./main.go

# Define o comando de inicialização do contêiner
CMD ["./main"]
