package neo4jclient

import (
	"fmt"

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

// getCustomKey instead of database index key
func getCustomKey(key string, props map[string]interface{}) string {
	customKeyItf, ok := props["key"]
	if ok {
		return customKeyItf.(string)
	}
	return key
}
