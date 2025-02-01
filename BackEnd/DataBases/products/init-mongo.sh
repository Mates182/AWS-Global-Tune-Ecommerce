#!/bin/bash
mongoimport --host localhost --db globaltune_products --collection products --type json --file /data/products.json --jsonArray