package main

import (
	"api-go/Controller/Users"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/users", Users.UserControoler).Methods(http.MethodPost)

	fmt.Println("Server porta 5000")
	log.Fatal(http.ListenAndServe(":5000", router))
}
