package resume

import (
	"encoding/json"
	"hr/config"
	"hr/functions"
	"net/http"
)

func Skills(w http.ResponseWriter, r *http.Request) {
	var req struct {
		User_Id     int `json:"user_id"`
		Skill_Name  string `json:"skill_name"`
		Proficiency string `json:"proficiency_level"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Insert the user data into the users table
	_, err := config.Database.Exec(
		"INSERT INTO skills (user_id, skill_name, proficiency_level) VALUES (?, ?, ?)",
		req.User_Id, req.Skill_Name, req.Proficiency,)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"message": "Skills Added Successfully",
	}
	functions.Response(w, response)
}
