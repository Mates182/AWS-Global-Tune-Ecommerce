FROM mongo:latest

WORKDIR /data

COPY backup.json /data/backup.json

COPY init-db.sh /docker-entrypoint-initdb.d/init-db.sh

RUN chmod +x /docker-entrypoint-initdb.d/init-db.sh

EXPOSE 27017