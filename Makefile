build: dist/brutemq_amd64

test:
	@bash test/*.sh

gofmt:
	@gofmt -s -w .

clean:
	@rm dist/*

dist/brutemq_amd64:
	@DOCKER_BUILDKIT=1 docker build -f build/binary.dockerfile --target binary --output dist/ .

all: build test
.PHONY: test clean gofmt
