FROM --platform=linux/amd64 golang:1.22

WORKDIR /app

# Copiar todo al WORKDIR
COPY . .

# Instalar dependencias
RUN go mod tidy

# Construir la aplicación
RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-backend ./cmd 

# Añadir el script de espera al contenedor
COPY wait-for-it.sh /wait-for-it.sh
RUN chmod +x /wait-for-it.sh

EXPOSE 80

# Usar el script de espera para esperar a MySQL y luego iniciar la aplicación
# TODO: Thi mysql:3306 should't be hardcoded
CMD ["/wait-for-it.sh", "mysql:3306", "--", "/docker-backend"]
