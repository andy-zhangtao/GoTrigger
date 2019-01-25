.PHONY: build
name = gt-http
version = v0.1.0

build:
	go build -ldflags "-X main._VERSION_=$(version) -X main._BUILD_=$(shell date +%Y%m%dT%H%M%S)" -o bin/$(name)

run: build
	./$(name)

release: *.go *.md
	GOOS=linux GOARCH=amd64 go build -ldflags "-X main._VERSION_=$(version) -X main._BUILD_=$(shell date +%Y%m%dT%H%M%S)" -a -o bin/$(name)
	docker build -t vikings/$(name):$(version) .
	docker tag vikings/$(name):$(version) vikings/$(name):latest
	docker push vikings/$(name):$(version)
	docker push vikings/$(name):latest
