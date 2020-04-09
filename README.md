Digital Publishing Platform
===========================

The digital publishing platform is a collection of services which provide
content management and publishing functionality for ONS content, and is
responsible for the public ONS website.

Guides and general info
-----------------------

* [Getting started](guides/GETTING_STARTED.md)
* [Creating a new app](guides/NEW_APP.md)
* [Release process](guides/RELEASES.md)
* [Contributing guidelines](guides/CONTRIBUTING.md)
* [GPG commit signing](guides/GPG.md)
* [Port numbering](guides/PORTS.md)
* [Technologies](guides/TECHNOLOGIES.md)
* [Tagging a release](guides/TAGS.md)
* [Migrating Go vendor projects to Go modules](guides/MODULES.md)
* [AWS credentials](guides/AWS_CREDENTIALS.md)

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

**[dp command line tool](https://github.com/ONSdigital/dp-cli)**

A command line utility to assist work in Digital Publishing. See its own repository
([dp-cli](https://github.com/ONSdigital/dp-cli))

**[websysd](websysd/README.md)**

You can use websysd to run your DP apps locally. However, currently the configuration is invalid as it assumes that they
are located in the GOPATH which is not the case for the majority of apps that have been migrated to go modules.

**[paasbox](paasbox)**

Paasbox is an alternative to websysd as a means of running the DP apps locally. Similar to websysd, its configuration
will need updating to reflect the new locations of apps migrated to go modules.

**[dp-zebedee-content](https://github.com/ONSdigital/dp-zebedee-content)**

Zebedee requires a very specific file structure in order to run along with additional environment variables. You can use [dp-zebedee-content](https://github.com/ONSdigital/dp-zebedee-content) to simplify the setup process.

Licence
-------

Copyright ©‎ 2016, Office for National Statistics (https://www.ons.gov.uk)

Released under MIT license, see [LICENSE](LICENSE.md) for details.
