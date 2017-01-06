package neo4jclient

import (
	"fmt"
	"log"

	"github.com/hippoai/graphgo"
)

type Delete struct {
	Nodes []string
	Edges []string
}

type GraphResponse struct {
	Merge  *graphgo.Graph `json:"$merge"`
	Delete *Delete        `json:"$delete"`
}

func NewGraphResponse() *GraphResponse {
	return &GraphResponse{
		Merge: graphgo.NewEmptyGraph(),
		Delete: &Delete{
			Nodes: []string{},
			Edges: []string{},
		},
	}
}

// getLabelKey returns the key for a label node
func getLabelKey(label string) string {
	return fmt.Sprintf("label.%s", label)
}

// getLabelRelKey returns the key for the relationship between a node and its label
func getLabelRelKey(nodeKey, label string) string {
	return fmt.Sprintf("%s.%s.%s", nodeKey, graphgo.NODE_LABEL_EDGE_LABEL, label)
}

// Convert neo4J response to Graphgo's format for Merge, and a list of what to delete
func Convert(r *Response) (*GraphResponse, error) {

	if len(r.Errors) > 0 {
		log.Println("Response error", r.Errors)
		return nil, errNeo(r.Errors)
	}

	// Initialize an empty graph response
	gr := NewGraphResponse()

	nodesKeyMap := map[string]string{}

	// Loop over all the results, and flatten this into the graph
	for _, result := range r.Results {

		// For each result (== statement), flatten all the rows in the graph
		for _, resultData := range result.Data {

			// First add all the nodes
			for _, resultNode := range resultData.Graph.Nodes {
				customKey := getCustomKey(resultNode.LegacyKey, resultNode.Props)

				nodesKeyMap[resultNode.LegacyKey] = customKey

				// Add deleted nodes
				if resultNode.Deleted {
					gr.Delete.Nodes = append(gr.Delete.Nodes, resultNode.LegacyKey)
					continue
				}

				// Add to legacy index
				gr.Merge.AddNodeLegacyIndex(resultNode.LegacyKey, customKey)

				// Add the node and its properties
				node, _ := gr.Merge.MergeNode(customKey, resultNode.Props)

				// Add the labels
				for _, resultLabel := range resultNode.Labels {
					_labelKey := getLabelKey(resultLabel)
					gr.Merge.MergeNode(_labelKey, nil)
					gr.Merge.MergeEdge(
						getLabelRelKey(node.Key, resultLabel),
						graphgo.NODE_LABEL_EDGE_LABEL,
						customKey,
						_labelKey,
						nil,
					)
				}

			}

			// Now, add all the relationships
			for _, resultEdge := range resultData.Graph.Edges {

				// Add deleted edges
				if resultEdge.Deleted {
					gr.Delete.Edges = append(gr.Delete.Edges, resultEdge.LegacyKey)
					continue
				}

				// Add edges to merge
				startNodeKey := nodesKeyMap[resultEdge.Start]
				endNodeKey := nodesKeyMap[resultEdge.End]
				customKey := getCustomKey(resultEdge.LegacyKey, resultEdge.Props)

				// Add to legacy index
				gr.Merge.AddEdgeLegacyIndex(resultEdge.LegacyKey, customKey)

				gr.Merge.MergeEdge(
					customKey,
					resultEdge.Label,
					startNodeKey,
					endNodeKey,
					resultEdge.Props,
				)
			}

		}

	}

	return gr, nil

}

// getCustomKey instead of database index key
func getCustomKey(key string, props map[string]interface{}) string {
	customKeyItf, ok := props["key"]
	if ok {
		return customKeyItf.(string)
	}
	return key
}
