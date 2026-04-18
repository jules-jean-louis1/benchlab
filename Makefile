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

docker-rebuild:
	./docker/docker.sh stop $(s) || true
	./docker/docker.sh rm -f $(s) || true
	./docker/docker.sh up -d --build $(s)

protoc-gen:
	protoc -I=proto --go_out=grpc-service --go_opt=module=grpc-sensor-service --go-grpc_out=grpc-service --go-grpc_opt=module=grpc-sensor-service proto/sensor/sensor.proto

k6-test:
	k6 run benchmarks/scripts/$(s)

k6-run-all:
	k6 run benchmarks/scripts/create-sensor-write.test.js \
	k6 run benchmarks/scripts/get-sensor-unit-read.test.js \
	k6 run benchmarks/scripts/ramp-up-load.test.js \
	k6 run benchmarks/scripts/get-sensor-1000.test.js \
	k6 run benchmarks/scripts/pagination-1000-sensors.test.js \
