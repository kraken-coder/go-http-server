package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var data []string = []string{}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/test", test)

	router.HandleFunc("/add/{item}", addItem).Methods("GET", "DELETE")

	err := http.ListenAndServe(":4000", router)

	if err != nil {
		log.Fatal(err)
	}

}

func test(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(struct {
		ID string
	}{"5555"})
}

func addItem(w http.ResponseWriter, r *http.Request) {
	rootData := mux.Vars(r)["item"]

	data = append(data, rootData)

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(data)
}
