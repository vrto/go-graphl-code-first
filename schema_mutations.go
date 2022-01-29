package main

import "github.com/graphql-go/graphql"

var mutationType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"createBeer": &graphql.Field{
				Type:        graphql.Boolean,
				Description: "create a new beer",
				Args: graphql.FieldConfigArgument{
					"maker": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"name": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"type": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: CreateBeerResolver,
			},
			"updateBeer": &graphql.Field{
				Type:        beerType,
				Description: "update an existing beer",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"type": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: UpdateBeerTypeResolver,
			},
		},
	})
