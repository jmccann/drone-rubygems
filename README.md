# drone-rubygems

[![Build Status](http://beta.drone.io/api/badges/jmccann/drone-rubygems/status.svg)](http://beta.drone.io/jmccann/drone-rubygems)

Drone plugin to publish ruby gems to a Rubygems server. For the usage information and a listing of the available options please take a look at [the docs](DOCS.md).

## Binary

Build the binary using `make`:

```
make deps build
```

## Usage

Execute from the working directory:

```
docker run --rm \
  -e RUBYGEMS_API_KEY=abcd12345 \
  -e PLUGIN_REPO=octocat/hello-world \
  -v $(pwd):$(pwd) \
  -w $(pwd) \
  plugins/drone-rubygems --help
```
