package main

import (
	"github.com/graphql-go/graphql"
	"log"
)

func CreateGraphQlSchema() graphql.Schema {
	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: createFields()}
	schemaConfig := graphql.SchemaConfig{
		Query: graphql.NewObject(rootQuery)}

	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("failed to create new schema, error :%v", err)
	}
	return schema
}

func createFields() graphql.Fields {
	return graphql.Fields{
		"pub": &graphql.Field{
			Type:        pubType,
			Description: "Get Pub by ID",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},
			Resolve: PubResolver,
		},
	}
}

var pubType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Pub",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
		},
	})
