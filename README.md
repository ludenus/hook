# hook
bitbucket webhook handler written in golang based on https://github.com/kataras/iris

## Build
default settings
```
$ go build hook.go
```

small binary
```
$ GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" hook.go
$ upx --ultra-brute ./hook
```

## Run locally

run hook server on default port 80
```
$ ./hook
```

specify custom port number via env var
```
$ hook_LISTENING_ADDRESS="0.0.0.0:7654" ./hook
```

or via command line arg (overrides env var)
```
$ ./hook :8181
```


## Build docker

```
$ docker build -t ludenus/hook:0.0.1 .
```


## Run docker

directly
```
$ docker run -ti --rm --name hook ludenus/hook:0.0.1
```

or via docker-compose
```
$ docker-compose up
```


## Usage

```
$ ./hook :1234
$ curl -XGET http://localhost:1234/ping/blablabla
{"hook":"blablabla"}
```


## Troubleshooting
```
$ ./hook 1234
[ERRO] 2018/09/23 19:56 listen tcp: address 1234: missing port in address
```

Forgot column, should be `:1234` instead of `1234`

## Versions
works on `Ubuntu 18.04.1 LTS`
```
$ uname -a
Linux hostname 4.15.0-34-generic #37-Ubuntu SMP Mon Aug 27 15:21:48 UTC 2018 x86_64 x86_64 x86_64 GNU/Linux
```

```
$ go version
go version go1.10.1 linux/amd64
```

