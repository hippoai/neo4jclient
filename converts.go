package neo4jclient

import (
	"github.com/hippoai/graphgo"
)

// Convert neo4J response to Output type
func ConvertToRGraph(r *Response) ([]*graphgo.Output, error) {

	if len(r.Errors) > 0 {
		return nil, errNeo(r.Errors)
	}

	// Initialize an empty graph response
	outs := []*graphgo.Output{}

	nodesKeyMap := map[string]string{}

	// Loop over all the results, and flatten this into the graph
	for _, result := range r.Results {

		// For each result (== statement), flatten all the rows in the graph
		for _, resultData := range result.Data {

			out := graphgo.NewOutput()

			// First add all the nodes
			for _, resultNode := range resultData.Graph.Nodes {
				customKey := getCustomKey(resultNode.LegacyKey, resultNode.Props)

				nodesKeyMap[resultNode.LegacyKey] = customKey

				// Add deleted nodes
				if resultNode.Deleted {
					out.Delete.LegacyNodes = append(out.Delete.LegacyNodes, resultNode.LegacyKey)
					continue
				}

				// Add to legacy index
				out.Merge.AddNodeLegacyIndex(resultNode.LegacyKey, customKey)

				// Add the node and its properties
				node, _ := out.Merge.MergeNode(customKey, resultNode.Props)

				// Add the labels
				for _, resultLabel := range resultNode.Labels {
					_labelKey := getLabelKey(resultLabel)
					out.Merge.MergeNode(_labelKey, nil)
					out.Merge.MergeEdge(
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
					out.Delete.LegacyEdges = append(out.Delete.LegacyEdges, resultEdge.LegacyKey)
					continue
				}

				// Add edges to merge
				startNodeKey := nodesKeyMap[resultEdge.Start]
				endNodeKey := nodesKeyMap[resultEdge.End]
				customKey := getCustomKey(resultEdge.LegacyKey, resultEdge.Props)

				// Add to legacy index
				out.Merge.AddEdgeLegacyIndex(resultEdge.LegacyKey, customKey)

				out.Merge.MergeEdge(
					customKey,
					resultEdge.Label,
					startNodeKey,
					endNodeKey,
					resultEdge.Props,
				)
			}

			outs = append(outs, out)

		}

	}

	return outs, nil

}
