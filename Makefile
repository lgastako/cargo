all:
	@cat Makefile

dependency-install:
	go get ./...

godep-install:
	go get github.com/tools/godep

godep-save:
	godep save

build:
	go build .

run:
	./cargo cult bootstrap.min.js

b: build
di: dependency-install
gi: godep-install
gs: godep-save
r: run
