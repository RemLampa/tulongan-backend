package models

// Repository is the repo object
type Repository struct {
	Owner string
	Name  string
}

// User is the user object
type User struct {
	Username     string
	Repositories []Repository
}
