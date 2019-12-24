#!/bin/bash

#
# So https://github.com/fernandoocampo/paymentgw
#
OWNER=fernandoocampo
BIN_NAME=paymentgwd
PROJECT_NAME=paymentgw
PLATFORMS=("linux/amd64" "darwin/amd64")

for platform in "${PLATFORMS[@]}"
do
    package="cmd/paymentgwd/main.go"
    platform_split=(${platform//\// })
    GOOS=${platform_split[0]}
    GOARCH=${platform_split[1]}
    output_name=$BIN_NAME'-'$GOOS'-'$GOARCH
    if [ $GOOS = "windows" ]; then
        output_name+='.exe'
    fi  
# go build -o cmd/event_bus/exec cmd/event_bus/main.go
    env GOOS=$GOOS GOARCH=$GOARCH CGO_ENABLED=0 go build -o bin/$output_name $package
    if [ $? -ne 0 ]; then
        echo 'An error has occurred! Aborting the script execution...'
        exit 1
    fi
done