FROM mongo:latest

WORKDIR /data

COPY shipping-details.json /data/shipping-details.json

COPY init-mongo.sh /docker-entrypoint-initdb.d/init-mongo.sh

RUN chmod +x /docker-entrypoint-initdb.d/init-mongo.sh

EXPOSE 27017