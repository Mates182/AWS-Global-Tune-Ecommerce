#!/bin/bash
mongoimport --host localhost --db billingdetails --collection billingdetails --file /data/backup.json --jsonArray
echo "Database initialized successfully"