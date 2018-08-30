package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"tulongan-backend/src/models"
)

// UserController handles user related operations
type UserController struct {
	User *models.User
}

// NewUserController creates a new user controller
func NewUserController() *UserController {
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

	userController := UserController{
		User: &user,
	}

	return &userController
}

// GetUserName returns the current user's username
func (u *UserController) GetUserName() string {
	return u.User.Username
}

// GetUserRepos returns the current user's repository list
func (u *UserController) GetUserRepos() []models.Repository {
	return u.User.Repositories
}
