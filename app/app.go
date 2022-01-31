package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func Start() {
	router := mux.NewRouter()

	router.HandleFunc("/createtoken", CreateToken).Methods(http.MethodPost)
	router.HandleFunc("/decodetoken", DecodeToken).Methods(http.MethodGet)

	address := getEnv("SERVER_ADDRESS", "localhost")
	port := getEnv("SERVER_PORT", "8181")
	log.Printf("Starting server on %s:%s ...", address, port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), router))
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
