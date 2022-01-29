package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/graphql-go/graphql"
)

func main() {
	schema := CreateGraphQlSchema()
	runQuery(schema, "Get Pub 1 (Full)", "queries/get-pub-full-1.graphql")
	runQuery(schema, "Get Pub 1 (Simple)", "queries/get-pub-simple-1.graphql")
	runQuery(schema, "Get Pub 2 (Full)", "queries/get-pub-full-2.graphql")
}

func runQuery(schema graphql.Schema, header, fileName string) {
	content, _ := ioutil.ReadFile(fileName)
	query := string(content)

	params := graphql.Params{Schema: schema, RequestString: query}
	response := graphql.Do(params)
	if len(response.Errors) > 0 {
		log.Fatalf("failed to execute graphql op, errors: %+v", response.Errors)
	}

	respJson, _ := json.MarshalIndent(response, "", "  ")
	fmt.Printf("\n===> %s\n%s \n", header, respJson)
}
