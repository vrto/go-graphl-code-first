package main

import "log"

func GetOneBeer(id int) *Beer {
	rows := RunQuery("SELECT * FROM beer WHERE id = $1", id)
	defer rows.Close()

	if !rows.Next() {
		return nil
	}

	var beer Beer
	if err := rows.Scan(&beer.Id, &beer.Name, &beer.Maker, &beer.Type); err != nil {
		log.Fatal(err)
	}
	return &beer
}
