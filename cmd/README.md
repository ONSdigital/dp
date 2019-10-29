dp command
==========

Helper command to work with Digital Publishing services

### Getting started

1. Add the following variables to your bash/zsh profile:
    - `DP_SSH_USER` - the user name you use for SSH access, also called `ANSIBLE_USER` in other circumstances
    - `DP_CONFIG` - the path to your YAML config file
2. `go get github.com/ONSdigital/dp/cmd/dp`
3. `dp`

### Configuration

A YAML file is needed to configure potential applications and environments which can be accessed using this tool. The file contains roughly the following structure:

```
setupRepo: github.com/ONSdigital/dp-setup

common_apps:
  - name: concourse
    url: https://concourse.<host>

environment_apps:
  - name: kibana
    url: "https://kibana.$environment.<host>"
  - name: publishing
    url: "https://publishing.$environment.<host>"

environments:
  - name: develop
```

This can be extended to reach a number of apps across a number of environments. If an environment is not included in your config file, you will not be able to access it through the `dp` command.

You can also set an AWS profile per an environment e.g:

```
environments:
  - name: production
    profile: production
```

To use a specific profile in your `~/.aws/credentials` without needing to update your default profile or change your local environment.

### Development

Remember to update the version in [the code](dp/main.go).

### Known issues

* `remote` commands (`allow` and `deny`) get current IP from ipify, but this means a `deny`
    will fail if the IP has changed since running `allow` (for SGs, ACLs use rule numbers)
