.PHONY: vet linux osx install test release

vet:
	go tool vet *.go marathon/*.go

linux:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bin/terraform-provider-marathon-linux .

osx:
	GOOS=darwin GOARCH=386 go build -o bin/terraform-provider-marathon-osx .

install:
	go install .

test: install
	docker pull python:3
	docker-compose up -d
	sleep 10
	TF_LOG=TRACE TF_LOG_PATH=./test-sh-tf.log TF_ACC=yes MARATHON_URL=http://localhost:8080 go test ./marathon -v
	docker-compose kill
	docker-compose rm -f

release: vet linux osx
	./bin/release.sh
