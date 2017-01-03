package neo4jclient

import "github.com/hippoai/goerr"

func errorDBEnv() error {
	return goerr.New(ERR_DB_CONNECT, map[string]interface{}{})
}

func errNeo() error {
	return goerr.New(ERR_NEO, map[string]interface{}{})
}
