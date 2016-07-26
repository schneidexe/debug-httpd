FROM busybox

ADD debug-httpd_linux_amd64 /debug-httpd

RUN chmod a+x /debug-httpd

CMD /debug-httpd