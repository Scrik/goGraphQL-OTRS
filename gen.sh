#!/bin/bash

function generate(){
    (
        cd api/$1;
        go run  ../../internal/gen/main.go -csv $1.csv
        gofmt -w $1.go
    )
}

generate tiket
# generate article
# generate article_attachment
