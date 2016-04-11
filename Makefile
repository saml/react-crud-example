all:
	go build -o build/server
	tsc -p .
	webpack
	rsync -av --delete templates/ build/templates/

run: all
	build/server -templates build/templates -htdocs build/public

clean:
	rm -rf build
