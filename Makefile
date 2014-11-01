INSTALL_DIR=~/local/bin

all:
	@cat Makefile

build: cargo

cargo:
	go build .

clean:
	\rm -rf cargo

dependency-install:
	go get ./...

godep-install:
	go get github.com/tools/godep

godep-save:
	godep save

install: cargo
	cp cargo $(INSTALL_DIR)

run:
	./cargo cult bootstrap.min.js

b: build
di: dependency-install
gi: godep-install
gs: godep-save
i: install
r: run

.PHONY: build dependency-install godep-install godep-save install run
