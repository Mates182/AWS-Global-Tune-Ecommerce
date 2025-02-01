const mysql = require("mysql2");

// Configuring the AWS RDS connection
const pool = mysql.createPool({
  host: "database-sql.c5eoehvxjn2a.us-east-1.rds.amazonaws.com", // Switch to AWS RDS Endpoint
  user: "globaltune",
  password: "globaltune202",
  database: "database-sql",
  waitForConnections: true,
  connectionLimit: 10,
  queueLimit: 0,
});

module.exports = pool.promise();

