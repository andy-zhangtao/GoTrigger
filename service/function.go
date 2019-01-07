package service

import (
	"github.com/andy-zhangtao/GoTrigger/model"
	"github.com/graphql-go/graphql"
)

var QueryTrigger = &graphql.Field{
	Type:        graphql.String,
	Description: "query trigger info",
	Args: graphql.FieldConfigArgument{
		"name": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"enable": &graphql.ArgumentConfig{
			Type: graphql.Boolean,
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {

		name, _ := p.Args["name"].(string)
		//enable, _ := p.Args["enable"].(bool)

		t, err := FindSpecifyTrigger(name)
		if err != nil {
			return nil, err
		}

		return t, nil
	},
}

var AddTrigger = &graphql.Field{
	Type:        graphql.String,
	Description: "register new trigger",
	Args: graphql.FieldConfigArgument{
		"name": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"enable": &graphql.ArgumentConfig{
			Type: graphql.Boolean,
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {

		name, _ := p.Args["name"].(string)
		enable, _ := p.Args["enable"].(bool)

		t := model.Trigger{
			Name:   name,
			Enable: enable,
		}

		return t, AddNewTrigger(t)
	},
}
