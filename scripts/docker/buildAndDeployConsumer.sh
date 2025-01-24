docker build . --no-cache -t idlerpg-consumer -f cmd/consumer/Dockerfile
docker tag idlerpg-consumer:latest idlerpg/consumer:latest
docker push idlerpg/consumer:latest
