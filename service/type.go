package service

import (
	"github.com/andy-zhangtao/GoTrigger/model"
	"github.com/graphql-go/graphql"
	"strconv"
	"time"
)

var Trigger = graphql.NewObject(graphql.ObjectConfig{
	Name: "trigger",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if t, ok := p.Source.(model.Trigger); ok {
					return t.ID.Hex(), nil
				}

				return nil, nil
			},
		},
		"name": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if t, ok := p.Source.(model.Trigger); ok {
					return t.Name, nil
				}

				return nil, nil
			},
		},
		"enable": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if t, ok := p.Source.(model.Trigger); ok {
					return t.Enable, nil
				}

				return nil, nil
			},
		},
		"next_time": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if t, ok := p.Source.(model.Trigger); ok {
					return t.NextTime, nil
				}

				return nil, nil
			},
		},
		"str_next_time": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if t, ok := p.Source.(model.Trigger); ok {
					i, err := strconv.ParseInt(strconv.Itoa(int(t.NextTime)), 10, 64)
					if err != nil {
						return "", err
					}
					tm := time.Unix(i, 0)
					return tm, nil
				}

				return nil, nil
			},
		},
		"parallel": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if t, ok := p.Source.(model.Trigger); ok {
					return t.Parallel, nil
				}

				return nil, nil
			},
		},
		//"type": &graphql.Field{
		//	Type: graphql.Int,
		//	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		//		if t, ok := p.Source.(model.Trigger); ok {
		//			return t.TriggerType, nil
		//		}
		//
		//		return nil, nil
		//	},
		//},
		//"typeID": &graphql.Field{
		//	Type: graphql.String,
		//	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		//		if t, ok := p.Source.(model.Trigger); ok {
		//			return t.TypeID.Hex(), nil
		//		}
		//
		//		return nil, nil
		//	},
		//},
		"create_time": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if t, ok := p.Source.(model.Trigger); ok {
					return t.CreateTime, nil
				}

				return nil, nil
			},
		},
	},
})
