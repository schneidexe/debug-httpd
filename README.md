# debug-httpd

A simple httpd server that helps you debugging data sent to it (esp. headers) e.g. if used behind a web-proxy

## build

```
go get github.com/mitchellh/gox
gox -arch="amd64" -os="darwin linux"
docker build -t debug-httpd .
```

## run

```
./debug-httpd_darwin_amd64
./debug-httpd_linux_amd64
docker run -p 8080:8080 debug-httpd
```

## use

http://localhost:8080