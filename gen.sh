#!/bin/bash

function generate(){
    (
        cd api/$1;
        go run  ../../internal/gen/main.go -csv $1.csv
        gofmt -w *.go
    )
}

generate ticket
generate article
generate article_attachment
generate customer_company

(
    cd api/
    gofmt -w *.go
)
