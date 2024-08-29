package auth

import (
	"encoding/json"
	"hr/config"
	"hr/functions"
	"net/http"
)

func Signupuser (w http.ResponseWriter, r *http.Request) {
	var req struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Email     string `json:"email"`
		Password  string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Insert the user data into the users table
	_, err := config.Database.Exec(
		"INSERT INTO users (first_name, last_name, email, password, created_at) VALUES (?, ?, ?, ?, NOW())",
		req.FirstName, req.LastName, req.Email, req.Password,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"message": "User registered successfully",
	}
	functions.Response(w, response)
}
