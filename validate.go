package neo4jclient

import "strings"

// Validate checks that the Cypher query looks valid
func (payload *Payload) Validate() error {

	for _, statement := range payload.Statements {
		err := validate(statement.Cypher)
		if err != nil {
			return err
		}
	}

	return nil

}

// validate a cypher query
func validate(cypher string) error {

	// Every delete query must contain a return
	if strings.Contains(cypher, "DELETE") {
		if !strings.Contains(cypher, "RETURN") {
			return errDeleteButNoReturn(cypher)
		}
	}

	return nil
}
