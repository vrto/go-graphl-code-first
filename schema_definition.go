package main

import (
	"fmt"
	"log"

	"github.com/graphql-go/graphql"
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
		"beer": &graphql.Field{
			Type:        beerType,
			Description: "Get Beer by ID",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},
			Resolve: BeerResolver,
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
			"beers": &graphql.Field{
				Type:    graphql.NewList(beerType),
				Resolve: BeersForPubResolver,
			},
		},
	})

var beerType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Beer",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"maker": &graphql.Field{
				Type: graphql.String,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"type": &graphql.Field{
				Type: graphql.String,
			},
			"pubs": &graphql.Field{
				Type:    graphql.NewList(graphql.String),
				Resolve: PubsForBeerResolver,
			},
		},
	})

type NotFoundError struct {
	Entity string
}

func NewNotFoundError(entity string) *NotFoundError {
	return &NotFoundError{Entity: entity}
}

func (e NotFoundError) Error() string {
	return fmt.Sprintf("%s not found!", e.Entity)
}
