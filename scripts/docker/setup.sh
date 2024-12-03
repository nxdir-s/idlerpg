## Usage: ./scripts/docker/setup.sh <connections> <number of clients> <server ip>
REPLICAS=10

for (( c=0; c<${REPLICAS}; c++ ))
do
    docker run -l ws-client -d ws-client
done
