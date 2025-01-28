docker build . --no-cache -t idlerpg-userevents -f cmd/consumers/userevents/Dockerfile
docker tag idlerpg-userevents:latest idlerpg/userevents:latest
docker push idlerpg/userevents:latest
