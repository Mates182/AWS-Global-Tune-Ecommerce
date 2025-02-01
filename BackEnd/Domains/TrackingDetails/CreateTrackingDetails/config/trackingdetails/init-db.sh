#!/bin/bash
mongoimport --host localhost --db trackingdetails --collection trackingdetails --file /data/backup.json --jsonArray
echo "Database initialized successfully"