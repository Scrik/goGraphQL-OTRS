package resolve

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/goGraphQL-OTRS/internal/gen/structs"
)

// GenResolve GenResolve
func GenResolve(name string, res io.Writer, reader *bufio.Reader) {
	fmt.Fprintf(res, `
type Resolver struct {
	s Type%s
}
func (R *Resolver) Set(s Type%s) {
	R.s = s
}
	`, name, name)

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
		typ := structs.ParseType(structs.Clean(ar[1]))
		GenFunc(res, name, typ)
	}
}

// GenFunc GenFunc
func GenFunc(r io.Writer, name string, typ string) (int, error) {
	if typ == "[]byte" {
		return fmt.Fprintf(r, `
func (R *Resolver) %s() string {
	return string(R.s.%s)
}
	`, name, name)
	}
	if typ == "time.Time" {
		return fmt.Fprintf(r, `
func (R *Resolver) %s() string {
	return R.s.%s.String()
}
	`, name, name)
	}

	if typ == "int8" {
		return fmt.Fprintf(r, `
func (R *Resolver) %s() int32 {
	return int32(R.s.%s)
}
	`, name, name)
	}

	if typ == "int64" {
		return fmt.Fprintf(r, `
func (R *Resolver) %s() string {
	return fmt.Sprintf("%s", R.s.%s)
}
	`, name, "%d", name)
	}
	return fmt.Fprintf(r, `
func (R *Resolver) %s() %s {
	return R.s.%s
}
	`, name, typ, name)
}
