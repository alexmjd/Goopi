package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {

	// Use of mysql.NewConfig() to take all default parameters such as mysql_native_password
	config := mysql.NewConfig()

	config.User = os.Getenv("DBUSER")
	config.Passwd = os.Getenv("DBPASS")
	config.Net = "tcp"
	config.Addr = "godb:3306"
	config.DBName = "godb"

	// Loop the DB ping until it works.
	loop := true
	for loop {

		db, err := sql.Open("mysql", config.FormatDSN())

		if err != nil {
			log.Fatal("Error while connecting to DB")
		}

		log.Print("SQL Connection openned...")

		pingErr := db.Ping()
		if pingErr != nil {
			log.Println("Ping Error: ", pingErr)
		}

		if err == nil && pingErr == nil {
			loop = false
		}

		time.Sleep(time.Second * 3)
	}

	fmt.Println("Hello db")
}
