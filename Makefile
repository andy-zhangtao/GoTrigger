
.PHONY: build
name = gotrigger
version = 0.1.0

build:
	go build -ldflags "-X main._BUILD_=$(shell date +%Y%m%d-%H%M%S) -X main._VERSION_=$(version)" -o bin/$(name)

run: build
	bin/$(name)

test: *.go *.md
	docker rm -f mongo ; echo
	docker run -d --name mongo -p 1301:27017 mongo:3.4
	sleep 5
	export BW_ENV_MONGO_ENDPOINT=127.0.0.1:1301; export BW_ENV_MONGO_DB=trigger; go test -v ./...
	docker rm -f mongo

release: *.go *.md
	docker run -it --rm --name golang -v $$PWD:/go/src/github.com/andy-zhangtao/GoTrigger vikings/golang-onbuild /go/src/github.com/andy-zhangtao/GoTrigger gotrigger
	docker build -t vikings/$(name):$(version) .
	docker push vikings/$(name):$(version)

plugin: plugin/*/*.go
	cd plugin/http; make release