package neo4jclient

// The following types help marshal a Payload

type Parameters struct {
	Props map[string]interface{} `json:"props"`
}
type Statement struct {
	Cypher             string     `json:"statement"`
	Parameters         Parameters `json:"parameters"`
	ResultDataContents []string   `json:"resultDataContents"`
	Description        string     `json:"description"`
	IsJustACount       bool       `json:"isJustACount"`
}
type Statements []*Statement
type Payload struct {
	Statements Statements `json:"statements"`
}

// NewStatement formats a new statement for a payload
func NewStatement(cypher string, description string, props map[string]interface{}) *Statement {
	s := &Statement{
		Cypher:             cypher,
		ResultDataContents: []string{RESULT_DATA_CONTENTS},
		Description:        description,
		IsJustACount:       false,
	}
	if props != nil {
		s.Parameters = Parameters{Props: props}
	}

	return s
}

func NewStatementNoProps(cypher, description string) *Statement {
	return NewStatement(cypher, description, map[string]interface{}{})
}

// NewPayload instanciates a payload from a list of statement
func NewPayload(statements ...*Statement) *Payload {
	return &Payload{
		Statements: statements,
	}
}

// NewSinglePayload instanciates a payload from a single statement
func NewSinglePayload(statement, description string, props map[string]interface{}) *Payload {
	s := NewStatement(statement, description, props)
	return NewPayload(s)
}

// NewSinglePayload instanciates a payload from a single statement
func NewSinglePayloadNoProps(statement, description string) *Payload {
	return NewSinglePayload(statement, description, map[string]interface{}{})
}

func (p *Payload) SetDataContentsToRow() {
	r := []string{"row"}
	for _, statement := range p.Statements {
		statement.ResultDataContents = r
	}
}

// NewPaginatedPayload creates a payload
// from one statement, adding ordering, skip + limit to it
// and returning the total number of responses it ran alone (for pages number)
func NewPaginatedPayload(
	statement *Statement,
	ascending bool, orderMe string,
	skip, limit int,
) *Payload {

	return NewPayload(
		statement.
			AddOrderBy(ascending, orderMe).
			AddSkipAndLimit(skip, limit),
		statement.
			OnlyReturnACount(orderMe),
	)

}
