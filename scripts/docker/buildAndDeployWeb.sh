docker build . --no-cache -t idlerpg-webserver -f cmd/web/Dockerfile
docker tag idlerpg-webserver:latest idlerpg/webserver:latest
docker push idlerpg/webserver:latest
