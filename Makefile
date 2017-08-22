COMMIT = $(shell git describe --always)
VERSION = $(shell grep Version paramedic/version.go | sed -E 's/.*"(.+)"$$/\1/')

default: build

build: 
	go build -ldflags "-X main.GitCommit=$(COMMIT)" -o bin/paramedic .

buildx:
	gox -ldflags "-X main.GitCommit=$(COMMIT)" -output "bin/v$(VERSION)/{{.Dir}}_{{.OS}}_{{.Arch}}_$(VERSION)" -arch "amd64" -os "linux darwin" .

test:
	go test -v $(shell go list ./... | grep -v /vendor/)

bench:
	go test -bench .

release: buildx
	git tag v$(VERSION)
	git push origin v$(VERSION)
	ghr v$(VERSION) bin/v$(VERSION)/

dep:
	dep ensure
	dep status