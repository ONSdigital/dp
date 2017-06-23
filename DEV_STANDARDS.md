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
* [Clean the campsite](http://programmer.97things.oreilly.com/wiki/index.php/The_Boy_Scout_Rule)
* Cross-browser support/testing
* Accessibility - meets [WCAG 2.0 AA](https://www.w3.org/TR/WCAG20/)

### Testing standards

* Test code should be subject to the same code quality checks and reviews as application code
* Test coverage should be considered by the developer and reviewer, though we should strive to keep it as high as reasonably possible
* All tests need to pass on PR
* Involve members of the QA team where appropriate in the development and review of integration / acceptance tests

##### Unit tests

* Should always be in the same solution as the application code
* Should test only the application code, and not depend on any external services
* Should follow [FIRST principles](https://pragprog.com/magazines/2012-01/unit-tests-are-first)

##### Integration tests

* Integration tests fall under two categories
    1) Test the integration of our code talking to external services
        * Code lives alongside the unit tests, but can run in isolation
        * Defined by the required usage of the external service
    2) Test the integration of external services talking to our code
        * Code lives alongside the acceptance tests, but tests only the interface of a single service
        * Defined by the service interface spec (swagger for API's)
* There should be no tests to specifically test the integration between micro services. The integration between services is tested via acceptance tests.
* Should be isolated and clean up after themselves

##### Acceptance tests

* Test via the public interface of the application (user interface or public API)
* Generally live in a separate dedicated repository
* Defined by application use cases

### Other requirements

* [Setup CI integration](https://github.com/ONSdigital/dp-ci)
* Git-flow
  * `master`/`develop` branches - must be protected
  * `feature/feature-name` branches
  * `hotfix/hotfix-name` branches
  * Use GPG signed annotated release tags
* [Pull request templates](.github/PULL_REQUEST_TEMPLATE.md)
