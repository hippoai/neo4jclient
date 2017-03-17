package neo4jclient

import (
	"fmt"
	"strings"
)

// AddAtTheEnd adds a string at the end of the Cypher query
func (s *Statement) AddAtTheEnd(addMe string) *Statement {
	s.clean()
	s.Cypher = fmt.Sprintf("%s\n%s", s.Cypher, addMe)
	return s
}

// AddOrderBy adds ordering to the query
func (s *Statement) AddOrderBy(ascending bool, orderMe ...string) *Statement {
	orderStr := "DESC"
	if ascending {
		orderStr = "ASC"
	}

	orderThem := strings.Join(orderMe, ", ")

	orderBy := fmt.Sprintf("ORDER BY %s %s", orderThem, orderStr)

	return s.AddAtTheEnd(orderBy)
}

// AddSkipAndLimit adds skip and limit to the end of the cypher query
func (s *Statement) AddSkipAndLimit(skip, limit int) *Statement {
	return s.
		AddAtTheEnd(fmt.Sprintf("SKIP %d", skip)).
		AddAtTheEnd(fmt.Sprintf("LIMIT %d", limit))
}

func (s *Statement) clean() {
	cypherRows := strings.Split(s.Cypher, "\n")
	newRows := []string{}
	for _, cypherRow := range cypherRows {
		if strings.Replace(cypherRow, " ", "", -1) == "" {
			continue
		}
		newRows = append(newRows, cypherRow)
	}
	s.Cypher = strings.Join(newRows, "\n")
}

// OnlyReturnACount breaks down the query row by row
// removes the previous RETURN if there was one on the last line
// and adds a count for the given "countMe" variable
func (s *Statement) OnlyReturnACount(countMe string) *Statement {
	s.clean()
	cypherRows := strings.Split(s.Cypher, "\n")

	cypherCount := fmt.Sprintf("RETURN COUNT(%s) AS _n", countMe)

	// Replace the last row if it does not have a RETURN
	if strings.Contains(cypherRows[len(cypherRows)-1], "RETURN") {
		cypherRows[len(cypherRows)-1] = cypherCount
	} else {
		cypherRows = append(cypherRows, cypherCount)
	}

	s.Cypher = strings.Join(cypherRows, "\n")

	return s
}
