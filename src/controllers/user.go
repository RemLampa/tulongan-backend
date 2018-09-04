package controllers

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"text/template"
	"tulongan-backend/src/models"
)

var user *models.User = &models.User{}

func init() {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal("Could not get current working directory", err)
	}

	file, err := ioutil.ReadFile(filepath.Join(cwd, "../mock-db/db.json"))
	if err != nil {
		log.Fatal("Could not read file.", err)
	}

	if err := json.Unmarshal(file, user); err != nil {
		log.Fatal("Error in converting file to JSON.", err)
	}
}

// UserController handles user related operations
type UserController struct {
	User *models.User
}

// NewUserController creates a new user controller
func NewUserController() *UserController {
	userController := UserController{
		User: user,
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

// AddUserRepo adds a valid repository into current
// user's list
func (u *UserController) AddUserRepo(newRepo models.Repository) error {
	urlTemplate := "https://api.github.com/repos/{{ .Owner }}/{{ .Name }}/commits?author=RemLampa"

	t := template.Must(template.New("url").Parse(urlTemplate))

	urlBuff := new(bytes.Buffer)

	err := t.Execute(urlBuff, newRepo)
	if err != nil {
		return err
	}

	githubURL := urlBuff.String()

	r, err := http.Get(githubURL)
	if err != nil {
		return err
	}

	if r.StatusCode != 200 {
		err := errors.New("invalid repository provided")

		return err
	}

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		return err
	}

	if err := r.Body.Close(); err != nil {
		return err
	}

	var bodyJSON []interface{}

	if err := json.Unmarshal(body, &bodyJSON); err != nil { // unmarshall body contents as a type query
		return err
	}

	if len(bodyJSON) < 1 {
		err := errors.New("user is not a contributor in the given repo")

		return err
	}

	u.User.Repositories = append(u.User.Repositories, newRepo)

	return nil
}
