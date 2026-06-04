package main

import (
	"backend/config"
	"backend/routes"
	"fmt"
	"net/http"
)

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func ApiMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		origin := r.Header.Get("Origin")

		if origin != "https://mywebsite.com" {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	pool := config.ConnectDB()
	defer pool.Close()

	// Pipeline Routes
	mux := http.NewServeMux()

	routes.RegisterPipelineRoutes(mux)

	fmt.Println("Server running on port 3007")

	// apiMiddlewareMux := ApiMiddleware(mux) if want to add API middleware
	// server := corsMiddleware(apiMiddlewareMux)

	server := corsMiddleware(mux)

	http.ListenAndServe(":3007", server)
}
