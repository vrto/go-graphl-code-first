package main

import (
	"github.com/graphql-go/graphql"
	"github.com/pkg/errors"
)

func CreateBeerResolver(params graphql.ResolveParams) (interface{}, error) {
	name := params.Args["name"].(string)
	maker := params.Args["maker"].(string)
	var bType string
	if params.Args["type"] != nil {
		bType = params.Args["type"].(string)
	} else {
		bType = "Unknown"
	}

	if !Insert("INSERT INTO beer(name, maker, type) VALUES ($1, $2, $3)", name, maker, bType) {
		return false, errors.New("failed to save data, see logs for details")
	}

	return true, nil
}
