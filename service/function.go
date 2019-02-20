package service

import (
	"github.com/andy-zhangtao/GoTrigger/model"
	"github.com/graphql-go/graphql"
	"strings"
)

var QueryTriggerJnl = &graphql.Field{
	Type:        graphql.NewList(TriggerJNL),
	Description: "query all trigger jnl",
	Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {

		return QueryAllTriggerJnl()
	},
}

var DeleteTriggerJnl = &graphql.Field{
	Type:        graphql.String,
	Description: "delete all trigger jnl",
	Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
		if err := DeleteALlTriggerJnl(); err != nil {
			return err.Error(), err
		}

		return "ok", nil
	},
}

var QueryPluginKind = &graphql.Field{
	Type:        graphql.NewList(PluginKind),
	Description: "query plugin info",
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {

		ps, err := FindAllPluginKind()
		if err != nil {
			if strings.Contains(err.Error(), "not found") {
				return nil, nil
			}

			return nil, err
		}

		return ps, nil
	},
}

var AddPluginKind = &graphql.Field{
	Type:        PluginKind,
	Description: "add plugin info",
	Args: graphql.FieldConfigArgument{
		"name": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"desc": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {

		name, _ := p.Args["name"].(string)
		desc, _ := p.Args["desc"].(string)

		return AddNewPluginKind(model.PluginType{
			Name: name,
			Desc: desc,
		})
	},
}
