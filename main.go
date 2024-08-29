package main

import (
	"fmt"
	"hr/config"
	"hr/routes/auth"
	"hr/routes/resume"
	"hr/routes/update"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	config.ConnectDB()

	// Set up the router
	router := mux.NewRouter()
	router.HandleFunc("/signup", auth.Signupuser).Methods("POST")
	router.HandleFunc("/login", auth.Loginuser).Methods("POST")
	router.HandleFunc("/update", update.UpdateUserDetails).Methods("POST")
	router.HandleFunc("/add-education", resume.Education).Methods("POST")
	router.HandleFunc("/add-experience", resume.WorkExperience).Methods("POST")
	router.HandleFunc("/add-skills", resume.Skills).Methods("POST")
	router.HandleFunc("/get-details", resume.GetUserDetails).Methods("POST")

	c := cors.AllowAll()

	fmt.Println("Running....")
	handler := c.Handler(router)
	loggedHandler := handlers.LoggingHandler(os.Stdout, handler)

	http.ListenAndServe(":8080", loggedHandler)
}
