# Docker image for the Drone Rubygems plugin
#
#     cd $GOPATH/src/github.com/drone-plugins/drone-rubygems
#     make deps build
#     docker build --rm=true -t plugins/drone-rubygems .

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
