# Usa una imagen base con Go instalado
FROM golang:1.21 AS builder

# Establece el directorio de trabajo
WORKDIR /app

# Copia los archivos del proyecto
COPY . .

# Descarga las dependencias y compila el binario
RUN go mod tidy
RUN go build -o app .

# Crea la imagen final mínima
FROM alpine:latest

# Establece el directorio de trabajo
WORKDIR /root/

# Copia el binario compilado desde la imagen anterior
COPY --from=builder /app/app .

# Expone el puerto en el que corre la aplicación
EXPOSE 8080

# Comando para ejecutar la aplicación
CMD ["./app"]
