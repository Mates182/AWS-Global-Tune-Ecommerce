const mysql = require("mysql2");

// Configuring the AWS RDS connection
const pool = mysql.createPool({
  host: "TU_ENDPOINT_RDS", // Switch to AWS RDS Endpoint
  user: "TU_USUARIO",
  password: "TU_PASSWORD",
  database: "Invoicing",
  waitForConnections: true,
  connectionLimit: 10,
  queueLimit: 0,
});

module.exports = pool.promise();

