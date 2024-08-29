package update

import (
	"encoding/json"
	"hr/config"
	"hr/functions"
	"net/http"
)

func UpdateUserDetails(w http.ResponseWriter, r *http.Request) {
	var req struct {
		User_ID   int `json:"user_id"`
		Phone     string `json:"phone_number"`
		DOB       string `json:"date_of_birth"`
		Address   string `json:"address"`
		Bio       string `json:"bio"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Prepare SQL query with the correct order
	query := `UPDATE users SET phone_number = ?, date_of_birth = ?, address = ?, bio = ? WHERE user_id = ?`
	// Execute SQL query with the correct parameter order
	_, err := config.Database.Exec(query, req.Phone, req.DOB, req.Address, req.Bio, req.User_ID)
	if err != nil {
		// Log the error for debugging
		http.Error(w, "Failed to update user details: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Send success response
	response := map[string]interface{}{
		"message": "User details updated successfully",
	}
	functions.Response(w, response)
}
