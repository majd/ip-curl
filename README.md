# ip-curl

[![License](https://img.shields.io/badge/License-MIT-yellow.svg)](https://github.com/majd/ip-curl/blob/main/LICENSE)
![Go](https://img.shields.io/badge/Go-1.17-blue.svg)

A lightweight tool that starts a web server and returns the client's IP address on every request.

```
$ curl ip.example.com
123.1.123.123
```

## Usage

The tool can be ran in a Docker container or compiled directly on the host.

### Docker (recommended)

Clone repository
```
$ git clone git@github.com:majd/ip-curl.git && cd ip-curl
```

Build image

```
$ docker build .
```

Start container

```
$ docker run --rm -p 3000:3000 <sha256>
```


### Manual

Clone repository
```
$ git clone git@github.com:majd/ip-curl.git && cd ip-curl
```

Compile code
```
$ go mod download
$ go build
```

Run compiled binary
```
$ ./ip-curl
```

## License

ip-curl is released under the [MIT license](https://github.com/majd/ip-curl/blob/main/LICENSE).