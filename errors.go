package neo4jclient

import "github.com/hippoai/goerr"

func errorDBEnv() error {
	return goerr.New(ERR_DB_CONNECT, map[string]interface{}{})
}

func errNeo(payload interface{}) error {
	return goerr.New(ERR_NEO, map[string]interface{}{
		"reason": payload,
	})
}

func errDeleteButNoReturn(cypher string) error {
	return goerr.New(ERR_DELETE_BUT_NO_RETURN, map[string]interface{}{
		"cypher": ERR_DELETE_BUT_NO_RETURN,
	})
}
