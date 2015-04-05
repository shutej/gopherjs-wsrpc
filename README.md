gopherjs-test
=============

Tests that GopherJS can speak rpcplus over a WebSocket connection!

To test (for now) do:

    go install ./...
    (cd cmd/gopherjs-test-client; gopherjs build -m main.go; python -m SimpleHTTPServer 8080)
    gopherjs-test-server

Then, browse to [http://127.0.0.1:8080/](http://127.0.0.1:8080/)!