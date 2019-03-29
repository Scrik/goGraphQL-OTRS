package main

import (
	"flag"
	"io/ioutil"
	"log"

	"github.com/goGraphQL-OTRS/internal/app"
)

func main() {
	var debug = flag.Int("d", 0, "debug level")
	var filename = flag.String("c", "config.yml", "config file name")
	var gen = flag.Bool("gen", false, "generate default config file")
	flag.Parse()
	if *debug > 5 {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
	}
	if *gen {
		app.Generate(*filename)
		return
	}
	data, err := ioutil.ReadFile(*filename)
	if err != nil {
		panic(err)
	}
	_, err = app.New(data)
}
