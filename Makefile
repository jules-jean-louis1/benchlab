.PHONY: docker-build docker-launch docker-remove docker-logs docker-ps docker-start docker-stop
docker-build:
	./docker/docker.sh build

docker-launch:
	./docker/docker.sh up -d --build

docker-remove:
	./docker/docker.sh down -v

docker-logs:
	./docker/docker.sh logs -f

docker-ps:
	./docker/docker.sh ps

docker-start:
	./docker/docker.sh start

docker-stop:
	./docker/docker.sh stop