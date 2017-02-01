package neo4jclient

import (
	"fmt"
	"log"
	"strings"

	"github.com/hippoai/goerr"
)

// compressQuery removes spaces and blank lines
func compressQuery(q string) string {
	q1 := strings.Replace(q, " ", "", -1)
	q2 := strings.Replace(q1, "\n", "", -1)
	return q2
}

// insertProp replaces a property placeholder by its value
func insertProp(q, key string, value interface{}) string {
	valueStr, isString := value.(string)
	if isString {
		return strings.Replace(
			q,
			fmt.Sprintf("{props}.%s", key),
			fmt.Sprintf("\"%s\"", compressQuery(valueStr)),
			-1,
		)
	} else {
		return strings.Replace(
			q,
			fmt.Sprintf("{props}.%s", key),
			fmt.Sprintf("%v", value),
			-1,
		)
	}
}

// IsSameQuery assesses whether two queries are the same
func IsSameQuery(expected string, statement *Statement) error {
	cExpected := compressQuery(expected)
	cOriginal := compressQuery(statement.Statement)

	if statement.Parameters.Props != nil {
		for key, value := range statement.Parameters.Props {
			cOriginal = insertProp(cOriginal, key, value)
		}
	}

	if cExpected != cOriginal {
		log.Println(cExpected)
		log.Println(cOriginal)
	}

	if cExpected != cOriginal {
		return ErrSameQuery(cExpected, cOriginal, expected, statement)
	}

	return nil
}

func ErrSameQuery(cExpected, cOriginal, expected string, statement *Statement) error {
	return goerr.New(
		"ERR_NOT_SAME_QUERY",
		map[string]interface{}{
			"expected":  expected,
			"statement": statement.Statement,
			"props":     statement.Parameters.Props,
			"cExpected": cExpected,
			"cOriginal": cOriginal,
		},
	)
}
