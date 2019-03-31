package resolve

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/goGraphQL-OTRS/internal/gen/structs"
)

// GenResolve GenResolve
func GenResolve(Name string, res io.Writer, reader *bufio.Reader) {
	fmt.Fprintf(res, `
type Resolver%s struct {
	s Type%s
}
func (R *Resolver%s) Set(s Type%s) {
	R.s = s
}
	`, Name, Name, Name, Name)

	prefix := ""
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
		GenFunc(res, Name, name, prefix+typ)
		prefix = "*"
	}
}

// GenFunc GenFunc
func GenFunc(r io.Writer, Title string, name string, typ string) (int, error) {
	if name == "Id" {
		return fmt.Fprintf(r, `
func (R Resolver%s) %s() string {
	return fmt.Sprintf("%s", R.s.%s)
}
	`, Title, name, "%d", name)
	}

	if typ == "*[]byte" {
		return fmt.Fprintf(r, `
func (R *Resolver%s) %s() *string {
	str:= string(*R.s.%s)
	return &str
}
	`, Title, name, name)
	}
	if typ == "*time.Time" {
		return fmt.Fprintf(r, `
func (R *Resolver%s) %s() *string {
	str:=R.s.%s.String()
	return &str
}
	`, Title, name, name)
	}

	if typ == "*int8" {
		return fmt.Fprintf(r, `
func (R Resolver%s) %s() *int32 {
	i := int32(*R.s.%s)
	return &i
}
	`, Title, name, name)
	}

	if typ == "*int64" {
		return fmt.Fprintf(r, `
func (R *Resolver%s) %s() *string {
	str := fmt.Sprintf("%s", R.s.%s)
	return &str
}
	`, Title, name, "%d", name)
	}

	if typ == "*decimal(10,2)" {
		return fmt.Fprintf(r, `
func (R *Resolver%s) %s() *string {
	str := fmt.Sprintf("%d", R.s.%s)
	return &str
}
	`, Title, name, "%d", name)
	}

	if typ == "int64" {
		return fmt.Fprintf(r, `
func (R Resolver%s) %s() string {
	str:= fmt.Sprintf("%s", R.s.%s)
	return &str
}
	`, Title, name, "%d", name)
	}

	return fmt.Fprintf(r, `
func (R Resolver%s) %s() %s {
	return R.s.%s
}
	`, Title, name, typ, name)
}
