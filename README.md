gopherjs-test
=============

Tests that GopherJS can speak rpcplus over a WebSocket connection!

To test (for now) do:

    go install ./...
    (cd cmd/gopherjs-test-client; gopherjs build -m main.go; python -m SimpleHTTPServer 8080)
    gopherjs-test-server

Then, browse to [http://127.0.0.1:8080/](http://127.0.0.1:8080/)!

TODO:

- [ ] Watch static file(s) that [watches for changes](gopkg.in/fsnotify.v1) and
  [compresses](http://golang.org/pkg/compress/gzip/#NewWriterLevel).
  - https://gist.github.com/the42/1956518
- [ ] Cache-Control:public/ETag/If-None-Match/304 Not Modified
  - http://en.wikipedia.org/wiki/HTTP_ETag#Typical_usage
