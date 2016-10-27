.PHONY: vet linux osx build test release

vet:
	go tool vet *.go marathon/*.go

linux:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bin/terraform-provider-marathon-linux .

osx:
	GOOS=darwin GOARCH=386 go build -o bin/terraform-provider-marathon-osx .

build: vet osx linux
	go install .

test: build
	docker pull python:3
	docker-compose build test
	docker-compose pull
	docker-compose up -d marathon mesos-master mesos-slave zookeeper
	sleep 5
	docker-compose up test
	docker-compose kill
	docker-compose rm -f

release:
	./bin/release.sh
