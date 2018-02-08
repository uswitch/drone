# docker build --rm -t drone/drone .

FROM alpine:3.7
RUN apk add --update --no-cache ca-certificates \
 && mkdir -p /etc/drone \
 && rm -rf /var/cache/apk/*

EXPOSE 8000 9000 80 443

ENV DRONE_ENV_FILE=/etc/drone/env
ENV DATABASE_DRIVER=sqlite3
ENV DATABASE_CONFIG=/var/lib/drone/drone.sqlite
ENV GODEBUG=netdns=go
ENV XDG_CACHE_HOME /var/lib/drone

ADD release/drone-server entrypoint.sh /bin/

ENTRYPOINT ["/bin/entrypoint.sh"]
