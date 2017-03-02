.PHONY: vet linux osx build test release

vet:
	go tool vet *.go marathon/*.go

linux:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bin/terraform-provider-marathon-linux .

osx:
	GOOS=darwin GOARCH=386 go build -o bin/terraform-provider-marathon-osx .

install:
	go install .

os=$(shell uname)
ifeq ($(os),Darwin)
  docker_compose_file=docker-compose.yml
else
  docker_compose_file=docker-compose-linux.yml
endif

test: install
	docker pull python:3
	docker-compose -f $(docker_compose_file) up -d
	sleep 10
	TF_LOG=TRACE TF_LOG_PATH=./test-sh-tf.log TF_ACC=yes MARATHON_URL=http://dev.banno.com:8080 go test ./marathon -v
	docker-compose -f $(docker_compose_file) kill
	docker-compose -f $(docker_compose_file) rm -f

release: vet linux osx
	./bin/release.sh
