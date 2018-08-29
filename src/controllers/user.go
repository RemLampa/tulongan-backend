package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"tulongan-backend/src/models"
)

// GetUser retrieves the user from database
func GetUser() models.User {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal("Could not get current working directory", err)
	}

	file, err := ioutil.ReadFile(filepath.Join(cwd, "../mock-db/db.json"))
	if err != nil {
		log.Fatal("Could not read file.", err)
	}

	var user models.User

	if err := json.Unmarshal(file, &user); err != nil {
		log.Fatal("Error in converting file to JSON.", err)
	}

	return user
}
