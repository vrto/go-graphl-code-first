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

	if !Exec("INSERT INTO beer(name, maker, type) VALUES ($1, $2, $3)", name, maker, bType) {
		return false, errors.New("failed to save data, see logs for details")
	}

	return true, nil
}

func UpdateBeerTypeResolver(params graphql.ResolveParams) (interface{}, error) {
	id := params.Args["id"].(int)
	bType := params.Args["type"].(string)

	if !Exec("UPDATE beer SET type = $1 WHERE id = $2", bType, id) {
		return nil, errors.New("failed to save data, see logs for details")
	}

	return GetOneBeer(id), nil
}
