# drone-rubygems

Drone plugin for publishing Ruby gems

## Docker

Build the Docker container

```sh
docker build --rm=true -t plugins/drone-rubygems .
```

Build and Publish a Ruby gems

```sh
docker run -i plugins/drone-rubygems <<EOF
{
    "workspace": {
        "path": "/drone/src"
    },
    "build" : {
        "number": 1,
        "head_commit": {
            "sha": "9f2849d5",
            "branch": "master",
            "ref": "refs/heads/master"
        }
    },
    "vargs": {
        "host": "https://rubygems.org",
        "api-key": "SOME_API_KEY",
        "gem": "awesomegem",
        "skip-cleanup": true
    }
}
EOF
```

## Development

Just install gem `dpl`

```sh
gem install dpl
```

Test plugin

```sh
ruby dpl.rb <<EOF
{
    "workspace": {
        "path": "/drone/src"
    },
    "build" : {
        "number": 1,
        "head_commit": {
            "sha": "9f2849d5",
            "branch": "master",
            "ref": "refs/heads/master"
        }
    },
    "vargs": {
        "host": "https://rubygems.org",
        "api-key": "SOME_API_KEY",
        "gem": "awesomegem",
        "skip-cleanup": true
    }
}
EOF
```
