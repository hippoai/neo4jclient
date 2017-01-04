package neo4jclient

import (
	"encoding/base64"
	"fmt"
	"log"

	"github.com/hippoai/env"
)

// Neo database structure
type Neo struct {
	username   string
	password   string
	endpoint   string
	authHeader string
}

// New tests a connection
func NewConnection(username, password, endpoint string) (*Neo, error) {

	data := []byte(fmt.Sprintf("%s:%s", username, password))
	token := base64.StdEncoding.EncodeToString(data)
	authHeader := fmt.Sprintf("Basic %s", token)

	neo := &Neo{
		username:   username,
		password:   password,
		authHeader: authHeader,
		endpoint:   endpoint,
	}

	return neo, nil

}

// NewConnectionFromEnv instanciates from environment variables
func NewConnectionFromEnv() (*Neo, error) {

	parsed, err := env.Parse(ENV_DB_USERNAME, ENV_DB_PASSWORD, ENV_DB_ENDPOINT)
	if err != nil {
		log.Fatalf(err.Error())
	}

	return NewConnection(parsed[ENV_DB_USERNAME], parsed[ENV_DB_PASSWORD], parsed[ENV_DB_ENDPOINT])

}
