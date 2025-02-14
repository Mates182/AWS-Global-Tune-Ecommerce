# Usar la imagen oficial de Node.js basada en Alpine
FROM node:lts-alpine

# Definir variable de entorno
ENV NODE_ENV=production

# Definir el directorio de trabajo dentro del contenedor
WORKDIR /app

# Instalar dependencias necesarias para compilar algunos paquetes
RUN apk add --no-cache python3 make g++

# Copiar archivos de dependencias y luego instalarlas
COPY package*.json ./
RUN npm ci --only=production

# Copiar el resto de los archivos al contenedor
COPY . .

# Exponer el puerto 3000
EXPOSE 3000

# Comando para iniciar la aplicación
CMD ["npm", "run", "dev"]
