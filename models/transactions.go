package models

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// GetAccount Receive Account ID param by URI
// Return Account ID and Document Number as JSON response
func GetAccount(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	for _, item := range Accounts {
		accountID, _ := strconv.Atoi(params["accountId"])
		if item.AccountID == accountID {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(&Account{})
}

// CreateAccount Receive Document Number in a JSON by Post
// Return a JSON with "success" or "error"
func CreateAccount(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var account Account
	err := json.NewDecoder(r.Body).Decode(&account)

	if err != nil || account.DocumentNumber == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "invalid JSON post"}`))

	} else {

		size := len(Accounts)
		// The next Account Id is the length of slice (Accounts DB) + 1
		account.AccountID = size + 1

		Accounts = append(Accounts, account)
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"success": "account created"}`))

	}

}

// CreateTransaction Receive Account ID, Operation Type ID and a Amount
// in a JSON by Post
// Return a JSON with "success" or "error"
func CreateTransaction(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var transaction Transaction
	err := json.NewDecoder(r.Body).Decode(&transaction)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "invalid JSON post"}`))

	} else {

		size := len(Transactions)
		// The next Transaction Id is the length of slice (Transactions DB) + 1
		transaction.TransactionID = size + 1

		// Verify if Account ID exists
		existsAccountID := false
		for _, value := range Accounts {
			if value.AccountID == transaction.AccountID {
				existsAccountID = true
				break
			}
		}

		// Verify if Amount exists in request
		existsAmount := false
		if transaction.Amount != 0 {
			existsAmount = true
		}

		// Verify if Operation Type ID exists
		existsOperationTypeID := false
		for _, value := range OperationTypes {
			if value.OperationTypeID == transaction.OperationTypeID {
				existsOperationTypeID = true
				break
			}
		}

		for i, v := range Accounts {
			if v.AccountID == transaction.AccountID {
				if transaction.OperationTypeID == 1 || transaction.OperationTypeID == 2 ||
					transaction.OperationTypeID == 3 {
					if v.AvaliableCreditLimit-float32(transaction.Amount) < 0 {
						w.WriteHeader(http.StatusNotAcceptable)
						w.Write([]byte(`{"error": "exceeded credit limit"}`))
					}
					Accounts[i].AvaliableCreditLimit = v.AvaliableCreditLimit - float32(transaction.Amount)

				}

				if transaction.OperationTypeID == 4 {
					Accounts[i].AvaliableCreditLimit = v.AvaliableCreditLimit + float32(transaction.Amount)

				}

			}
		}

		if existsAccountID && existsOperationTypeID && existsAmount {
			// If Transaction Type ID is 1, 2 or 3 make ammount negative
			if transaction.OperationTypeID == 1 || transaction.OperationTypeID == 2 || transaction.OperationTypeID == 3 {
				transaction.Amount = transaction.Amount * (-1)
			}

			transaction.EventDate = time.Now()

			Transactions = append(Transactions, transaction)
			w.WriteHeader(http.StatusCreated)
			w.Write([]byte(`{"success": "transaction created"}`))
		} else {
			w.WriteHeader(http.StatusNotAcceptable)
			w.Write([]byte(`{"error": "invalid Account ID, Operation Type ID or Amount field"}`))
		}
	}

}
