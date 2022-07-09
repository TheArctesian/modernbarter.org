package main

import (
    "encoding/json"
    "fmt"
    "log"
    "time"

    "github.com/graphql-go/graphql"
)

type Commodities struct {
    id int32
    Name string 
    Price float32
    dateUpdated  time.Time
}


type Relationship {
    id int32
    comparingCommodities []Commodities
    Relationship string // example 4-1
    perceivedRelationship []string
}

func populate() []Commodities {

}
func main() {
    // Schema
    fields := graphql.Fields{
        "hello": &graphql.Field{
            Type: graphql.String,
            Resolve: func(p graphql.ResolveParams) (interface{}, error) {
                return "world", nil
            },
        },
    }
    rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
    schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
    schema, err := graphql.NewSchema(schemaConfig)
    if err != nil {
        log.Fatalf("failed to create new schema, error: %v", err)
    }

    // Query
    query := `
        {
            hello
        }
    `
    params := graphql.Params{Schema: schema, RequestString: query}
    r := graphql.Do(params)
    if len(r.Errors) > 0 {
        log.Fatalf("failed to execute graphql operation, errors: %+v", r.Errors)
    }
    rJSON, _ := json.Marshal(r)
    fmt.Printf("%s \n", rJSON) // {“data”:{“hello”:”world”}}
}