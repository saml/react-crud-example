.PHONY: all run clean

all: build/server build/public/main.js build/templates

build/server: main.go
	go build -o build/server

build/frontend/main.js: frontend/*
	tsc -p .

build/public/main.js: build/frontend/main.js
	webpack

build/templates: templates/*
	rsync -av --delete templates/ build/templates/

run: all
	build/server -templates build/templates -htdocs build/public

clean:
	rm -rf build
