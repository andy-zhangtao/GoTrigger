package service

import (
	"context"
	"encoding/json"
	"github.com/andy-zhangtao/GoTrigger/model"
	"github.com/gorilla/mux"
	"github.com/graphql-go/graphql"
	"github.com/rs/cors"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"strings"
)

func GraphQL() {
	router := mux.NewRouter()
	router.Path("/api").HandlerFunc(handleDevExGraphQL)
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"HEAD", "GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
		Debug:            false,
		AllowOriginFunc: func(origin string) bool {
			return true
		},
	})

	handler := c.Handler(router)
	logrus.Fatal(http.ListenAndServe(":80", handler))
}

func handleDevExGraphQL(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	//fmt.Println(r.Header)
	ip := strings.Split(r.RemoteAddr, ":")[0]
	args := make(map[string]interface{})

	args["ip"] = ip

	g := make(map[string]interface{})

	if r.Method == http.MethodGet {
		g["query"] = r.URL.Query().Get("query")

		result := executeDevExQuery(g, schemaDevex, args)
		json.NewEncoder(w).Encode(result)
	}

	if r.Method == http.MethodPost {
		data, _ := ioutil.ReadAll(r.Body)
		err := json.Unmarshal(data, &g)
		if err != nil {
			json.NewEncoder(w).Encode(err.Error())
		}

		result := executeDevExQuery(g, schemaDevex, args)
		json.NewEncoder(w).Encode(result)
	}
}

var schemaDevex, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    rootDevexQuery,
	Mutation: rootMutation,
})

func executeDevExQuery(query map[string]interface{}, schema graphql.Schema, args map[string]interface{}) *graphql.Result {

	params := graphql.Params{
		Schema:        schema,
		RequestString: query["query"].(string),
	}

	if query["variables"] != nil {
		params.VariableValues = query["variables"].(map[string]interface{})
	}

	if len(args) > 0 {
		_context := context.Background()
		for key, value := range args {
			_context = context.WithValue(_context, key, value)
		}
		params.Context = _context
	}

	result := graphql.Do(params)

	if len(result.Errors) > 0 {
		logrus.WithFields(logrus.Fields{"wrong result, unexpected errors:": result.Errors}).Error(model.MODULENAME)
	}
	return result
}

var rootDevexQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		"queryTrigger":       QueryTrigger,
		"queryTriggerPlugin": QueryTriggerPlugin,
		"queryPluginKind":    QueryPluginKind,
		"queryTriggerJnl":    QueryTriggerJnl,
	},
})

var rootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootMutation",
	Fields: graphql.Fields{
		"enableTrigger":       UpdateTriggerEnable,
		"addTrigger":          AddTrigger,
		"delTrigger":          DelTrigger,
		"addTriggerPlugin":    AddTriggerPlugin,
		"addPluginKind":       AddPluginKind,
		"delTriggerPlugin":    DelTriggerPlugin,
		"updateTriggerPlugin": UpdateTriggerPlugin,
		"delTriggerJnl":       DeleteTriggerJnl,
		"fireTrigger":         TriggerFire,
		"updatePluginKind":    UpdatePluginKind,
	},
})
