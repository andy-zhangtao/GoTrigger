
.PHONY: build
name = gotrigger
version = 0.1.0

build:
	go build -ldflags "-X main._BUILD_=$(shell date +%Y%m%d-%H%M%S) -X main._VERSION_=$(version)" -o bin/$(name)

run: build
	bin/$(name)

release: *.go *.md
	GOOS=linux GOARCH=amd64 go build -ldflags "-X main._BUILD_=$(shell date +%Y%m%d-%H%M%S) -X main._VERSION_=$(version)" -a -o bin/$(name)
	docker build -t vikings/$(name) .
	docker push vikings/$(name)
