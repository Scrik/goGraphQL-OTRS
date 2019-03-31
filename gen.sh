#!/bin/bash

function generate(){
    (
        ../../.bin/gen -csv $1
    )
}
(
    cd internal/gen/
    go build
)
mv internal/gen/gen .bin/
(
    cd api/csv
    for file in *.csv
    do
        generate $file
    done
    cd ..
    gofmt -w *.go
)
