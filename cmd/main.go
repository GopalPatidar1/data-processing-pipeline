package main

import (
	"backend/config"
	"backend/routes"
	"encoding/json"
	"fmt"
	"net/http"
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
	pool := config.ConnectDB()
	defer pool.Close()

	// Pipeline Routes
	routes.RegisterPipelineRoutes(http.DefaultServeMux)

	fmt.Println("Server running on port 3007")

	http.ListenAndServe(":3007", nil)
}
