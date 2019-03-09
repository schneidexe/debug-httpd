FROM golang:1.12 as builder
RUN go get -v github.com/schneidexe/debug-httpd

FROM busybox
COPY --from=builder /go/bin/debug-httpd /usr/bin
CMD debug-httpd