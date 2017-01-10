package neo4jclient

// RequestAndConvert requests the database
// and converts it to Graphgo format
func (neo *Neo) RequestAndConvert(payload *Payload) (*Output, error) {

	response, err := neo.Request(payload)

	if err != nil {
		return nil, err
	}

	return Convert(response)

}
