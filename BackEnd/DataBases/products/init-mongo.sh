#!/bin/bash
mongoimport --host localhost --db products --collection products --type json --file /data/products.json --jsonArray