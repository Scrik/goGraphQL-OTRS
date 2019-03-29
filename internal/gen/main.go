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
	res, err := os.Create("../" + Name + "_type.go")
	if err != nil {
		panic(err)
	}

	// GenStruct
	rd := bytes.NewReader(fields)
	r := bufio.NewReader(rd)
	n := structs.KebabToCamelCase(Name)

	fmt.Fprint(res, `
package api

import "time"
	`)
	structs.GenStruct(n, res, r)

	// GenResolve
	rd = bytes.NewReader(fields)
	r = bufio.NewReader(rd)

	res, err = os.Create("../" + Name + "_resolv.go")
	if err != nil {
		panic(err)
	}

	fmt.Fprint(res, `
package api
	`)
	resolve.GenResolve(n, res, r)
	rd = bytes.NewReader(fields)
	r = bufio.NewReader(rd)
	gql.GenType(n, res, r)
}
