package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/proxy", ImageProxy).Methods("GET")

	fmt.Println("Successfully started server")

	log.Fatal(http.ListenAndServe(":80", r))
}
