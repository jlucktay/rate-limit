# jlucktay's Rate Limiter (`jrl`)

My implementation of (and small demo wrapper around) a distributed rate limiting library.

## Installation

### Prerequisites

You should have a [working Go environment](https://golang.org/doc/install) and have `$GOPATH/bin` in your `$PATH`.

### Compiling

To download the source, compile, and install the demo binary, run:

``` shell
go get github.com/jlucktay/rate-limit/...
```

The source code will be located in `$GOPATH/src/github.com/jlucktay/rate-limit/`.

A newly-compiled `jrl` binary will be in `$GOPATH/bin/`.

## Usage

Launching the demo server:

``` shell
jrl
```

Hitting the server from another terminal session:

``` shell
$ while true; do curl --silent --include localhost:8080 | head --lines 1; done
HTTP/1.1 200 OK
HTTP/1.1 200 OK
HTTP/1.1 200 OK
...
```

## License

[MIT](https://choosealicense.com/licenses/mit/)
