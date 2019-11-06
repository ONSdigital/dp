Digital Publishing Platform
===========================

The digital publishing platform is a collection of services which provide
content management and publishing functionality for ONS content, and is
responsible for the public ONS website.

Guides and general info
-----------------------

* [Getting started](guides/GETTING_STARTED.md)
* [Release process](guides/RELEASES.md)
* [Contributing guidelines](guides/CONTRIBUTING.md)
* [GPG commit signing](guides/GPG.md)
* [Port numbering](guides/PORTS.md)
* [Technologies](guides/TECHNOLOGIES.md)
* [Tagging a release](guides/TAGS.md)
* [Migrating Go vendor projects to Go modules](guides/MODULES.md)

Standards and specifications
----------------------------

* [Development standards](standards/DEV_STANDARDS.md)
* [Definition of Ready](standards/DEFINITION_OF_READY.md)
* [Definition of Done](standards/DEFINITION_OF_DONE.md)
* [Digital Publishing principles](https://github.com/ONSdigital/dp-principles)
* [API standards](standards/API_STANDARDS.md)
* [Logging standards](standards/LOGGING_STANDARDS.md)
* [Health check specification](standards/HEALTH_CHECK_SPECIFICATION.md)

Dev tooling
-----------

**[dp tool](cmd/README.md)**

A command line utility to help make working in Digital Publishing simpler.

**[websysd](websysd/README.md)**

You can use websysd to run your DP apps locally.

**[paasbox](paasbox)**

Paasbox is an alternative to websysd as a means of running the DP apps locally.

**[dp-zebedee-content](https://github.com/ONSdigital/dp-zebedee-content)**

Zebedee requires a very specific file structure in order to run along with additional environment variables. You can use [dp-zebedee-content](https://github.com/ONSdigital/dp-zebedee-content) to simplify the setup process.

Licence
-------

Copyright ©‎ 2016, Office for National Statistics (https://www.ons.gov.uk)

Released under MIT license, see [LICENSE](LICENSE.md) for details.
