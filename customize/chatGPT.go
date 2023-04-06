package customize

import (
	"custom-go/pkg/plugins"
	"github.com/graphql-go/graphql"
)

type Person struct {
	Id        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

var (
	personType = graphql.NewObject(graphql.ObjectConfig{
		Name:        "Person",
		Description: "A person in the system",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"firstName": &graphql.Field{
				Type: graphql.String,
			},
			"lastName": &graphql.Field{
				Type: graphql.String,
			},
		},
	})

	fields = graphql.Fields{
		"person": &graphql.Field{
			Type:        personType,
			Description: "Get Person By ID",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				_ = plugins.GetGraphqlContext(params)
				id, ok := params.Args["id"].(int)
				if ok {
					testPeopleData := []Person{
						{Id: 1, FirstName: "John", LastName: "Doe"},
						{Id: 2, FirstName: "Jane", LastName: "Doe"},
					}
					for _, p := range testPeopleData {
						if p.Id == id {
							return p, nil
						}
					}
				}
				return nil, nil
			},
		},
	}

	rootQuery = graphql.ObjectConfig{Name: "RootQuery", Fields: fields}

	ChatGPT_schema, _ = graphql.NewSchema(graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)})
)
