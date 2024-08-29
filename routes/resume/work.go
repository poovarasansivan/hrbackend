package resume

import (
	"encoding/json"
	"hr/config"
	"hr/functions"
	"net/http"
)

func WorkExperience(w http.ResponseWriter, r *http.Request) {
	var req struct {
		User_Id          int `json:"user_id"`
		Job_title        string `json:"job_title"`
		Company          string `json:"company_name"`
		Location         string `json:"location"`
		Start_Date       string `json:"start_date"`
		End_Date         string `json:"end_date"`
		Responsibilities string `json:"responsibilities"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Insert the user data into the users table
	_, err := config.Database.Exec(
		"INSERT INTO work_experience (user_id, job_title, company_name, location, start_date, end_date, responsibilities) VALUES (?, ?, ?, ?, ?, ?, ?)",
		req.User_Id, req.Job_title, req.Company, req.Location, req.Start_Date, req.End_Date, req.Responsibilities,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"message": "Work Experience Added Successfully",
	}
	functions.Response(w, response)
}
