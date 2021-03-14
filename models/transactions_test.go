package models

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestCreateAccount(t *testing.T) {
	var jsonStr = []byte(`{"document_number": "12345678900"}`)

	req, err := http.NewRequest("POST", "/accounts", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateAccount)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	// Add a new line char at final
	expected := `"success"` + "\n"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestCreateAccountWithMalformedJson1(t *testing.T) {
	var jsonStr = []byte(`{"Xdocument_number": "12345678900"}`)

	req, err := http.NewRequest("POST", "/accounts", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateAccount)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}

	// Check the response body is what we expect.
	// Add a new line char at final
	expected := `"error: invalid JSON post"` + "\n"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestCreateAccountWithMalformedJson2(t *testing.T) {
	var jsonStr = []byte(`{"document_number": 12345678900}`)

	req, err := http.NewRequest("POST", "/accounts", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateAccount)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}

	// Check the response body is what we expect.
	// Add a new line char at final
	expected := `"error: invalid JSON post"` + "\n"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
func TestGetAccount(t *testing.T) {

	r, _ := http.NewRequest("GET", "/accounts/1", nil)
	w := httptest.NewRecorder()

	//Hack to try to fake gorilla/mux vars
	vars := map[string]string{
		"accountId": "1",
	}

	r = mux.SetURLVars(r, vars)

	GetAccount(w, r)

	if status := w.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	// Add a new line char at final
	expected := `{"account_id":1,"document_number":"12345678900"}` + "\n"
	if w.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			w.Body.String(), expected)
	}

}

func TestGetAccountWithUnknowAccount(t *testing.T) {

	r, _ := http.NewRequest("GET", "/accounts/a", nil)
	w := httptest.NewRecorder()

	//Hack to try to fake gorilla/mux vars
	vars := map[string]string{
		"accountId": "a",
	}

	r = mux.SetURLVars(r, vars)

	GetAccount(w, r)

	if status := w.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	// Add a new line char at final
	expected := `{}` + "\n"
	if w.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			w.Body.String(), expected)
	}

}
func TestCreateTransaction(t *testing.T) {
	InitOperationTypesIDsTable()
	var jsonStr = []byte(`{"account_id":1,"operation_type_id":4,"amount":123.45}`)
	req, err := http.NewRequest("POST", "/transactions", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateTransaction)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	// Add a new line char at final
	expected := `"success"` + "\n"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestCreateTransactionWithUnknowAccount(t *testing.T) {
	InitOperationTypesIDsTable()
	var jsonStr = []byte(`{"account_id":2,"operation_type_id":4,"amount":123.45}`)
	req, err := http.NewRequest("POST", "/transactions", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateTransaction)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusNotAcceptable {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNotAcceptable)
	}

	// Check the response body is what we expect.
	// Add a new line char at final
	expected := `"error: invalid Account ID, Operation Type ID or Amount field"` + "\n"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
