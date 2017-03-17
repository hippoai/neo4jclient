package neo4jclient

import (
	"strings"
	"testing"
)

func removeWS(s string) string {
	return strings.Replace(s, " ", "", -1)
}

func TestAddAtTheEnd(t *testing.T) {

	cypher := `
    MATCH (c:Hello:World)
    MATCH (d:Hello:World)<-[r:HOLA]-(e:Playa)
    RETURN c
  `

	statement := NewStatementNoProps(cypher, "fake query")
	statement.AddAtTheEnd("I AM HERE")

	newCypher := `
    MATCH (c:Hello:World)
    MATCH (d:Hello:World)<-[r:HOLA]-(e:Playa)
    RETURN c
    I AM HERE
  `

	s2 := NewStatementNoProps(newCypher, "ok")
	s2.clean()

	if removeWS(statement.Cypher) != removeWS(s2.Cypher) {
		t.Fatalf("Expected cypher %s but got %s",
			removeWS(statement.Cypher),
			removeWS(s2.Cypher),
		)
	}

}

func TestAddOrderBy(t *testing.T) {

	cypher := `
    MATCH (c:Hello:World)
    MATCH (d:Hello:World)<-[r:HOLA]-(e:Playa)
    RETURN c
  `

	statement := NewStatementNoProps(cypher, "fake query")
	statement.AddOrderBy(false, "d.name", "e.id")

	newCypher := `
    MATCH (c:Hello:World)
    MATCH (d:Hello:World)<-[r:HOLA]-(e:Playa)
    RETURN c
    ORDER BY d.name, e.id DESC
  `

	s2 := NewStatementNoProps(newCypher, "ok")
	s2.clean()

	if removeWS(statement.Cypher) != removeWS(s2.Cypher) {
		t.Fatalf("Expected cypher %s but got %s",
			removeWS(statement.Cypher),
			removeWS(s2.Cypher),
		)
	}

}

func TestAddSkipAndLimit(t *testing.T) {
	cypher := `
    MATCH (c:Hello:World)
    MATCH (d:Hello:World)<-[r:HOLA]-(e:Playa)
    RETURN c
  `

	statement := NewStatementNoProps(cypher, "fake query")
	statement.
		AddOrderBy(false, "d.name", "e.id").
		AddSkipAndLimit(100, 10)

	newCypher := `
    MATCH (c:Hello:World)
    MATCH (d:Hello:World)<-[r:HOLA]-(e:Playa)
    RETURN c
    ORDER BY d.name, e.id DESC
    SKIP 100
    LIMIT 10
  `

	s2 := NewStatementNoProps(newCypher, "ok")
	s2.clean()

	if removeWS(statement.Cypher) != removeWS(s2.Cypher) {
		t.Fatalf("Expected cypher %s but got %s",
			removeWS(statement.Cypher),
			removeWS(s2.Cypher),
		)
	}

}

func TestOnlyReturnCount(t *testing.T) {
	cypher := `
    MATCH (c:Hello:World)
    MATCH (d:Hello:World)<-[r:HOLA]-(e:Playa)
    RETURN c
  `

	statement := NewStatementNoProps(cypher, "fake query")
	newStatement := statement.
		OnlyReturnACount("d.hello")

	newCypher := `
    MATCH (c:Hello:World)
    MATCH (d:Hello:World)<-[r:HOLA]-(e:Playa)
    RETURN COUNT(d.hello) AS _n
  `

	s2 := NewStatementNoProps(newCypher, "ok")
	s2.clean()

	if removeWS(newStatement.Cypher) != removeWS(s2.Cypher) {
		t.Fatalf("Expected cypher %s but got %s",
			removeWS(newStatement.Cypher),
			removeWS(s2.Cypher),
		)
	}

}
