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

	rows := RunQuery("SELECT * FROM pub WHERE id = $1", id)
	defer rows.Close()

	if !rows.Next() {
		return nil, NewNotFoundError("pub")
	}

	var pub Pub
	if err := rows.Scan(&pub.Id, &pub.Name); err != nil {
		log.Fatal(err)
	}
	return pub, nil
}

func BeersForPubResolver(params graphql.ResolveParams) (interface{}, error) {
	pubId := params.Source.(Pub).Id

	rows := RunQuery("SELECT id, name, maker, type FROM beer JOIN pubs_beers ON id = beer_id WHERE pub_id = $1", pubId)
	defer rows.Close()

	var beers []Beer
	for rows.Next() {
		var beer Beer
		if err := rows.Scan(&beer.Id, &beer.Name, &beer.Maker, &beer.Type); err != nil {
			log.Fatal(err)
		}
		beers = append(beers, beer)
	}
	return beers, nil
}

func BeerResolver(params graphql.ResolveParams) (interface{}, error) {
	id, ok := params.Args["id"].(int)
	if !ok {
		return nil, nil
	}

	rows := RunQuery("SELECT * FROM beer WHERE id = $1", id)
	defer rows.Close()

	if !rows.Next() {
		return nil, NewNotFoundError("beer")
	}

	var beer Beer
	if err := rows.Scan(&beer.Id, &beer.Name, &beer.Maker, &beer.Type); err != nil {
		log.Fatal(err)
	}
	return beer, nil
}
