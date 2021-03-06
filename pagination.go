package neo4jclient

import (
	"fmt"
	"strings"
)

// AddAtTheEnd adds a string at the end of the Cypher query
func (s *Statement) AddAtTheEnd(addMe string) *Statement {
	copiedS := s.Copy()
	copiedS.clean()
	copiedS.Cypher = fmt.Sprintf("%s\n%s", copiedS.Cypher, addMe)
	return copiedS
}

// AddOrderBy adds ordering to the query
func (s *Statement) AddOrderBy(ascending bool, orderMe ...string) *Statement {
	if (len(orderMe) == 0) || orderMe[0] == "" {
		return s
	}

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

func (s *Statement) Copy() *Statement {
	copiedResultDataContents := []string{}
	for _, elt := range s.ResultDataContents {
		copiedResultDataContents = append(copiedResultDataContents, elt)
	}

	copiedProps := map[string]interface{}{}
	for k, v := range s.Parameters.Props {
		copiedProps[k] = v
	}

	return &Statement{
		Cypher:             s.Cypher,
		Parameters:         Parameters{Props: copiedProps},
		ResultDataContents: copiedResultDataContents,
		Description:        s.Description,
		IsJustACount:       s.IsJustACount,
	}

}

// OnlyReturnACount breaks down the query row by row
// removes the previous RETURN if there was one on the last line
// and adds a count for the given "countMe" variable
func (s *Statement) OnlyReturnACount(countMe string) *Statement {
	// COUNT(*) if nothing was specified for the countMe parameter
	if countMe == "" {
		countMe = "*"
	}

	copiedS := s.Copy()

	copiedS.clean()
	cypherRows := strings.Split(copiedS.Cypher, "\n")

	cypherCount := fmt.Sprintf("RETURN COUNT(%s) AS _n", countMe)

	// Replace the last row if it does not have a RETURN
	newCypherRows := []string{}
	for _, cypherRow := range cypherRows {
		if !strings.Contains(cypherRow, "RETURN") {
			newCypherRows = append(newCypherRows, cypherRow)
		}
	}

	newCypherRows = append(newCypherRows, cypherCount)

	copiedS.Cypher = strings.Join(newCypherRows, "\n")
	copiedS.ResultDataContents = []string{"row"}
	copiedS.IsJustACount = true

	return copiedS
}
