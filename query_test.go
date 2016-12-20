package neo4jclient

import (
	"encoding/json"
	"testing"
)

func TestQuery(t *testing.T) {

	username := "neo4j"
	password := "password"
	endpoint := "http://0.0.0.0:7474/db/data/transaction/commit"

	n, err := NewConnection(username, password, endpoint)
	if err != nil {
		t.Errorf(err.Error())
	}

	statement := "MATCH (x)-[r]->(y) RETURN x, r, y"
	props := map[string]interface{}{}
	payload := NewSinglePayload(statement, props)

	response, err := n.Request(payload)
	if err != nil {
		t.Errorf(err.Error())
	}

	g, err := Convert(response)
	if err != nil {
		t.Errorf(err.Error())
	}

	serializedG, err := json.Marshal(g)
	if err != nil {
		t.Errorf(err.Error())
	}

	t.Logf(string(serializedG))

}
