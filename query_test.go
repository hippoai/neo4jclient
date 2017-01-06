package neo4jclient

import (
	"testing"

	"github.com/hippoai/goutil"
)

func TestQuery(t *testing.T) {

	username := "neo4j"
	password := "password"
	endpoint := "http://0.0.0.0:7474/db/data/transaction/commit"

	n, err := NewConnection(username, password, endpoint)
	if err != nil {
		t.Errorf(err.Error())
	}

	var out *Output

	out = insert(n, t)
	// Nodes - Must have two labels: Person and City, two nodes: Patrick and Denver
	// Edges - Two labels + One relationship between Patrick and Denver
	if (len(out.Merge.Nodes) != 4) || (len(out.Merge.Edges) != 3) {
		t.Fatal("Wrong number nodes and edges")
	}

	out = getPersons(n, t)
	// Nodes - One label and one person
	// Edges - One label
	if (len(out.Merge.Nodes) != 2) || (len(out.Merge.Edges) != 1) {
		t.Fatal("Wrong number nodes and edges for the persons")
	}

	out = delete(n, t)
	// Must be empty graph
	if (len(out.Merge.Nodes) > 0) || (len(out.Merge.Edges) > 0) {
		t.Fatal("Not empty graph after deletion")
	}
	// And the size of the deleted nodes and edges must be exact
	if (len(out.Delete.Nodes) != 2) || (len(out.Delete.Edges) != 1) {
		t.Fatal("Wrong deleted nodes and edges size")
	}

	out = getPersons(n, t)
	// Nodes - One label and one person
	// Edges - One label
	if (len(out.Merge.Nodes) != 0) || (len(out.Merge.Edges) != 0) {
		t.Fatal("Should be empty now")
	}

	t.Logf(goutil.Pretty(out))

}

func delete(n *Neo, t *testing.T) *Output {
	statement := `
    MATCH (person {key: {props}.personKey})
    MATCH (city {key: {props}.cityKey})
    MATCH (person)-[r {key: {props}.relKey}]-(city)
    DELETE city, r, person
    RETURN city, r, person
  `
	props := map[string]interface{}{
		"personKey": "person.patrick",
		"cityKey":   "city.denver",
		"relKey":    "person.patrick.LIVES_IN.city.denver",
	}
	payload := NewSinglePayload(statement, props)

	response, err := n.Request(payload)
	if err != nil {
		t.Errorf(err.Error())
	}

	out, err := Convert(response)
	if err != nil {
		t.Errorf(err.Error())
	}

	t.Logf(goutil.Pretty(out))

	return out

}

func insert(n *Neo, t *testing.T) *Output {
	statement := `
    MERGE (x:Person {name: {props}.personName, key: {props}.personKey})-[r:LIVES_IN {key: {props}.relKey}]->(s:City {name: {props}.city, key: {props}.cityKey})
    RETURN x, r, s
  `
	props := map[string]interface{}{
		"personKey":  "person.patrick",
		"personName": "patrick",
		"city":       "denver",
		"cityKey":    "city.denver",
		"relKey":     "person.patrick.LIVES_IN.city.denver",
	}
	payload := NewSinglePayload(statement, props)

	response, err := n.Request(payload)
	if err != nil {
		t.Errorf(err.Error())
	}

	out, err := Convert(response)
	if err != nil {
		t.Errorf(err.Error())
	}

	t.Logf(goutil.Pretty(out))

	return out

}

func getPersons(n *Neo, t *testing.T) *Output {
	statement := `
    MATCH (x:Person)
    RETURN x
  `
	props := map[string]interface{}{}
	payload := NewSinglePayload(statement, props)

	gr, err := n.RequestAndConvert(payload)
	if err != nil {
		t.Errorf(err.Error())
	}

	t.Logf(goutil.Pretty(gr))

	return gr

}
