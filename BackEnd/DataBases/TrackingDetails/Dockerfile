FROM mongo:latest

WORKDIR /data

COPY billing-details.json /data/billing-details.json

COPY init-mongo.sh /docker-entrypoint-initdb.d/init-mongo.sh

RUN chmod +x /docker-entrypoint-initdb.d/init-mongo.sh

EXPOSE 27017