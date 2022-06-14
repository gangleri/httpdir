# httpdir

A simple http server written in Go inspired by Python's [`python -m http.server`](https://docs.python.org/3/library/http.server.html) to serve a directory over http. 

## Installation
```
go get -u github.com/httpdir/cmd/httpdir
```

## Usage
```
httpdir [-p port {8080}] [path to serve]
```

- The `--port` flag can be used to server on a prt other than the default of `8080`

- Running without an argument will serve the currect directory. Otherwise the directory at the supplied path will be served


### Purpose
Sometimes I've found it useful to spin up an http sever, I had a previous version written only using Go's stdlib but this was extended to play with cobra overkill and under utilisation of cobra for this but it was a play thing :shrug: :grin: