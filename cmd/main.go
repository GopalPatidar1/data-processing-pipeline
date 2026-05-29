package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"backend/routes"
	"backend/config"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Go Server")
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	user := User{
		ID:   1,
		Name: "Gopal",
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(user)
}

func main() {
	// Connect DB
	conn := config.ConnectDB()

	defer conn.Close(context.Background())
	// User Routes
	routes.UserRoutes()

	// File Routes
	routes.FileRoutes()

	// Pipeline Routes
	routes.RegisterPipelineRoutes(http.DefaultServeMux)

	fmt.Println("Server running on port 3007")

	http.ListenAndServe(":3007", nil)
}