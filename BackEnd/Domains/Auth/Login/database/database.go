package database

import (
	"database/sql"
	"fmt"
	"log"
	"login-service/secrets"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db   *sql.DB
	once sync.Once
)

func GetDBInstance() *sql.DB {
	once.Do(func() {
		dbHost, dbUser, dbPassword, dbPort, dbName := secrets.GetDBCredentials()

		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)

		var err error
		db, err = sql.Open("mysql", dsn)
		if err != nil {
			log.Fatalf("Failed to connect to database: %v", err)
		}

		// Test the connection
		err = db.Ping()
		if err != nil {
			log.Fatalf("Database connection error: %v", err)
		}

		log.Println("Connected to RDS successfully")
	})

	return db
}
