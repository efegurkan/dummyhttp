# DUMMY HTTP

This is an experimental project to have a dummy backend when trying out new
stuff. Very simple backend for serving an index.html and json files as fake API.

Serves from index.html from ./public/index.html as a relative path to the
dummyhttp binary.

``` bash
./dummyhttp
```
to serve from localhost:1234

or

``` bash
./dummyhttp portnumber
```
to serve from localhost:portnumber.


.json files are also served from /api/

``` bash
localhost:1234/api/fakeapi/lol.json
```
serves ./fakeapi/lol.json
