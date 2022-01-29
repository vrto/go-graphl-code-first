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

//TODO return err instead of logging?
func BeersForPubResolver(params graphql.ResolveParams) (interface{}, error) {
	pubId := params.Source.(Pub).Id
	db := OpenConnection()
	defer db.Close()

	rows, err := db.Query("SELECT id, name, maker, type FROM beer JOIN pubs_beers ON id = beer_id WHERE pub_id = $1", pubId)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var beers []Beer
	for rows.Next() {
		var beer Beer
		err = rows.Scan(&beer.Id, &beer.Name, &beer.Maker, &beer.Type)
		if err != nil {
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

	db := OpenConnection()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM beer WHERE id = $1", id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	if !rows.Next() {
		return nil, nil
	}

	var beer Beer
	err = rows.Scan(&beer.Id, &beer.Name, &beer.Maker, &beer.Type)
	if err != nil {
		log.Fatal(err)
	}
	return beer, nil
}
