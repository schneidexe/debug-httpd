# debug-httpd

A simple httpd server that helps you debugging data sent to it (esp. headers) e.g. if used behind a web-proxy

## build
```
docker build --no-cache -t debug-httpd .
```

## run
```
docker run -p 8080:8080 debug-httpd
```

## use

[http://localhost:8080](http://localhost:8080)
