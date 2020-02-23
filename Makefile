all: run

run:
	@go run main.go 

build:
	docker build . -t deepsvmwarecom/hello-world
	docker image prune -f
	docker images

docker: build
	docker run --rm -p 8090:8090 -e STATSD_ADDR=host.docker.internal:8125 deepsvmwarecom/hello-world
