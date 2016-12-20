package neo4jclient

import "errors"

func errorDBEnv() error {
	return errors.New(ERR_DB_CONNECT)
}

func errNeo() error {
	return errors.New(ERR_NEO)
}
