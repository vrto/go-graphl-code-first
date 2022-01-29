package main

import (
	"fmt"
	"log"

	"github.com/graphql-go/graphql"
)

func CreateGraphQlSchema() graphql.Schema {
	schemaConfig := graphql.SchemaConfig{Query: queryType, Mutation: mutationType}

	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("failed to create new schema, error :%v", err)
	}
	return schema
}

type NotFoundError struct {
	Entity string
}

func NewNotFoundError(entity string) *NotFoundError {
	return &NotFoundError{Entity: entity}
}

func (e NotFoundError) Error() string {
	return fmt.Sprintf("%s not found!", e.Entity)
}
