package main

import (
	"log"
	"net/http"

	"github.com/andrepereira/card-transactions/models"

	"github.com/gorilla/mux"
)

// função principal
func main() {
	models.InitOperationTypesIDsTable()

	router := mux.NewRouter()
	router.HandleFunc("/accounts", models.CreateAccount).Methods("POST")
	router.HandleFunc("/accounts/{accountId}", models.GetAccount).Methods("GET")
	router.HandleFunc("/transactions", models.CreateTransaction).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", router))

}
