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
	s := NewStatement(statement, map[string]interface{}{})
	return NewPayload(s)
}

// The following types help unmarshal a response

type ResultNode struct {
	Key    string                 `json:"id"`
	Labels []string               `json:"labels"`
	Props  map[string]interface{} `json:"properties"`
}
type ResultEdge struct {
	Key   string                 `json:"id"`
	Label string                 `json:"type"`
	Start string                 `json:"startNode"`
	End   string                 `json:"endNode"`
	Props map[string]interface{} `json:"properties"`
}
type ResultGraph struct {
	Nodes []ResultNode `json:"nodes"`
	Edges []ResultEdge `json:"relationships"`
}
type ResultData struct {
	Graph ResultGraph `json:"graph"`
}
type Result struct {
	Columns []string     `json:"columns"`
	Data    []ResultData `json:"data"`
}
type Results []Result
type Error interface{}
type Errors []Error
type Response struct {
	Results Results `json:"results"`
	Errors  Errors  `json:"errors"`
}
