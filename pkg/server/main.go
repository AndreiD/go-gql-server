package server

import (
	"github.com/AndreiD/go-gql-server/internal/handlers"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

var host, port string

func init() {
	log.SetReportCaller(true)
	host = "localhost"
	port = "5555"
}

// Run .
func Run() {
	pathGQL := "/graphql"
	r := gin.Default()
	// Setup a route
	r.GET("/ping", handlers.Ping())

	log.Println("Running @ http://" + host + ":" + port + pathGQL)

	log.Fatalln(r.Run(host + ":" + port))
}
