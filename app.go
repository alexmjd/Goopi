package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"godb/src/models"

	"github.com/gorilla/mux"

	"github.com/go-sql-driver/mysql"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

type Product models.Product

// Init the DB Connection
func (a *App) Init(user, pass, dbname string) {
	config := mysql.NewConfig()

	config.User = user
	config.Passwd = pass
	config.Net = "tcp"
	config.Addr = "godb:3306"
	config.DBName = dbname

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

// Run API Server
func (a *App) Run(addr string) {
	log.Println("Starting server...")
	router := mux.NewRouter().StrictSlash(true)

	log.Println("Server is listening.")

	router.HandleFunc("/", Hello)
	router.HandleFunc("/products/{id}", a.ReadProductById).Methods("GET")
	router.HandleFunc("/products", a.ReadAllProduct).Methods("GET")
	router.HandleFunc("/products", a.CreateProduct).Methods("POST")
	router.HandleFunc("/products/{id}", a.UpdateProduct).Methods("PATCH")
	router.HandleFunc("/products/{id}", a.DeleteProduct).Methods("DELETE")

	log.Fatal(http.ListenAndServe(addr, router))
}

func Hello(w http.ResponseWriter, r *http.Request) {
	log.Println("Hello API.")
}

func (a *App) ReadProductById(w http.ResponseWriter, r *http.Request) {
	log.Println("Read one product.")
}

func (a *App) ReadAllProduct(w http.ResponseWriter, r *http.Request) {
	log.Println("Read All product.")
}

func (a *App) CreateProduct(w http.ResponseWriter, r *http.Request) {
	log.Println("Create a product.")
}

func (a *App) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	log.Println("Update a product.")
}

func (a *App) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	log.Println("Delete a product.")
}
