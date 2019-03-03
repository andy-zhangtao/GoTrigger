package service

import (
	"github.com/andy-zhangtao/GoTrigger/model"
	"github.com/graphql-go/graphql"
)

var QueryTriggerPlugin = &graphql.Field{
	Type:        graphql.NewList(TriggerPlugin),
	Description: "query trigger plugin info",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
		id, ok := p.Args["id"].(int)
		if !ok {
			return FindAllPlugin()
		}

		f, err := FindSpecifyPlugin(id)
		return []model.TriggerPlugin{f}, err
	},
}

var AddTriggerPlugin = &graphql.Field{
	Type:        TriggerPlugin,
	Description: "register new trigger plugin",
	Args: graphql.FieldConfigArgument{
		"name": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"endpoint": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"desc": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
		name, _ := p.Args["name"].(string)
		id, _ := p.Args["id"].(int)
		desc, _ := p.Args["desc"].(string)
		endpoint, _ := p.Args["endpoint"].(string)

		return AddNewPlugin(model.TriggerPlugin{
			Name:     name,
			PID:      id,
			Endpoint: endpoint,
			Desc:     desc,
		})
	},
}

var DelTriggerPlugin = &graphql.Field{
	Type:        graphql.String,
	Description: "delete trigger plugin",
	Args: graphql.FieldConfigArgument{
		"name": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
		name, _ := p.Args["name"].(string)

		return "ok", DeleteSpecifyPlugin(name)
	},
}

var UpdateTriggerPlugin = &graphql.Field{
	Type:        graphql.String,
	Description: "update trigger plugin",
	Args: graphql.FieldConfigArgument{
		"name": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"endpoint": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"desc": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
		name, _ := p.Args["name"].(string)
		desc, _ := p.Args["desc"].(string)
		endpoint, _ := p.Args["endpoint"].(string)

		return "ok", UpdateSpecifyPlugin(model.TriggerPlugin{
			Name:     name,
			Desc:     desc,
			Endpoint: endpoint,
		})
	},
}
