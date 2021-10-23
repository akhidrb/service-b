FROM alpine:3.14

COPY bin/app /opt/app

RUN chmod +x /opt/app

ENTRYPOINT ["/opt/app"]
