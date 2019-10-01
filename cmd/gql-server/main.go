package main

import (
	log "github.com/AndreiD/go-gql-server/internal/logger"
	"github.com/AndreiD/go-gql-server/internal/orm"
	"github.com/AndreiD/go-gql-server/pkg/server"
)

func main() {

	// Create a new ORM instance to send it to our
	ormx, err := orm.Factory()
	if err != nil {
		log.Panic(err)
	}

	server.Run(ormx)
}
