package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Couldn't get env var")
	}
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT env var is not set")
	}

	mux := http.NewServeMux()
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	mux.HandleFunc("GET /v1/healthz", handlerReadiness)
	mux.HandleFunc("GET /v1/err", handlerErr)

	log.Printf("Server starting on port: %s", port)
	log.Fatal(srv.ListenAndServe())
}
