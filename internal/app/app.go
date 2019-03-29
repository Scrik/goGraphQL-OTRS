package app

import (
	"io/ioutil"
	"log"
	"net/http"

	graphiql "github.com/alexsuslov/go-graphiql"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"

	"github.com/goGraphQL-OTRS/internal/query"
)

// App App
type App struct {
	Config *TypeConfig
}

// New New
func New(data []byte) (*App, error) {

	Config, err := NewConfig(data)
	if err != nil {
		panic(err)
	}
	schema, err := GetSchema(Config.Schema)
	if err != nil {
		panic(err)
	}

	A := App{Config: Config}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Proxy GraphQL-OTRS API http://github.com/goGraphQL-OTRS"))
	})
	log.Println("Client URL", Config.ClientURL)
	http.HandleFunc(Config.ClientURL, graphiql.ServeGraphiQL)

	Schema := graphql.MustParseSchema(schema, &query.Query{})

	log.Println("GraphQL URL", Config.GraphiQLURL)
	http.Handle(Config.GraphiQLURL, &relay.Handler{Schema: Schema})

	log.Println("try Listen", Config.GetURL())
	err = http.ListenAndServe(Config.GetURL(), nil)
	return &A, err
}

// GetSchema GetSchema
func GetSchema(filename string) (string, error) {
	log.Println("try get GraphQL schema from file", filename)
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
