Digital Publishing Platform
===========================

The digital publishing platform is a collection of services which provide
content management and publishing functionality for ONS content, and is
responsible for the public ONS website.

### Guides

* [Getting started](GETTING_STARTED.md)
* [Development standards](DEV_STANDARDS.md)
* [Definition of Ready](DEFINITION_OF_READY.md)
* [Definition of Done](DEFINITION_OF_DONE.md)
* [Release process](RELEASES.md)
* [Digital Publishing principles](https://github.com/ONSdigital/dp-principles)
* [Contributing guidelines](CONTRIBUTING.md)
* [GPG commit signing](GPG.md)
* [Port numbering](PORTS.md)
* [Technologies](TECHNOLOGIES.md)
* ONS CDN
  * [ONS assets](https://github.com/ONSdigital/cdn.ons.gov.uk-assets)
  * [Vendored dependencies](https://github.com/ONSdigital/cdn.ons.gov.uk-vendor)

### Websysd

You can use [websysd](https://github.com/ONSdigital/dp/blob/master/websysd) to run the dp services from a centralised location:

If you are running the publishing stack, then you must set "ENABLE_PRIVATE_ENDPOINTS=true" and "ENCRYPTION_DISABLED=true” if encryption is not setup.

If you are running the publishing stack in terminals [dp-dimension-extractor](https://github.com/ONSdigital/dp-dimension-extractor), [dp-observation-extractor](https://github.com/ONSdigital/dp-observation-extractor) and [florence](https://github.com/ONSdigital/florence) should have "ENCRYPTION_DISABLED=true” if encryption is not setup. For [dp-dataset-api](https://github.com/ONSdigital/dp-dataset-api), [dp-filter-api](https://github.com/ONSdigital/dp-filter-api) and [dp-search-api](https://github.com/ONSdigital/dp-search-api) you require having "ENABLE_PRIVATE_ENDPOINTS=true".

### Licence

Copyright ©‎ 2016, Office for National Statistics (https://www.ons.gov.uk)

Released under MIT license, see [LICENSE](LICENSE.md) for details.
