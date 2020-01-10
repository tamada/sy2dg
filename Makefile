GO=go
NAME := sy2dg
VERSION := 1.0.0
DIST := $(NAME)-$(VERSION)

all: test build

update_version:
	@for i in README.md; do\
	    sed -e 's!Version-[0-9.]*-yellowgreen!Version-${VERSION}-yellowgreen!g' -e 's!tag/v[0-9.]*!tag/v${VERSION}!g' $$i > a ; mv a $$i; \
	done

	@sed 's/const Version = .*/const Version = "${VERSION}"/g' sy2dg.go > a
	@mv a sy2dg.go
	@echo "Replace version to \"${VERSION}\""

setup: update_version
	git submodule update --init

test: build setup
	$(GO) test -covermode=count -coverprofile=coverage.out $$(go list ./... | grep -v wasm)

build: setup
	go build -o sy2dg cmd/sy2dg/main.go
	go build -o ksu2json cmd/ksu2json/main.go
	go build -o json2dg cmd/json2dg/main.go

define _createDist
	mkdir -p dist/$(1)_$(2)/$(DIST)
	GOOS=$1 GOARCH=$2 go build -o dist/$(1)_$(2)/$(DIST)/$(NAME) cmd/$(NAME)/main.go
	GOOS=$1 GOARCH=$2 go build -o dist/$(1)_$(2)/$(DIST)/json2dg  cmd/json2dg/main.go
	GOOS=$1 GOARCH=$2 go build -o dist/$(1)_$(2)/$(DIST)/ksu2json cmd/ksu2json/main.go
	cp -r data README.md LICENSE dist/$(1)_$(2)/$(DIST)
	mkdir -p dist/$(1)_$(2)/$(DIST)/djview
	cp docs/draw_graph_d3.js docs/index.html docs/style.css docs/dataset.js dist/$(1)_$(2)/$(DIST)/djview
	tar cfz dist/$(DIST)_$(1)_$(2).tar.gz -C dist/$(1)_$(2) $(DIST)
endef

dist: build
	@$(call _createDist,darwin,amd64)
	@$(call _createDist,darwin,386)
	@$(call _createDist,windows,amd64)
	@$(call _createDist,windows,386)
	@$(call _createDist,linux,amd64)
	@$(call _createDist,linux,386)

run: build
	go run cmd/sy2dg/main.go --url=https:/syllabus.kyoto-su.ac.jp/syllabus/html/2019/ data/cse_ksu | tee docs/dataset.js

install: test build
	$(GO) install $(LDFLAGS)

clean:
	$(GO) clean
	rm -rf sy2dg ksu2json json2dg coverage.out dist
