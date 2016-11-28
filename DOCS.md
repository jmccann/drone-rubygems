Use this plugin for publishing gems to a Rubygems server. You can override the
default configuration with the following parameters:

* `api_key` - Rubygems API Key, get it [here](https://rubygems.org/profile/edit)
* `username` - Username, only required without API key
* `password` - Password, only required without API key
* `gemname` - Name of the gem, defaults to repo name
* `gemspec` - Path to gemspec, defaults to repo name with `.gemspec` suffix
* `host` - Rubygems host, only required for a selfhosted gem server
* `skip_cleanup` - Flag to deploy from the current file state

The following secret values can be set to configure the plugin.

* **RUBYGEMS_API_KEY** - corresponds to **api_key**
* **RUBYGEMS_USERNAME** - corresponds to **username**
* **RUBYGEMS_PASSWORD** - corresponds to **password**

It is highly recommended to put the **RUBYGEMS_USERNAME**, **RUBYGEMS_PASSWORD**
or **RUBYGEMS_API_KEY** into secrets so it is not exposed to users. This can be
done using the drone-cli.

```bash
drone secret add --image=plugins/drone-rubygems:latest \
    octocat/hello-world RUBYGEMS_USERNAME kevinbacon

drone secret add --image=plugins/drone-rubygems:latest \
    octocat/hello-world RUBYGEMS_PASSWORD pa55word

drone secret add --image=plugins/drone-rubygems:latest \
    octocat/hello-world RUBYGEMS_API_KEY abcd12345
```

Then sign the YAML file after all secrets are added.

```bash
drone sign octocat/hello-world
```

See [secrets](http://readme.drone.io/0.5/usage/secrets/) for additional
information on secrets

## Example

The following is a sample configuration in your `.drone.yml` file:

```yaml
pipeline:
  rubygems:
    image: plugins/drone-rubygems:latest
    when:
      branch: master
```
