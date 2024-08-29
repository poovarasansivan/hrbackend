package resume

import (
	"encoding/json"
	"hr/config"
	"hr/functions"
	"net/http"
)

func Education(w http.ResponseWriter, r *http.Request) {
	var req struct {
		User_Id          int    `json:"user_id"`
		Institution_Name string `json:"institution_name"`
		Degree           string `json:"degree"`
		Field            string `json:"field_of_study"`
		Start_Date       string `json:"start_date"`
		End_Date         string `json:"end_date"`
		Grade            string `json:"grade"`
		Description      string `json:"description"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Insert the user data into the education table
	_, err := config.Database.Exec(
		"INSERT INTO education (user_id, institution_name, degree, field_of_study, start_date, end_date, grade, description) VALUES (?, ?, ?, ?, ?, ?, ?, ?)",
		req.User_Id, req.Institution_Name, req.Degree, req.Field, req.Start_Date, req.End_Date, req.Grade, req.Description,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"message": "Education Added Successfully",
	}
	functions.Response(w, response)
}
