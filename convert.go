package neo4jclient

import (
	"fmt"
	"log"

	"github.com/hippoai/graphgo"
)

// getLabelKey returns the key for a label node
func getLabelKey(label string) string {
	return fmt.Sprintf("label.%s", label)
}

// getLabelRelKey returns the key for the relationship between a node and its label
func getLabelRelKey(nodeKey, label string) string {
	return fmt.Sprintf("%s.%s.%s", nodeKey, graphgo.NODE_LABEL_EDGE_LABEL, label)
}

// Convert neo4J response to Graphgo's format
func Convert(r *Response) (*graphgo.Graph, error) {

	if len(r.Errors) > 0 {
		log.Println("Response error", r.Errors)
		return nil, errNeo(r.Errors)
	}

	// Initialize an empty graph
	g := graphgo.NewEmptyGraph()

	nodesKeyMap := map[string]string{}

	// Loop over all the results, and flatten this into the graph
	for _, result := range r.Results {

		// For each result (== statement), flatten all the rows in the graph
		for _, resultData := range result.Data {

			// First add all the nodes
			for _, resultNode := range resultData.Graph.Nodes {
				customKey := getCustomKey(resultNode.Key, resultNode.Props)

				nodesKeyMap[resultNode.Key] = customKey

				// Add the node and its properties

				node, _ := g.MergeNode(customKey, resultNode.Props)

				// Add the labels
				for _, resultLabel := range resultNode.Labels {
					_labelKey := getLabelKey(resultLabel)
					g.MergeNode(_labelKey, nil)
					g.MergeEdge(
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
				startNodeKey := nodesKeyMap[resultEdge.Start]
				endNodeKey := nodesKeyMap[resultEdge.End]
				customKey := getCustomKey(resultEdge.Key, resultEdge.Props)
				g.MergeEdge(
					customKey,
					resultEdge.Label,
					startNodeKey,
					endNodeKey,
					resultEdge.Props,
				)
			}

		}

	}

	return g, nil

}

// getCustomKey instead of database index key
func getCustomKey(key string, props map[string]interface{}) string {
	customKeyItf, ok := props["key"]
	if ok {
		return customKeyItf.(string)
	}
	return key
}
