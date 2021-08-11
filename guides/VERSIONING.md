# Version control of Repositories

Following [Semantic Versioning 2.0.0](https://semver.org/).
* The first stable release for repositories, either libraries or services, is `1.0.0`

* When creating release branches or tags it is digital publishing standard to not increment the patch version and hence should always be zero; (`Major.Minor.Patch = 1.0.0`)

* Any new release will force the version to go up by either a minor or major version, see examples below:

```
Current version is 1.1.0

* Minor release updates version to 1.2.0
* Major release updates version to 2.0.0
```

## Maintaining Old Versions

When upgrading to a new major version, e.g. `1.x` to `2.x` it is Digital Publishing standard to cut a `v1` branch for maintenance of the old major version. Then the `main` branch will continue to contain the latest code for the newest major version (in this scenario v2).

The purposes of maintaing old versions could be for the following reasons:

- bug fixes
- security fixes
- dependency upgrades

## Pre-release versioning (Alpha and Beta)

When going to a new major version of an app or library it is advisable to move to a pre-release version using the suffix `-alpha` or `-beta` e.g. `2.0.0-beta`. 
This allows for further changes that might be breaking backward compatibility without moving quickly to another version.

Any new tag releases should follow the convention to add `.<number>` with the number incrementing by 1 for each new tag, like so `v2.0.0-beta.2`.

See below example of order of increasing versions:

```
1.3.2 < 2.0.0-alpha < 2.0.0-alpha.1 < 2.0.0-beta < 2.0.0-beta.2 < 2.0.0-beta.11 < 2.0.0-rc.1 < 2.0.0
```

Once the team are happy with the new app/library then the new major version (e.g. `2.0.0`) should be created and this signals that the code must maintain backward compatibility - until any later breaking forces a move to a new major version.

It is worth referring back to the [semantic versioning documentation (2.0)](https://semver.org/) for more details on pre-release versions.
