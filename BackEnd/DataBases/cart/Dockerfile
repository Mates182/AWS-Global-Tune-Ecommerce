FROM redis:7.0-alpine

# Copiar el archivo de configuración desde el contexto de construcción al contenedor
COPY redis.conf /etc/redis/redis.conf

EXPOSE 6379

CMD ["redis-server", "/etc/redis/redis.conf"]
