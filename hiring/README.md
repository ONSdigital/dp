Digital Publishing Overview
===========================

In Digital Publishing we are responsible for the ONS website, its corresponding APIs and content management system that allows our Publishing Team to prepare and publish all content for the organisation. To give an idea of what this looks like in practice, we've put together this list of relevant links to our products, documentation and code bases. 

The only code we keep private relates to our infrastructure, which is also why there's no mention of platform services in this list. To learn more about our underlying platform see this [list of technologies](../TECHNOLOGIES.md)

### Product Links

* [ONS website](https://www.ons.gov.uk)
* Our [original API](https://api.ons.gov.uk) and [developer site](https://developer.ons.gov.uk)
* Our [new API](https://api.beta.ons.gov.uk) and [developer site](https://developer.beta.ons.gov.uk)

### Useful documents
* [DP](https://github.com/ONSdigital/dp) contains a getting started guide, and some of our hiring documents. Any documentation that doesn't live with code is likely in this repository.
* [Our development standards](https://github.com/ONSdigital/dp-standards) contains documentation on development standards and specifications
* [Our principles](https://github.com/ONSdigital/dp-principles) are core to the way we work, and apply across all of Digital Publishing, not just the engineering team.

### Key Repositories
* [Florence](https://github.com/ONSdigital/florence) is our internal CMS UI. It is written in React/JS with a small Go server side component
* [Zebedee](https://github.com/ONSdigital/zebedee) is the CMS backend, written in Java and using a custom web framework.
* [Babbage](https://github.com/ONSdigital/babbage) runs the bulk of the ONS website, it is also written in Java using the same custom framework as Zebedee. Some of the original functionality is being broken out into Go microservices such as:
    * [Router](https://github.com/ONSdigital/dp-frontend-router)
    * [Renderer](https://github.com/ONSdigital/dp-frontend-renderer)
    * [Dataset Controller](https://github.com/ONSdigital/dp-frontend-dataset-controller)
* [Dataset API](https://github.com/ONSdigital/dp-dataset-api) provides the core of our new API focussed journey, written in Go and backed by our graph database.
* [Dataset exporter](https://github.com/ONSdigital/dp-dataset-exporter) is one of our event-driven Go services, which generates custom CSV files determined by user queries.
