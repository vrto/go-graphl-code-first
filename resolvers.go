package main

import (
	"log"

	"github.com/graphql-go/graphql"
	_ "github.com/lib/pq"
)

func PubResolver(params graphql.ResolveParams) (interface{}, error) {
	id, ok := params.Args["id"].(int)
	if !ok {
		return nil, nil
	}

	db := OpenConnection()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM pub WHERE id = $1", id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	if !rows.Next() {
		return nil, nil
	}

	var pub Pub
	err = rows.Scan(&pub.Id, &pub.Name)
	if err != nil {
		log.Fatal(err)
	}
	return pub, nil
}
