#!/bin/bash

tag="$(git describe --tags $(git rev-list --tags --max-count=1))"
filename="pi-led-$tag-arm64-linux"

env GOOS=linux GOARCH=arm GOARM=6 go build -o release/pi-led example/main.go &&
    cd release && tar -czvf $filename.tar.gz pi-led