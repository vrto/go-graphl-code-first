package main

import "github.com/graphql-go/graphql"

var queryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name:   "RootQuery",
		Fields: createQueryFields(),
	})

func createQueryFields() graphql.Fields {
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
		"beers": &graphql.Field{
			Type:        graphql.NewList(beerType),
			Description: "List All Beers",
			Resolve:     BeersResolver,
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
