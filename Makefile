all: run

run:
	@go run main.go 

build:
	env GOOS=linux go build main.go
	docker build . -t deepsvmwarecom/hello-world -t deepsvmwarecom/hello-world:0.1.0
	rm main
	docker image prune -f
	docker images

docker: build
	docker run --rm -p 8090:8090 -e PORT=8090 deepsvmwarecom/hello-world
