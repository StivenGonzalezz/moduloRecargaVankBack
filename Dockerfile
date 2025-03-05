# Usa una imagen base con Go instalado
FROM golang:1.21 AS builder

# Establece el directorio de trabajo
WORKDIR /app

# Copia los archivos del proyecto
COPY . .

# Descarga las dependencias y compila el binario
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app .

# Crea la imagen final mínima
FROM alpine:latest

# Establece el directorio de trabajo
WORKDIR /root/

# Copia el binario compilado desde la imagen anterior
COPY --from=builder /app/app .
RUN chmod +x ./app

# Expone el puerto correcto
EXPOSE 3000

# Comando para ejecutar la aplicación
CMD ["./app"]
