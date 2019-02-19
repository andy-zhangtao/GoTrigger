.PHONY: build
name = bwidow

test: *.go *.md
	docker rm -f postgres ; echo
	docker run -d --name postgres -e POSTGRES_PASSWORD=123456 -p 5432:5432 postgres:alpine
	sleep 5
	docker cp u1.sql postgres:/tmp/u1.sql
	docker exec -u postgres -it postgres psql -f /tmp/u1.sql
	go test github.com/andy-zhangtao/bwidow
	docker rm -f postgres
all: test
