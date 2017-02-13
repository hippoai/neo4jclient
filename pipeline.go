package neo4jclient

import "github.com/hippoai/graphgo"

// RequestAndConvert requests the database
// and converts it to Graphgo format
func (neo *Neo) RequestAndConvert(payload *Payload) (*graphgo.Output, error) {

	response, err := neo.Request(payload)

	if err != nil {
		return nil, errNeo4JRequest()
	}

	return Convert(response)

}
