FROM alpine:edge

RUN apk --update add \
  ca-certificates \
  ruby \
  ruby-bundler \
  ruby-json \
  git \
  ruby-dev && \
  gem install dpl --no-ri --no-rdoc && \
  rm -fr /usr/share/ri

ADD dpl.rb /dpl.rb
ENTRYPOINT ["/usr/bin/ruby", "/dpl.rb"]
