package neo4jclient

// The following types help unmarshal a response from Neo4J

type ResultNode struct {
	LegacyKey string                 `json:"id"`
	Labels    []string               `json:"labels"`
	Props     map[string]interface{} `json:"properties"`
	Deleted   bool                   `json:"deleted"`
}
type ResultEdge struct {
	LegacyKey string                 `json:"id"`
	Label     string                 `json:"type"`
	Start     string                 `json:"startNode"`
	End       string                 `json:"endNode"`
	Props     map[string]interface{} `json:"properties"`
	Deleted   bool                   `json:"deleted"`
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
