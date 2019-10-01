package main

import (
	"strings"

	"github.com/AndreiD/go-gql-server/internal/logger"

	"github.com/AndreiD/go-gql-server/internal/orm"
	"github.com/AndreiD/go-gql-server/pkg/server"
)

// handle configs here!

func main() {
	var serverconf = &utils.ServerConfig{
		Host:          utils.MustGet("SERVER_HOST"),
		Port:          utils.MustGet("SERVER_PORT"),
		URISchema:     utils.MustGet("SERVER_URI_SCHEMA"),
		Version:       utils.MustGet("SERVER_PATH_VERSION"),
		SessionSecret: utils.MustGet("SESSION_SECRET"),
		JWT: utils.JWTConfig{
			Secret:    utils.MustGet("JWT_SECRET"),
			Algorithm: utils.MustGet("JWT_SIGNING_ALGORITHM"),
		},
		GraphQL: utils.GQLConfig{
			Path:                utils.MustGet("GQL_SERVER_GRAPHQL_PATH"),
			PlaygroundPath:      utils.MustGet("GQL_SERVER_GRAPHQL_PLAYGROUND_PATH"),
			IsPlaygroundEnabled: utils.MustGetBool("GQL_SERVER_GRAPHQL_PLAYGROUND_ENABLED"),
		},
		Database: utils.DBConfig{
			Dialect:     utils.MustGet("GORM_DIALECT"),
			DSN:         utils.MustGet("GORM_CONNECTION_DSN"),
			SeedDB:      utils.MustGetBool("GORM_SEED_DB"),
			LogMode:     utils.MustGetBool("GORM_LOGMODE"),
			AutoMigrate: utils.MustGetBool("GORM_AUTOMIGRATE"),
		},
		AuthProviders: []utils.AuthProvider{
			utils.AuthProvider{
				Provider:  "google",
				ClientKey: utils.MustGet("GOOGLE_KEY"),
				Secret:    utils.MustGet("GOOGLE_SECRET"),
			},
			utils.AuthProvider{
				Provider:  "auth0",
				ClientKey: utils.MustGet("AUTH0_KEY"),
				Secret:    utils.MustGet("AUTH0_SECRET"),
				Domain:    utils.MustGet("AUTH0_DOMAIN"),
				Scopes:    strings.Split(utils.MustGet("AUTH0_SCOPES"), ","),
			},
		},
	}
	orm, err := orm.Factory(serverconf)
	defer orm.DB.Close()
	if err != nil {
		logger.Panic(err)
	}
	server.Run(serverconf, orm)
}
