package neo4jclient

import (
	"encoding/base64"
	"fmt"
	"os"
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
	username := os.Getenv(ENV_DB_USERNAME)
	password := os.Getenv(ENV_DB_PASSWORD)
	endpoint := os.Getenv(ENV_DB_ENDPOINT)

	if username == "" || password == "" || endpoint == "" {
		return nil, errorDBEnv()
	}

	return NewConnection(username, password, endpoint)

}
