package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// função principal
func main() {

	router := mux.NewRouter()

	router.HandleFunc("/transactions", GetTransactions).Methods("GET")
	router.HandleFunc("/authorizations", GetAuthorizations).Methods("GET")
	router.HandleFunc("/reversals", GetReversals).Methods("GET")
	router.HandleFunc("/reversals", PostReversal).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", router))
}

type Reversal struct {
	ID int `json:"id,omitempty"`
}

func GetTransactions(w http.ResponseWriter, r *http.Request) {
	dat, err := os.ReadFile("get_transactions.json")
	check(err)
	w.Write([]byte(dat))
}

func GetAuthorizations(w http.ResponseWriter, r *http.Request) {
	dat, err := os.ReadFile("get_authorizations.json")
	check(err)
	w.Write([]byte(dat))
}

func GetReversals(w http.ResponseWriter, r *http.Request) {
	dat, err := os.ReadFile("get_reversals.json")
	check(err)
	w.Write([]byte(dat))
}
func PostReversal(w http.ResponseWriter, r *http.Request) {
	dat := Reversal{ID: rand.Int()}
	json.NewEncoder(w).Encode(dat)
}
