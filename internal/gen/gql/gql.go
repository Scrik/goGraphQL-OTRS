package gql

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/goGraphQL-OTRS/internal/gen/structs"
)

// GenGqlType GenGqlType
func GenType(name string, res io.Writer, reader *bufio.Reader) {
	fmt.Fprintf(res, `
func GenGqlType%s(extra string)string{
	return "type %s { " + extra + %s
	`, name, name, "`")
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		ar := strings.Split(line, ";")
		name := structs.KebabToCamelCase(structs.Clean(ar[0]))
		if name == "Field" {
			continue
		}
		typ := ParseType(structs.Clean(ar[1]))
		GenTyp(res, name, typ)
	}
	fmt.Fprintf(res, `}%s
}
	`, "`")
}

// GenTyp GenTyp
func GenTyp(res io.Writer, name string, typ string) {
	fmt.Fprintf(res, `%s: %s,
	`, name, typ)
}

// ParseType ParseType
func ParseType(t string) string {
	if strings.HasPrefix(t, "varchar") {
		return "String"
	}
	if strings.HasPrefix(t, "text") {
		return "String"
	}
	if strings.HasPrefix(t, "longblob") {
		return "String"
	}

	if strings.HasPrefix(t, "mediumtext") {
		return "String"
	}
	if strings.HasPrefix(t, "int") {
		return "Int"
	}

	if strings.HasPrefix(t, "bigint") {
		return "String"
	}

	if strings.HasPrefix(t, "smallint") {
		return "Int"
	}

	switch t {
	case "datetime":
		return "String"
	default:
		return t
	}
}
