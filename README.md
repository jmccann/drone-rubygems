# drone-rubygems

[![Build Status](http://beta.drone.io/api/badges/drone-plugins/drone-rubygems/status.svg)](http://beta.drone.io/drone-plugins/drone-rubygems)
[![Coverage Status](https://aircover.co/badges/drone-plugins/drone-rubygems/coverage.svg)](https://aircover.co/drone-plugins/drone-rubygems)
[![](https://badge.imagelayers.io/plugins/drone-rubygems:latest.svg)](https://imagelayers.io/?images=plugins/drone-rubygems:latest 'Get your own badge on imagelayers.io')

Drone plugin to publish ruby gems to a Rubygems server

## Binary

Build the binary using `make`:

```
make deps build
```

### Example

```sh
./drone-rubygems <<EOF
{
    "repo": {
        "clone_url": "git://github.com/drone/drone",
        "owner": "drone",
        "name": "dron",
        "full_name": "drone/drone"
    },
    "system": {
        "link_url": "https://beta.drone.io"
    },
    "build": {
        "number": 22,
        "status": "success",
        "started_at": 1421029603,
        "finished_at": 1421029813,
        "message": "Update the Readme",
        "author": "johnsmith",
        "author_email": "john.smith@gmail.com"
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

## Docker

Build the container using `make`:

```
make deps docker
```

### Example

```sh
docker run -i plugins/drone-rubygems <<EOF
{
    "repo": {
        "clone_url": "git://github.com/drone/drone",
        "owner": "drone",
        "name": "dron",
        "full_name": "drone/drone"
    },
    "system": {
        "link_url": "https://beta.drone.io"
    },
    "build": {
        "number": 22,
        "status": "success",
        "started_at": 1421029603,
        "finished_at": 1421029813,
        "message": "Update the Readme",
        "author": "johnsmith",
        "author_email": "john.smith@gmail.com"
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
