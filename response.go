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
	Row   []interface{} `json:"row"`
	Meta  []interface{} `json:"meta"`
	Graph ResultGraph   `json:"graph"`
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

// GetN returns the number of nodes, if available in one of the statements
func (response *Response) GetN() (int, error) {
	for _, result := range response.Results {
		if (len(result.Columns) != 1) || (result.Columns[0] != "_n") {
			continue
		}

		// At this point we now we have one column _n, we expect data to be of size 1 here
		if len(result.Data) != 1 {
			continue
		}

		nAsInt, ok := result.Data[0].Row[0].(int)
		if !ok {
			continue
		}
		return nAsInt, nil

	}

	return 0, ErrNNotAvailable()
}
