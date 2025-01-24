docker build . --no-cache -t idlerpg-gameserver -f cmd/server/Dockerfile
docker tag idlerpg-gameserver:latest idlerpg/gameserver:latest
docker push idlerpg/gameserver:latest
