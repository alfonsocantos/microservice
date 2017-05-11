package main

import (
	"fmt"
	handler "handler"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {

	port := os.Getenv("HOST")

	if port == "" {
		log.Println("HOST environment variable not setted, using default 8080")
		port = "8080"
	}

	fmt.Println("Listening on port ", port)

	r := mux.NewRouter()
	r.HandleFunc("/healthz", handler.Healthz).Methods("GET")

	log.Fatal(http.ListenAndServe(":"+port, r))

}
