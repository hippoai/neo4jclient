package main

import (
	"log"

	"github.com/hippoai/env"
)

// Initializes the environment variables and Keen client
func init() {
	_, err := env.Parse(
		ENV_DB_USERNAME, ENV_DB_PASSWORD, ENV_DB_ENDPOINT,
	)
	if err != nil {
		log.Fatalf(err.Error())
	}

}
