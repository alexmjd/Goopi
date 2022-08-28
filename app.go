package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"godb/src/models"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"

	"github.com/go-sql-driver/mysql"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

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
	var err error
	for loop {

		a.DB, err = sql.Open("mysql", config.FormatDSN())

		if err != nil {
			log.Fatal("Error while connecting to DB")
		}

		log.Print("SQL Connection openned...")

		pingErr := a.DB.Ping()
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

	// Get all variables from requests
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		fmt.Fprintf(w, "Cannot convert string to int.\n")
		return
	}

	p := models.Product{Id: id}

	err = p.GetProduct(a.DB)
	switch {
	case err == sql.ErrNoRows:
		fmt.Fprintf(w, "No product with id %d.\n", id)
		return
	case err != nil:
		fmt.Printf("Error here %s\n", err)
		fmt.Fprintf(w, "An error occured when getting the %d product.\n", id)
	}

	responseJson, err := json.Marshal(p)
	if err != nil {
		fmt.Fprintf(w, "An error occured: %s.\n", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s\n", responseJson)
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
