#!/bin/bash
mongoimport --host localhost --db tracking-system --collection tracking-details --type json --file /data/tracking-details.json --jsonArray