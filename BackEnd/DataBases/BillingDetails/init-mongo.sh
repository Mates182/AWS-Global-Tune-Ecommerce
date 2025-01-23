#!/bin/bash
mongoimport --host localhost --db billing-system --collection billing-details --type json --file /data/billing-details.json --jsonArray