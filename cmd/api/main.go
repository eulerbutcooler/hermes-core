package main

import (
	"log"
	"net/http"

	"github.com/eulerbutcooler/hermes-core/internal/api"
	"github.com/eulerbutcooler/hermes-core/internal/config"
	"github.com/eulerbutcooler/hermes-core/internal/db"
	"github.com/eulerbutcooler/hermes-core/internal/store"
)

func main() {
	cfg := config.LoadConfig()

	pool, err := db.New(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Fatal DB Error: %v", err)
	}
	defer pool.Close()

	relayStore := store.NewRelayStore(pool)
	handler := api.NewHandler(relayStore)
	router := api.NewRouter(handler)

	log.Printf("Core running on port %s", cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, router); err != nil {
		log.Fatal(err)
	}
}
