package service

import (
	"github.com/andy-zhangtao/GoTrigger/model"
	"github.com/andy-zhangtao/GoTrigger/util"
	"github.com/graphql-go/graphql"
	"strings"
)

var UpdateTriggerEnable = &graphql.Field{
	Type:        graphql.String,
	Description: "Enable / Disable specify trigger",
	Args: graphql.FieldConfigArgument{
		"name": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"enable": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Boolean),
		},
	},
	Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
		name, _ := p.Args["name"].(string)
		enable, _ := p.Args["enable"].(bool)

		t, err := FindSpecifyTrigger(name)
		if err != nil {
			if strings.Contains(err.Error(), "not found") {
				return nil, err
			}

			return nil, err
		}

		t.Enable = enable
		if t.Enable {
			e = EnableTrigger(name)
		} else {
			e = DisableTrigger(name)
		}

		if e != nil {
			return e.Error(), e
		}

		util.GetTriggerChan() <- t.ID
		return "ok", nil
	},
}

var QueryTrigger = &graphql.Field{
	Type:        graphql.NewList(Trigger),
	Description: "query trigger info",
	Args: graphql.FieldConfigArgument{
		"name": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		//"enable": &graphql.ArgumentConfig{
		//	Type: graphql.Boolean,
		//},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {

		name, ok := p.Args["name"].(string)
		//enable, _ := p.Args["enable"].(bool)
		if ok {
			t, err := FindSpecifyTrigger(name)
			if err != nil {
				if strings.Contains(err.Error(), "not found") {
					return nil, nil
				}

				return nil, err
			}

			return t, nil
		}

		ts, err := FindSpecifyAllTrigger()
		if err != nil {
			if strings.Contains(err.Error(), "not found") {
				return nil, nil
			}

			return nil, err
		}

		return ts, nil
	},
}

var DelTrigger = &graphql.Field{
	Type:        graphql.String,
	Description: "delete specify trigger",
	Args: graphql.FieldConfigArgument{
		"name": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		name, _ := p.Args["name"].(string)

		return "ok", DeleTrigger(name)
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
		"desc": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"interval": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"parallel": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"kind": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"endpoint": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {

		name, _ := p.Args["name"].(string)
		desc, _ := p.Args["desc"].(string)
		enable, _ := p.Args["enable"].(bool)
		interval, _ := p.Args["interval"].(string)
		parallel, _ := p.Args["parallel"].(int)

		kind, _ := p.Args["kind"].(int)
		endpoint, _ := p.Args["endpoint"].(string)

		if parallel == 0 {
			parallel = 1
		}

		name = strings.TrimSpace(name)
		interval = strings.TrimSpace(interval)
		endpoint = strings.TrimSpace(endpoint)

		intervalTime, err := util.ParseInterval(interval)
		if err != nil {
			return nil, err
		}

		t := model.Trigger{
			Name:     name,
			Enable:   enable,
			Desc:     desc,
			Parallel: parallel,
			Interval: intervalTime,
			NextTime: int64(util.NextTime(intervalTime)),
			Type: model.TriggerType{
				Kind:     kind,
				Endpoint: endpoint,
			},
		}

		if t, err = AddNewTrigger(t); err != nil {
			return t, err
		}

		util.GetTriggerChan() <- t.ID
		return t, nil
	},
}