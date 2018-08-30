package schema

import (
	"github.com/graphql-go/graphql"
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

	TulonganSchema, _ = graphql.NewSchema(graphql.SchemaConfig{
		Query: queryType,
	})
}
