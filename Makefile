INSTALL_DIR=~/local/bin

SRCS = cargo.go copy.go

all:
	@cat Makefile

build: cargo

cargo: $(SRCS)
	go build .

clean:
	\rm -rf cargo

distclean: clean
	\rm -rf Godeps/_workspace

dependency-install:
	go get ./...

godep-install:
	go get github.com/tools/godep

godep-save:
	godep save

install: $(INSTALL_DIR)/cargo

$(INSTALL_DIR)/cargo: cargo
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
