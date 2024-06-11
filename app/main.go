package main

import (
	"log/slog"
	"net/http"

	petstore_api "github.com/tetsurowada/backend-golang-echo/petstore/delivery"
)

func main() {
	// Create service instance.
	service := petstore_api.NewPetsService()
	// Create generated server.
	srv, err := petstore_api.NewServer(&service)
	if err != nil {
		panic(err)
	}
	slog.Info("server started...")
	if err := http.ListenAndServe(":8080", srv); err != nil {
		slog.Error("server error: %v", err)
	}
}
