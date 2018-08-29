package tulonganutil

import (
	"github.com/graphql-go/graphql"
)

// TestData is our test data
var TestData string

var testDataType *graphql.Object

// TulonganSchema is our application schema
var TulonganSchema graphql.Schema

func init() {
	TestData = "foo bar"

	// testDataType = graphql.NewObject(graphql.ObjectConfig{
	// 	Name: "Test Data",
	// 	Description: "Just Some Test Deta",
	// 	Fields: graphql.Fields{
	// 		"id": &graphql.Field{
	// 			Type: graphql.NewNonNull(graphql.String),
	// 			Description: "Unique identifier of our test data",
	// 			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
	// 				return "some-id"
	// 			},
	// 		},
	// 		"value": &graphql.Field{
	// 			Types: grapqhl.String,
	// 			Description: "The value of our test data",
	// 			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
	// 				return TestData
	// 			}
	// 		},
	// 	},
	// })

	queryType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"testData": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return TestData, nil
				},
			},
		},
	})

	TulonganSchema, _ = graphql.NewSchema(graphql.SchemaConfig{
		Query: queryType,
	})
}
