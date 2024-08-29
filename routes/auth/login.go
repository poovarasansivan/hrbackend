package auth

import (
	"database/sql"
	"encoding/json"
	"hr/config"
	"hr/functions"
	"net/http"
)

type UserInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserAccount struct {
	User_Id int    `json:"user_id`
	Email   string `json:"email"`
}

func Loginuser(w http.ResponseWriter, r *http.Request) {
	var input UserInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		response := map[string]interface{}{
			"success": false,
			"error":   "Invalid Request",
		}
		functions.Response(w, response)
		return
	}

	var data UserAccount
	err := config.Database.QueryRow("SELECT user_id, email FROM users WHERE email = ? AND password = ?", input.Email, input.Password).Scan(&data.User_Id, &data.Email)

	var response map[string]interface{}
	if err != nil {
		if err == sql.ErrNoRows {
			response = map[string]interface{}{
				"success": false,
				"error":   "No User Found",
			}
			w.WriteHeader(http.StatusUnauthorized) // Set status code for unauthorized access
		} else {
			response = map[string]interface{}{
				"success": false,
				"error":   err.Error(),
			}
			w.WriteHeader(http.StatusInternalServerError) // Internal error
		}
		functions.Response(w, response)
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
