# Establecer la imagen base para la etapa de construcción
FROM golang:1.15 AS builder

# Establecer el directorio de trabajo en el contenedor
WORKDIR /app

# Copiar los archivos go.mod y go.sum al directorio de trabajo
COPY go.mod go.sum ./

# Descargar las dependencias del módulo
RUN go mod download

# Copiar el directorio que contiene el código fuente de la aplicación
COPY . .

# Cambiar al directorio donde está el main.go y compilar la aplicación
# Compilar de forma estática para compatibilidad con Alpine
RUN cd cmd/api && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o main .

# Utilizar una imagen base ligera para la etapa de producción
FROM alpine:latest

# Instalar ca-certificates para permitir llamadas HTTPS
RUN apk --no-cache add ca-certificates

# Establecer el directorio de trabajo en el contenedor
WORKDIR /root/

# Copiar el ejecutable desde la etapa de construcción
COPY --from=builder /app/cmd/api/main .

# Asegurarse de que el ejecutable tiene permisos para ejecutarse
RUN chmod +x main

# Exponer el puerto en el que la aplicación estará escuchando
EXPOSE 8080

# Comando para ejecutar la aplicación
CMD ["./main"]
