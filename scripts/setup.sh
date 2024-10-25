#!/bin/bash
## Usage: ./connect <connections> <number of clients> <server ip>

CONNECTIONS=$1
REPLICAS=$2
IP=$3
go build --tags "static netgo" -o client client.go
for (( c=0; c<${REPLICAS}; c++ ))
do
    docker run -l 1m-go-websockets -v $(pwd)/client:/client -d alpine /client -conn=${CONNECTIONS} -ip=${IP}
done
