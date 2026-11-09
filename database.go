package folcon

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var DB *sql.DB

// Connect MySql database initializes the MySQL database connection using .env variables
func ConnectMySql() error {

	if err := godotenv.Load(".env"); err != nil {
		log.Println("warning: .env file not found, using system environment variables")
	}

	dbUser := os.Getenv("MYSQL_DB_USER")
	dbPassword := os.Getenv("MYSQL_DB_PASSWORD")
	dbHost := os.Getenv("MYSQL_DB_HOST")
	dbPort := os.Getenv("MYSQL_DB_PORT")
	dbName := os.Getenv("MYSQL_DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbPassword, dbHost, dbPort, dbName)

	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}

	if err = DB.Ping(); err != nil {
		return err
	}

	log.Println("database connected successfully!")

	return nil

}
