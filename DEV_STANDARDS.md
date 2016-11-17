Development standards
=====================

These are the Digital Publishing development standards.

### App standards

* Use the [dp-repo-template](https://github.com/ONSdigital/dp-repo-template)
* README
  * App overview and 'Getting started' guide
  * Configuration info
  * Link to [LICENSE](LICENSE.md) file
  * Link to [contribution guidelines](CONTRIBUTING.md)
* Change log
  * Don't list every commit
  * Include high-level features/bug fixes
* Repo naming
  * Use a `dp-` prefix for DP-specific repos
  * No 'pet' names, must be descriptive
  * e.g. `dp-frontend-renderer`
* Makefile - `make` target for development
* CodeDeploy scripts

### Coding standards

* Test coverage
* Documentation
  * Prefer concise and readable code
  * Include explanatory comments where appropriate
  * Libraries should be fully documented
* API documentation
  * Include a Swagger spec file which documents any HTTP APIs
* Pull requests / code review
* 12-factor
  * All code must be 12-factor compliant - see [https://12factor.net](https://12factor.net)
* [Sign commits using GPG](GPG.md)
* Use meaningful commit messages
* Design patterns and best practices
* Follow existing code style
* All tests pass on PR

### Other requirements

* [Setup CI integration](https://github.com/ONSdigital/dp-ci)
* Git-flow
  * `master`/`develop` branches - must be protected
  * `feature/feature-name` branches
  * `hotfix/hotfix-name` branches
  * Use GPG signed annotated release tags
* [Pull request templates](.github/PULL_REQUEST_TEMPLATE.md)
