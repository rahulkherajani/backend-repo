package main

import (
	"fmt"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/rahulkherajani/backend-repo/controllers"
	"github.com/rahulkherajani/backend-repo/models"
)

func main() {
	godotenv.Load()

	handler := controllers.New()

	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: handler,
	}

	models.ConnectDatabase()

	fmt.Println("Server listening on http://localhost:8080")

	err := server.ListenAndServe()

	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
