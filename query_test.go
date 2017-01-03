package neo4jclient

import (
	"encoding/json"
	"testing"

	"github.com/hippoai/graphgo"
)

func TestQuery(t *testing.T) {

	username := "neo4j"
	password := "password"
	endpoint := "http://0.0.0.0:7474/db/data/transaction/commit"

	n, err := NewConnection(username, password, endpoint)
	if err != nil {
		t.Errorf(err.Error())
	}

	var g *graphgo.Graph

	g = insert(n, t)
	// Nodes - Must have two labels: Person and City, two nodes: Patrick and Denver
	// Edges - Two labels + One relationship between Patrick and Denver
	if (len(g.Nodes) != 4) || (len(g.Edges) != 3) {
		t.Fatal("Wrong number nodes and edges")
	}

	g = getPersons(n, t)
	// Nodes - One label and one person
	// Edges - One label
	if (len(g.Nodes) != 2) || (len(g.Edges) != 1) {
		t.Fatal("Wrong number nodes and edges for the persons")
	}

	g = delete(n, t)
	// Must be empty graph
	if (len(g.Nodes) > 0) || (len(g.Edges) > 0) {
		t.Fatal("Not empty graph after deletion")
	}

	g = getPersons(n, t)
	// Nodes - One label and one person
	// Edges - One label
	if (len(g.Nodes) != 0) || (len(g.Edges) != 0) {
		t.Fatal("Should be empty now")
	}

	t.Logf(prettyPrint(g))

}

func delete(n *Neo, t *testing.T) *graphgo.Graph {
	statement := "MATCH (x {key: \"person.patrick\"}) MATCH (y {key: \"city.denver\"}) MATCH (x)-[r {key: \"person.patrick.LIVES_IN.city.denver\"}]-(y) DELETE x, r, y"
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

	t.Logf(prettyPrint(g))

	return g

}

func insert(n *Neo, t *testing.T) *graphgo.Graph {
	statement := "MERGE (x:Person {name: \"patrick\", key: \"person.patrick\"})-[r:LIVES_IN {key: \"person.patrick.LIVES_IN.city.denver\"}]->(s:City {name: \"Denver\", key: \"city.denver\"}) RETURN x, r, s"
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

	t.Logf(prettyPrint(g))

	return g

}

func getPersons(n *Neo, t *testing.T) *graphgo.Graph {
	statement := "MATCH (x:Person) RETURN x"
	props := map[string]interface{}{}
	payload := NewSinglePayload(statement, props)

	g, err := n.RequestAndConvert(payload)
	if err != nil {
		t.Errorf(err.Error())
	}

	t.Logf(prettyPrint(g))

	return g

}

func prettyPrint(x interface{}) string {
	b, _ := json.MarshalIndent(x, "", "  ")
	return string(b)
}
