# DUMMY HTTP

This is an experimental project to have a dummy backend when trying out new
stuff. Very simple backend for serving an index.html and json files as fake API.

Serves from index.html from ./public/index.html as a relative path of where you
executed dummyhttp binary.

### Install
Assuming you have a proper go setup

``` bash
go get github.com/efegurkan/dummyhttp
```

### Usage
``` bash
$GOPATH/bin/dummyhttp
```

By default it starts on port 1234
``` bash
$GOPATH/bin/dummyhttp
```

or you can give a specific port number

``` bash
$GOPATH/bin/dummyhttp 5555
```
will serve from localhost:5555


.json files are served from /api/

``` bash
localhost:1234/api/fakeapi/lol.json
```
serves /fakeapi/lol.json
