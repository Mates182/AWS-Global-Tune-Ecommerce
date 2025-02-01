const { Sequelize } = require('sequelize');

const sequelize = new Sequelize('database_name', 'username', 'password', {
  host: 'localhost',
  dialect: 'mysql', // Changes depending on the database engine
});

(async () => {
  try {
    await sequelize.authenticate();
    console.log('Database connection successful.');
  } catch (error) {
    console.error('Could not connect to the database:', error);
  }
})();

module.exports = sequelize;



//for mongo

const mongoose = require('mongoose');

const connectDB = async () => {
  try {
    await mongoose.connect('mongodb://localhost:27017/database_name', {
      useNewUrlParser: true,
      useUnifiedTopology: true,
    });
    console.log('Database connection successful.');
  } catch (error) {
    console.error('Could not connect to the database:', error);
    process.exit(1);
  }
};

module.exports = connectDB;



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
