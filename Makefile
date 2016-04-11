all:
	go build -o build/server
	tsc -p .
	webpack
	cp frontend/index.html build/

run: all
	build/server -htdocs build
