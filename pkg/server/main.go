package server

import (
	"github.com/AndreiD/go-gql-server/internal/handlers"
	"github.com/AndreiD/go-gql-server/internal/orm"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

var host, port string
var isPgEnabled bool

const GQL_SERVER_GRAPHQL_PATH = "/graphql"
const GQL_SERVER_GRAPHQL_PLAYGROUND_PATH = "/"

func init() {
	log.SetReportCaller(true)
	host = "localhost"
	port = "5555"
	isPgEnabled = true
}

// Run .
func Run(orm *orm.ORM) {
	r := gin.Default()
	// Setup a route
	r.GET("/ping", handlers.Ping())

	// GraphQL handlers
	// Playground handler
	if isPgEnabled {
		r.GET(GQL_SERVER_GRAPHQL_PLAYGROUND_PATH, handlers.PlaygroundHandler(GQL_SERVER_GRAPHQL_PATH))
		log.Println("GraphQL Playground @ " + host + ":" + port + GQL_SERVER_GRAPHQL_PLAYGROUND_PATH)
	}
	r.POST(GQL_SERVER_GRAPHQL_PATH, handlers.GraphqlHandler(orm))
	log.Println("GraphQL @ " + host + ":" + port + GQL_SERVER_GRAPHQL_PATH)

	log.Println("Running @ http://" + host + ":" + port + GQL_SERVER_GRAPHQL_PATH)

	log.Fatalln(r.Run(host + ":" + port))
}
