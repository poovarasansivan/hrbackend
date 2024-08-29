package resume

import (
	"encoding/json"
	"hr/config"
	"hr/functions"
	"net/http"
)

func GetUserDetails(w http.ResponseWriter, r *http.Request) {
	var req struct {
		User_Id int `json:"user_id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userID := req.User_Id

	// Retrieve User Personal Details
	userRow := config.Database.QueryRow(
		"SELECT first_name, last_name, email, phone_number, date_of_birth, address, bio FROM users WHERE user_id = ?",
		userID,
	)

	var userDetails struct {
		FirstName   string `json:"first_name"`
		LastName    string `json:"last_name"`
		Email       string `json:"email"`
		PhoneNumber string `json:"phone_number"`
		DateOfBirth string `json:"date_of_birth"`
		Address     string `json:"address"`
		Bio         string `json:"bio"`
	}

	if err := userRow.Scan(&userDetails.FirstName, &userDetails.LastName, &userDetails.Email, &userDetails.PhoneNumber, &userDetails.DateOfBirth, &userDetails.Address, &userDetails.Bio); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Retrieve Education Details
	educationRows, err := config.Database.Query(
		"SELECT institution_name, degree, field_of_study, start_date, end_date, grade, description FROM education WHERE user_id = ?",
		userID,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer educationRows.Close()

	var educationDetails []map[string]interface{}
	for educationRows.Next() {
		var education struct {
			InstitutionName string `json:"institution_name"`
			Degree          string `json:"degree"`
			Field           string `json:"field_of_study"`
			StartDate       string `json:"start_date"`
			EndDate         string `json:"end_date"`
			Grade           string `json:"grade"`
			Description     string `json:"description"`
		}

		if err := educationRows.Scan(&education.InstitutionName, &education.Degree, &education.Field, &education.StartDate, &education.EndDate, &education.Grade, &education.Description); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		educationDetails = append(educationDetails, map[string]interface{}{
			"institution_name": education.InstitutionName,
			"degree":           education.Degree,
			"field_of_study":   education.Field,
			"start_date":       education.StartDate,
			"end_date":         education.EndDate,
			"grade":            education.Grade,
			"description":      education.Description,
		})
	}

	// Retrieve Skills Details
	skillsRows, err := config.Database.Query(
		"SELECT skill_id, skill_name, proficiency_level FROM skills WHERE user_id = ?",
		userID,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer skillsRows.Close()

	var skillsDetails []map[string]interface{}
	for skillsRows.Next() {
		var skill struct {
			SkillID          int    `json:"skill_id"`
			SkillName        string `json:"skill_name"`
			ProficiencyLevel string `json:"proficiency_level"`
		}

		if err := skillsRows.Scan(&skill.SkillID, &skill.SkillName, &skill.ProficiencyLevel); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		skillsDetails = append(skillsDetails, map[string]interface{}{
			"skill_id":          skill.SkillID,
			"skill_name":        skill.SkillName,
			"proficiency_level": skill.ProficiencyLevel,
		})
	}

	// Retrieve Work Experience Details
	workRows, err := config.Database.Query(
		"SELECT job_title, company_name, location, start_date, end_date, responsibilities FROM work_experience WHERE user_id = ?",
		userID,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer workRows.Close()

	var workDetails []map[string]interface{}
	for workRows.Next() {
		var work struct {
			JobTitle         string `json:"job_title"`
			CompanyName      string `json:"company_name"`
			Location         string `json:"location"`
			StartDate        string `json:"start_date"`
			EndDate          string `json:"end_date"`
			Responsibilities string `json:"responsibilities"`
		}

		if err := workRows.Scan(&work.JobTitle, &work.CompanyName, &work.Location, &work.StartDate, &work.EndDate, &work.Responsibilities); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		workDetails = append(workDetails, map[string]interface{}{
			"job_title":        work.JobTitle,
			"company_name":     work.CompanyName,
			"location":         work.Location,
			"start_date":       work.StartDate,
			"end_date":         work.EndDate,
			"responsibilities": work.Responsibilities,
		})
	}

	response := map[string]interface{}{
		"user":      userDetails,
		"education": educationDetails,
		"skills":    skillsDetails,
		"work":      workDetails,
	}

	functions.Response(w, response)
}
