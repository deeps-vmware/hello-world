all: run

run:
	@go run main.go 

build:
	docker build . -t deepsvmwarecom/hello-world -t deepsvmwarecom/hello-world:0.1.0
	docker image prune -f
	docker images

docker: build
	docker run --rm -p 8090:8090 -e PORT=8090 deepsvmwarecom/hello-world
