package neo4jclient

import "github.com/hippoai/graphgo"

type Delete struct {
	Nodes []string
	Edges []string
}

// Output defines the expected output from this client
// after a call to Neo4J
type Output struct {
	Merge  *graphgo.Graph `json:"$merge"`
	Delete *Delete        `json:"$delete"`
}

func NewOutput() *Output {
	return &Output{
		Merge: graphgo.NewEmptyGraph(),
		Delete: &Delete{
			Nodes: []string{},
			Edges: []string{},
		},
	}
}
