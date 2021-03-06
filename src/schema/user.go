package schema

import (
	"github.com/graphql-go/graphql"
	"tulongan-backend/src/controllers"
	"tulongan-backend/src/models"
)

var userType *graphql.Object
var repoType *graphql.Object
var createRepoType *graphql.InputObject
var deleteRepoType *graphql.InputObject

func initUserSchema() {
	repoType = graphql.NewObject(graphql.ObjectConfig{
		Name:        "Repo",
		Description: "A GitHub repository wherein the user is a validated contributor",
		Fields: graphql.Fields{
			"owner": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "The owner of the repo",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if repo, ok := p.Source.(models.Repository); ok {
						return repo.Owner, nil
					}

					return nil, nil
				},
			},
			"name": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "The name of the repo",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if repo, ok := p.Source.(models.Repository); ok {
						return repo.Name, nil
					}

					return nil, nil
				},
			},
		},
	})

	userType = graphql.NewObject(graphql.ObjectConfig{
		Name:        "User",
		Description: "Retrieves data for the given user",
		Fields: graphql.Fields{
			"userName": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "Unique username",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					u := controllers.NewUserController()

					return u.GetUserName(), nil
				},
			},
			"repositories": &graphql.Field{
				Type:        graphql.NewList(repoType),
				Description: "The user's validated Github repos",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					u := controllers.NewUserController()

					return u.GetUserRepos(), nil
				},
			},
		},
	})

	createRepoType = graphql.NewInputObject(graphql.InputObjectConfig{
		Name: "CreateRepo",
		Fields: graphql.InputObjectConfigFieldMap{
			"owner": &graphql.InputObjectFieldConfig{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "The owner of the repository",
			},
			"name": &graphql.InputObjectFieldConfig{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "The name of the repository",
			},
		},
	})
}
