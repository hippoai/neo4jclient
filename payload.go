package neo4jclient

// The following types help marshal a Payload

type Parameters struct {
	Props map[string]interface{} `json:"props"`
}
type Statement struct {
	Statement          string     `json:"statement"`
	Parameters         Parameters `json:"parameters"`
	ResultDataContents []string   `json:"resultDataContents"`
}
type Statements []*Statement
type Payload struct {
	Statements Statements `json:"statements"`
}

// NewStatement formats a new statement for a payload
func NewStatement(statement string, props map[string]interface{}) *Statement {
	s := &Statement{
		Statement:          statement,
		ResultDataContents: []string{RESULT_DATA_CONTENTS},
	}
	if props != nil {
		s.Parameters = Parameters{Props: props}
	}

	return s
}

func NewStatementNoProps(statement string) *Statement {
	return NewStatement(statement, map[string]interface{}{})
}

// NewPayload instanciates a payload from a list of statement
func NewPayload(statements ...*Statement) *Payload {
	return &Payload{
		Statements: statements,
	}
}

// NewSinglePayload instanciates a payload from a single statement
func NewSinglePayload(statement string, props map[string]interface{}) *Payload {
	s := NewStatement(statement, props)
	return NewPayload(s)
}

// NewSinglePayload instanciates a payload from a single statement
func NewSinglePayloadNoProps(statement string) *Payload {
	return NewSinglePayload(statement, map[string]interface{}{})
}
