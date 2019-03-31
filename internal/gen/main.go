package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path"
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
	p := path.Join("../" + Name + "_type.go")
	_, err = os.Stat(p)
	if !os.IsNotExist(err) {
		err = os.Remove(p)
	}

	res, err := os.Create(p)
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
	p = path.Join("../" + Name + "_resolv.go")
	_, err = os.Stat(p)
	if !os.IsNotExist(err) {
		err = os.Remove(p)
	}
	res, err = os.Create(p)
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
