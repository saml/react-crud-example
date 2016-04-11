# Quick Start

Assuming you have the following dependencies installed:

- [go](https://golang.org/doc/install)
- [node](https://nodejs.org/) 
- [nvm](https://github.com/creationix/nvm)
- wgo (`go get github.com/skelterjohn/wgo`)
- tsc (`npm install -g typescript`)
- typings (`npm install -g typings`)

Do this once:

```sh
nvm use
typings install
wgo restore
./govendorsymlink.bash
```

Do this to build and run:

```sh
make run
```

# About

React CRUD application example using golang and typescript.