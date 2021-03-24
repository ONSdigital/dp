Dependency Upgrading Standard
=============================

## General approach

When working on a repository, developers should check that the dependencies of that repository are sufficiently up to
date. The approach to updating the dependencies depends on the type of app and the underlying technology.

There is no formal decision process for upgrading language dependencies. As soon as a newer minor version of a language
is released (providing it is covered by the policy for that language), then you may upgrade it straight away. For 
example, when go@1.x.1 is released then you should use it in your next pull request.

If you notice any upgrades that are not covered by this policy (including major versions of DBs becoming available)
raise them on the 20% board for triage.

We don’t update dependencies in legacy apps unless there is a security vulnerability (babbage, zebedee, train, brian,
legacy florence).

## Go based apps and libraries

Go based repositories have their dependencies defined in three different ways.
1. The version of go used to test and build the component (the docker image tag defined in `ci/build.yml` etc. eg. `1.15.8`)
2. The language specification of the module (defined in `go.mod` eg. `1.15`)
3. The versions of each of the module dependencies (defined in `go.mod` and `go.sum`)

### Build and test version (docker image)

The general approach for updating the version of go used to build an app is to update to the latest version of go as
defined [on the `golang` image on docker hub](https://hub.docker.com/_/golang) as long as the patch version is >= 1. 
For example…

| Available docker image tags | Choose… |
| --------------------------- | ------- |
| 1.15.8 , 1.16.0             | 1.15.8  |
| 1.15.9 , 1.16.1             | 1.16.1  |
| 1.15.9 , 1.16.2             | 1.16.2  |

Once the new version of go has been selected it needs to be added to the `ci/build.yml` and `ci/unit.yml` files.

NB. For libraries the version of go will be pinned by the app using the library so for the ci jobs we want to use the 
latest version of go instead. So for libraries only, the tag in the above files should be `latest` instead of a specific
version. 

### Language specification (go.mod)

If the docker tag is upgraded from one minor version to another (eg. from `1.15.x` to `1.16.x`) then the language
specification of the module in `go.mod` should be upgraded as well. Whilst not strictly necessary it helps to keep
things in sync. Note that this must not contain the patch version so should simply be `1.16` etc. not `1.16.2` and can
be updated with the command `go mod edit -go=1.16` or similar.

### Module dependencies

The module dependencies are audited by running `make audit` and also by an audit CI job triggered by PRs. Any vulnerable
dependencies identified should be updated. Other dependencies may be updated as necessary eg. in order to use new
functionality.

To update a specific dependency run `go get -u path.to/some/library` or for a specific version 
`go get path.to/some/library@v1.2.3`. If you want to update all direct dependencies simply run `go get -u ./...`.

This will not update sub-dependencies of dependencies (unless a new version of a dependency requires a new version of a 
sub-dependency). So if after updating there are still vulnerabilities in any of these transitive dependencies there are
two possibilities for resolution.

1. (Preferred) If the vulnerable dependency is being brought in by one of our libraries, then fix the vulnerability in
   the library first and release a new version. This dependency can then be in turn updated in the app using it.
2. If it is not possible to update the direct dependency (say it is an external library which itself has no newer
   version) then you can directly update the vulnerable dependency using `go get path.to/some/library@v1.2.3`

## Java based apps and libraries

latest LTS and wait for a patch (update installing to go from https://adoptopenjdk.net/)


## Other technologies

### Node 
latest ‘active’ LTS and wait for a patch

### React 
latest minor version and wait for a patch

### Spring boot 
latest minor and wait for a patch

### Frameworks and libraries
latest minor + patch as clean campsite. Down to senior developers to determine when major version upgrades are worth doing

### Platform technologies 
latest + patch

### Databases 
stay up to date on patches, but all minor/major updates will need discussion and planning due to dependency with clients.

### AWS 
managed should just be latest


