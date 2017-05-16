package main

import (
	healthz "healthz"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

var version string

type stati struct {
	key, value string
}

func main() {

	log.Println("Version: ", version)

	port := os.Getenv("PORT")
	host := os.Getenv("HOST")

	if port == "" {
		log.Println("PORT environment variable not set, using default 8080")
		port = "8080"
	}

	if host == "" {
		log.Println("HOST environment variable not set, using default 127.0.0.1")
		host = "127.0.0.1"
	}

	log.Printf("%s %v:%v", "Listening on", host, port)

	r := mux.NewRouter()
	r.HandleFunc("/healthz", healthz.Healthz).Methods("GET")

	log.Fatal(http.ListenAndServe(host+":"+port, r))

}

func initService(c chan stati) {
	// Do init stuff

	// Connect to database
	_dbStatus := stati{key: "mongodb", value: "ok"}
	c <- _dbStatus

	// status is finnally ok
	_status := stati{key: "status", value: "ok"}
	c <- _status

}
