package neo4jclient

import "github.com/hippoai/graphgo"

// RequestAndConvert requests the database
// and converts it to Graphgo format
func (neo *Neo) RequestAndConvert(payload *Payload) (*graphgo.Graph, error) {

	response, err := neo.Request(payload)

	if err != nil {
		return nil, err
	}

	return Convert(response)

}
