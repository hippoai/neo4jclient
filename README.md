# Neo4J Client

This is a client for the popular graph database [Neo4J](https://neo4j.com).

Features:
* Connects to neo4j transactional endpoint using the REST API. Allows multiple requests per transaction.
* Converts data to [GraphGo](https://github.com/hippoai/graphgo) format, which makes it compatible with the in-memory traversal engine [AskGo](https://github.com/hippoai/askgo)

## Install

`go get -u github.com/hippoai/neo4jclient`

## How to use

### 1. Create a connection

Best practice is to have the following environment variables:
* `DB_USERNAME`: the username for your Neo4J database
* `DB_PASSWORD`: the password for your Neo4J database
* `DB_ENDPOINT`: the transactional endpoint, of the format `http://IP_ADDRESS:PORT/db/data/transaction/commit`

```go
n, err := neo4jclient.NewConnectionFromEnv()
if err != nil {
  log.Println("Could not connect", err)
}
```

### 2. Send a Cypher query and get the result

#### Single query in a transaction

Let's get a random node in the graph.

```go
statement := "MATCH (x) RETURN x LIMIT 1"
props := map[string]interface{}{}
payload := neo4jclient.NewSinglePayload(statement, props)

g, err := n.RequestAndConvert(payload)
if err != nil {
  log.Println("Something went wrong", err)
}
```

**Tip** - Use
```go
func prettyPrint(x interface{}) string {
	b, _ := json.MarshalIndent(x, "", "  ")
	return string(b)
}
```

on the graph instance returned to have a look at it.

#### Multiple queries

Use:

```go
s1 := neo4jclient.NewStatement(
  "MATCH (x) RETURN x LIMIT 1",
  map[string]interface{}{},
)
s2 := neo4jclient.NewStatement(
  "MATCH (x)-[r]-(y) RETURN r LIMIT 1",
  map[string]interface{}{},
)
g, err := n.RequestAndConvert(neo4jclient.NewPayload(s1, s2))
if err != nil {
  log.Println("Something went wrong", err)
}

prettyPrint(g)
```
