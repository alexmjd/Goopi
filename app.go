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
	fmt.Fprintf(w, "Hello API.\n")
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
		return
	}

	responseJson, err := json.Marshal(p)
	if err != nil {
		fmt.Fprintf(w, "An error occured: %s.\n", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", responseJson)
}

func (a *App) ReadAllProduct(w http.ResponseWriter, r *http.Request) {
	log.Println("Read All product.")

	productList, err := models.GetAllProduct(a.DB)

	switch {
	case err == sql.ErrNoRows:
		fmt.Fprintf(w, "%s.\n", err)
		return
	case err != nil:
		fmt.Fprintf(w, "Error on requesting DB.\n")
		return
	}

	responseJson, err := json.Marshal(productList)
	if err != nil {
		fmt.Fprintf(w, "An error occured: %s.\n", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", responseJson)
}

func (a *App) CreateProduct(w http.ResponseWriter, r *http.Request) {
	log.Println("Create a product.")

	var p models.Product

	// Read the request's body content
	decodedContent := json.NewDecoder(r.Body)

	// Translate the decodedContent to a product
	err := decodedContent.Decode(&p)
	if err != nil {
		log.Printf("No readable content %s.\n", err)
		fmt.Fprintf(w, "Invalid request payload.\n")
		return
	}

	// Close the Request body at the end of the scope
	defer r.Body.Close()

	err = p.CreateProduct(a.DB)
	if err != nil {
		fmt.Fprintf(w, "Error while creating the product: %s.\n", err)
		return
	}

	responseJson, err := json.Marshal(p)
	if err != nil {
		fmt.Fprintf(w, "An error occured: %s.\n", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s\n", responseJson)
}

func (a *App) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	log.Println("Update a product.")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		fmt.Fprintf(w, "Invalid request: %s.\n", err)
		return
	}

	p := models.Product{Id: id}

	decodedContent := json.NewDecoder(r.Body)
	err = decodedContent.Decode(&p)
	if err != nil {
		fmt.Fprintf(w, "Error on payload: %s.\n", err)
		return
	}

	defer r.Body.Close()

	err = p.UpdateProduct(a.DB)
	if err != nil {
		fmt.Fprintf(w, "Error on updating product %d: %s.\n", p.Id, err)
		return
	}

	responseJson, err := json.Marshal(p)
	if err != nil {
		fmt.Fprintf(w, "An error occured: %s.\n", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", responseJson)

}

func (a *App) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	log.Println("Delete a product.")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		fmt.Fprint(w, "Invalid Product Id.\n")
		return
	}

	p := models.Product{Id: id}

	err = p.DeleteProduct(a.DB)
	if err != nil {
		fmt.Fprintf(w, "Error while deleting product %d.\n", p.Id)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Successfully delete product %d.\n", p.Id)
}
