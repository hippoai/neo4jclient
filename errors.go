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

func errNeo4JRequest(err error) error {
	return goerr.New(ERR_NEO4J_REQUEST, map[string]interface{}{
		"error": err,
	})
}

func ErrNNotAvailable() error {
	return goerr.NewS(ERR_N_NOT_AVAILABLE)
}

func ErrCypherQuery(errs Errors) error {
	return goerr.New(ERR_CYPHER_QUERY, map[string]interface{}{
		"errors": errs,
	})
}
