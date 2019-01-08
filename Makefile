
.PHONY: build
name = gotrigger
version = 0.1.0

build:
	go build -ldflags "-X main._BUILD_=$(shell date +%Y%m%d-%H%M%S) -X main._VERSION_=$(version)" -o bin/$(name)

run: build
	bin/$(name)

release: *.go *.md
	docker run -it --rm --name golang -v $$PWD:/go/src/github.com/andy-zhangtao/GoTrigger vikings/golang-onbuild /go/src/github.com/andy-zhangtao/GoTrigger gotrigger
	docker build -t vikings/$(name) .
	docker push vikings/$(name)
