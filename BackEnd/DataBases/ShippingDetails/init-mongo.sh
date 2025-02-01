#!/bin/bash
mongoimport --host localhost --db shipping-system --collection shipping-details --type json --file /data/shipping-details.json --jsonArray