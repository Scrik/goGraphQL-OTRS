package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/goGraphQL-OTRS/internal/gen/gql"
	"github.com/goGraphQL-OTRS/internal/gen/resolve"
	"github.com/goGraphQL-OTRS/internal/gen/structs"
)

func main() {
	var basename = flag.String("csv", "tiket.csv", "csv file name")
	flag.Parse()
	fmt.Println(*basename)
	Name := strings.TrimSuffix(*basename, filepath.Ext(*basename))
	fields, err := ioutil.ReadFile(*basename)
	if err != nil {
		panic(err)
	}
	res, err := os.Create(Name + ".go")
	if err != nil {
		panic(err)
	}

	rd := bytes.NewReader(fields)
	r := bufio.NewReader(rd)
	n := structs.KebabToCamelCase(Name)

	fmt.Fprintf(res, `
package %s

	`, Name)
	structs.GenStruct(n, res, r)
	rd = bytes.NewReader(fields)
	r = bufio.NewReader(rd)
	resolve.GenResolve(n, res, r)
	rd = bytes.NewReader(fields)
	r = bufio.NewReader(rd)
	gql.GenType(n, res, r)
}
