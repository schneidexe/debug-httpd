FROM golang:1.12 as builder
RUN CGO_ENABLED=0 go get -v github.com/schneidexe/debug-httpd

FROM alpine
RUN apk add --no-cache ca-certificates && \
    update-ca-certificates
COPY --from=builder /go/bin/debug-httpd /debug-httpd
CMD /debug-httpd