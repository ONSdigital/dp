Development standards
=====================

These are the Digital Publishing development standards.

### App standards

* Use the [dp-repo-template](https://github.com/ONSdigital/dp-repo-template)
* README
  * App overview and 'Getting started' guide
  * Configuration info
  * Link to [LICENSE](../LICENSE.md) file
  * Link to [contribution guidelines](../guides/CONTRIBUTING.md)
* Change log
  * Don't list every commit
  * Include high-level features/bug fixes
* Repo naming
  * Use a `dp-` prefix for DP-specific repos
  * No 'pet' names, must be descriptive
  * e.g. `dp-frontend-renderer`
* Makefile - `make` target for development

### Coding standards

* Documentation
  * Prefer concise and readable code
  * Include explanatory comments where appropriate
  * Libraries should be fully documented
* API documentation
  * Include a Swagger spec file which documents any HTTP APIs
* Pull requests / code review
* 12-factor
  * All code must be 12-factor compliant - see [https://12factor.net](https://12factor.net) and [our training module](../training/architecture/12_FACTOR_APP_PRINCIPLES.md)
* [Sign commits using GPG](../guides/GPG.md)
* Use meaningful commit messages
* Design patterns and best practices
* Follow existing code style
* [Clean the campsite](https://learning.oreilly.com/library/view/97-things-every/9780596809515/ch08.html) "Always leave the campground cleaner than you found it"
* Cross-browser support/testing
* Accessibility - meets [WCAG 2.0 AA](https://www.w3.org/TR/WCAG20/)

See also [Logging standards](LOGGING_STANDARDS.md).

### Testing standards

* Test code should be subject to the same code quality checks and reviews as application code
* Test coverage should be considered by the developer and reviewer, though we should strive to keep it as high as reasonably possible
* All tests need to pass on PR
* Involve members of the QA team where appropriate in the development and review of integration / acceptance tests
* Any bugs found in the system should have a test created at the lowest level that exercises that error condition

##### Unit tests

* Should always be in the same solution as the application code
* Should test only the application code, and not depend on any external services
* Should follow [FIRST principles](https://web.archive.org/web/20140227191934/http://pragprog.com/magazines/2012-01/unit-tests-are-first)

##### Integration tests

* Generally you should avoid creating integration tests
* The only potential case for integration tests is testing database access code against a test database instance
    * Isolated, not sharing data across tests
    * Coded as unit tests
    * Must be able to run unit tests and integration tests separately (e.g. passing a flag to test runner to run integration tests)
* Any other case for integration tests should be discussed with the team. The tests should provide sufficient ROI (Return on Investment) if they cannot be covered appropriately by unit or acceptance tests
* If you want to test the interface of a service, do so as a unit test at the highest level possible in code. If testing an HTTP API, call the router in a unit test rather than making HTTP requests to a running instance of the service.
* Do not create integration tests that test the configuration of the environment. This is covered in the acceptance tests
* Do not create integration tests that test the integration between micro services. This is covered in the acceptance tests

##### Acceptance tests

* Test via the public interface of the application (user interface or API)
* Acceptance tests should be cover public and private API interfaces
* Generally live in a separate dedicated repository
* Defined by application use cases

### Other requirements

* [Setup CI integration](https://github.com/ONSdigital/dp-ci)
* Git-flow (see [A successful git branching model](https://nvie.com/posts/a-successful-git-branching-model/))
  * `master`/`develop` branches - must be protected
  * `feature/feature-name` branches
  * `fix/fix-name` branches
  * `hotfix/hotfix-name` branches
  * Use GPG signed annotated release tags
* [Pull request templates](../.github/PULL_REQUEST_TEMPLATE.md)
