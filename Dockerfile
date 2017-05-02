FROM alpine:3.2

RUN apk update && \
  apk add \
    ca-certificates \
    git \
    ruby && \
  gem install --no-ri --no-rdoc \
    dpl && \
  rm -rf /var/cache/apk/*

ADD drone-rubygems /bin/
ENTRYPOINT ["/bin/drone-rubygems"]
