package neo4jclient

const (
	ENV_DB_USERNAME = "DB_USERNAME"
	ENV_DB_PASSWORD = "DB_PASSWORD"
	ENV_DB_ENDPOINT = "DB_ENDPOINT"

	ERR_DB_CONNECT = "No DB username or password provided"
	ERR_NEO        = "Errors in the query"

	RESULT_DATA_CONTENTS = "graph"

	REQUEST_METHOD       = "POST"
	HEADER_AUTHORIZATION = "Authorization"
)
