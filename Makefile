COMMIT = $(shell git describe --always)
VERSION = $(shell grep Version paramedic/version.go | sed -E 's/.*"(.+)"$$/\1/')

default: build

build: 
	go build -ldflags "-X github.com/ryotarai/paramedic/paramedic.GitCommit=$(COMMIT)" -o _bin/paramedic .

install: 
	go install -ldflags "-X github.com/ryotarai/paramedic/paramedic.GitCommit=$(COMMIT)" .

buildx:
	gox -ldflags "-X github.com/ryotarai/paramedic/paramedic.GitCommit=$(COMMIT)" -output "_bin/v$(VERSION)/{{.Dir}}_{{.OS}}_{{.Arch}}_$(VERSION)" -arch "amd64" -os "linux darwin" .
	gzip -k _bin/v$(VERSION)/*

test:
	go test -v ./...

bench:
	go test -bench .

release: ensureclean buildx
	git tag v$(VERSION)
	git push origin v$(VERSION)
	ghr v$(VERSION) _bin/v$(VERSION)/

dep:
	dep ensure
	dep status

ensureclean:
	git status | grep -q 'working tree clean'
