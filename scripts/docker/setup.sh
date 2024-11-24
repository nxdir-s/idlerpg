## Usage: ./scripts/docker/setup.sh <connections> <number of clients> <server ip>
CONNECTIONS=$1
REPLICAS=$2
IP=$3

go build --tags "static netgo" -o ws-client cmd/ws-client/main.go

for (( c=0; c<${REPLICAS}; c++ ))
do
    docker run -l ws-client -v $(pwd)/tmp:/tmp -d alpine /ws-client -conn=${CONNECTIONS} -ip=${IP}
done
