package schema

import (
	"github.com/graphql-go/graphql"
	"tulongan-backend/src/controllers"
	"tulongan-backend/src/models"
)

// TulonganSchema is our application schema
var TulonganSchema graphql.Schema

func init() {
	initUserSchema()

	queryType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"user": &graphql.Field{
				Type: userType,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return userType, nil
				},
			},
		},
	})

	mutationType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"createRepoMutation": &graphql.Field{
				Type: graphql.NewList(repoType),
				Args: graphql.FieldConfigArgument{
					"repo": &graphql.ArgumentConfig{
						Description: "The repository details",
						Type:        graphql.NewNonNull(createRepoType),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					repo := p.Args["repo"].(map[string]interface{})

					newRepo := models.Repository{
						Owner: repo["owner"].(string),
						Name:  repo["name"].(string),
					}

					u := controllers.NewUserController()

					err := u.AddUserRepo(newRepo)
					if err != nil {
						return nil, err
					}

					repos := u.GetUserRepos()

					return repos, nil
				},
			},
			"deleteRepoMutation": &graphql.Field{
				Type: graphql.NewList(repoType),
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Description: "The ID of the repo to be deleted",
						Type:        graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id := p.Args["id"].(int)

					u := controllers.NewUserController()

					err := u.DeleteUserRepo(id)
					if err != nil {
						return nil, err
					}

					repos := u.GetUserRepos()

					return repos, nil
				},
			},
		},
	})

	TulonganSchema, _ = graphql.NewSchema(graphql.SchemaConfig{
		Query:    queryType,
		Mutation: mutationType,
	})
}
