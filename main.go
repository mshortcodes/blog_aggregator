package main

import (
	"blogagg/internal/database"
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Couldn't get env var")
	}
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT env var is not set")
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL env var is not set")
	}

	dbConn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Can't connect to database:", err)
	}

	if err != nil {
		log.Fatal("Can't create to db connection:", err)
	}

	apiCfg := apiConfig{
		DB: database.New(dbConn),
	}

	mux := http.NewServeMux()
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	mux.HandleFunc("GET /v1/healthz", handlerReadiness)
	mux.HandleFunc("GET /v1/err", handlerErr)

	mux.HandleFunc("GET /v1/users", apiCfg.handlerGetUser)
	mux.HandleFunc("POST /v1/users", apiCfg.handlerCreateUser)

	log.Printf("Server starting on port: %s", port)
	log.Fatal(srv.ListenAndServe())
}
