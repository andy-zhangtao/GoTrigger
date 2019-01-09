package service

import (
	"github.com/andy-zhangtao/GoTrigger/model"
	"github.com/andy-zhangtao/GoTrigger/util"
	"github.com/graphql-go/graphql"
	"strings"
)

var QueryTrigger = &graphql.Field{
	Type:        Trigger,
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
			if strings.Contains(err.Error(), "not found") {
				return nil, nil
			}

			return nil, err
		}

		return t, nil
	},
}

var AddTrigger = &graphql.Field{
	Type:        Trigger,
	Description: "register new trigger",
	Args: graphql.FieldConfigArgument{
		"name": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"enable": &graphql.ArgumentConfig{
			Type: graphql.Boolean,
		},
		"interval": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"parallel": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {

		name, _ := p.Args["name"].(string)
		enable, _ := p.Args["enable"].(bool)
		interval, _ := p.Args["interval"].(string)
		parallel, _ := p.Args["parallel"].(int)

		if parallel == 0 {
			parallel = 1
		}

		intervalTime, err := util.ParseInterval(interval)
		if err != nil {
			return nil, err
		}

		t := model.Trigger{
			Name:     name,
			Enable:   enable,
			Parallel: parallel,
			Interval: intervalTime,
			NextTime: uint64(util.NextTime(intervalTime)),
		}

		return t, AddNewTrigger(t)
	},
}
