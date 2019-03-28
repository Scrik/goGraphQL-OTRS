package structs

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

// GenStruct GenStruct
func GenStruct(name string, res io.Writer, r *bufio.Reader) {
	fmt.Fprintf(res, `type Type%s struct {
	`, name)

	for {
		line, err := r.ReadString('\n')
		if err != nil {
			break
		}
		ar := strings.Split(line, ";")
		name := KebabToCamelCase(Clean(ar[0]))
		if name == "Field" {
			continue
		}
		json := Clean(ar[0])
		typ := ParseType(Clean(ar[1]))
		TplField(res, name, json, typ)
	}
	fmt.Fprint(res, `}
	`)
}

// ParseType ParseType
func ParseType(t string) string {
	if strings.HasPrefix(t, "varchar") {
		return "string"
	}
	if strings.HasPrefix(t, "text") {
		return "string"
	}
	if strings.HasPrefix(t, "mediumtext") {
		return "string"
	}

	switch t {
	case "smallint(6)":
		return "int8"
	case "bigint(20)":
		return "int64"
	case "datetime":
		return "time.Time"
	case "longblob":
		return "[]byte"
	case "int(11)":
		return "int32"
	default:
		return t
	}
}

// TplField TplField
func TplField(buf io.Writer, name string, json string, t string) (int, error) {
	j := "`" + fmt.Sprintf(`json:"%s"`, json) + "`"
	return fmt.Fprintf(buf, `%s %s %s
	`, name, t, j)
}

// Clean Clean
func Clean(str string) string {
	return str[1 : len(str)-1]
}

// KebabToCamelCase KebabToCamelCase
func KebabToCamelCase(kebab string) (camelCase string) {
	isToUpper := true
	for _, runeValue := range kebab {
		if isToUpper {
			camelCase += strings.ToUpper(string(runeValue))
			isToUpper = false
		} else {
			if runeValue == '_' {
				isToUpper = true
			} else {
				camelCase += string(runeValue)
			}
		}
	}
	return
}
