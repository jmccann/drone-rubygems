# drone-rubygems

Drone plugin for publishing gems to a Rubygems server

## Usage

```
./drone-rubygems <<EOF
{
    "repo": {
        "clone_url": "git://github.com/drone/drone",
        "full_name": "drone/drone",
        "owner": "drone",
        "name": "drone"
    },
    "build": {
        "event": "push",
        "branch": "master",
        "commit": "436b7a6e2abaddfd35740527353e78a227ddcb2c",
        "ref": "refs/heads/master"
    },
    "workspace": {
        "root": "/drone/src",
        "path": "/home/tboerger/Projects/golang/src/github.com/drone-plugins/drone-rubygems"
    },
    "vargs": {
        "api_key": "e123f83113f79eb17308bbf1ca8885bb",
        "name": "my_awesome_tool"
    }
}
EOF
```

## Docker

Build the Docker container using `make`:

```
make deps build
docker build --rm=true -t plugins/drone-rubygems .
```

### Example

```sh
docker run -i plugins/drone-rubygems <<EOF
{
    "repo": {
        "clone_url": "git://github.com/drone/drone",
        "full_name": "drone/drone",
        "owner": "drone",
        "name": "drone"
    },
    "build": {
        "event": "push",
        "branch": "master",
        "commit": "436b7a6e2abaddfd35740527353e78a227ddcb2c",
        "ref": "refs/heads/master"
    },
    "workspace": {
        "root": "/drone/src",
        "path": "/drone/src/github.com/drone/drone"
    },
    "vargs": {
        "api_key": "e123f83113f79eb17308bbf1ca8885bb",
        "name": "my_awesome_tool"
    }
}
EOF
```
