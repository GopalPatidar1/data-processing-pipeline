package routes

import (
	"encoding/json"
	"net/http"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	users := []User{
		{
			ID:    1,
			Name:  "Gopal",
			Email: "gopal@gmail.com",
		},
		{
			ID:    2,
			Name:  "Rahul",
			Email: "rahul@gmail.com",
		},
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func UserRoutes() {
	http.HandleFunc("/users", getUsers)
}
