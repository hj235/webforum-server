package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var DBInstance *sql.DB

func InitialiseDB() {
	// Verify DSN
	// rootPw := os.Getenv("MYSQL_ROOT_PASSWORD")
	// port := os.Getenv("MYSQL_PORT")
	// database := os.Getenv("MYSQL_DATABASE")
	// dsn := fmt.Sprintf("root:%v@tcp(mysql-db:%v)/%v", rootPw, port, database)
	dsn := os.Getenv("MYSQL_DSN")
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("An error has occurred with the data source name.\nOpening a driver will typically not attempt to connect to the database.")
		log.Fatal(err)
	}
	DBInstance = db
	log.Println("Opened MySql driver")
	fmt.Println("Connected to: " + dsn)
	Ping()
	log.Println("Connection to database established")
}

func Ping() {
	err := DBInstance.Ping()
	if err != nil {
		log.Println("An error has occurred establishing connection to the database")
	}
}

func GetDB() *sql.DB {
	if DBInstance == nil {
		InitialiseDB()
	}
	return DBInstance
}
