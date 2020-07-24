Install, configure and run services
===============
As listed step-by-step in the [Getting Started](https://github.com/ONSdigital/dp/blob/master/guides/GETTING_STARTED.md) guide. **You must follow the steps in the Getting Started guide** to ensure steps not documented here are not missed.

## Prerequisites

* [Java 8 JDK (OpenJDK)](https://openjdk.java.net/install/)
* [Maven](https://maven.apache.org/)
* [Docker](https://www.docker.com/get-started)
* [Cypher Shell](https://neo4j.com/docs/operations-manual/current/tools/cypher-shell/) - installed with `brew install cypher-shell`
* [Node.js and npm](https://nodejs.org/en/), known working versions:
  - Node.js version 10.15.3
  - npm version 6.9.0
* [dp-compose](https://github.com/ONSdigital/dp-compose)

  NB. See `dp-compose` README for configuration of Docker Desktop resources

  - Services for Web
    - Elasticsearch 2.4.2
    - Highcharts
    - Postgres
    - MathJax 
  - Services for CMD
    - MongoDB
    - Elasticsearch 5 (on non-standard port)
    - Kafka (plus required Zookeeper dependency)
    - Neo4J
* [go v1.13](https://golang.org/doc/install)
* [GoConvey](https://github.com/smartystreets/goconvey#installation)
* [Govendor](https://github.com/kardianos/govendor)
* [GhostScript](https://www.ghostscript.com/download.html) - Required for [Babbage](https://github.com/onsdigital/babbage)
* [Vault](https://www.vaultproject.io/intro/getting-started/install.html) - This will be required for running Florence.

* [jq](https://stedolan.github.io/jq/) - installed with `brew install jq`
  - This isn't essential, but it's a useful tool for debugging website content.

## Clone the services

Clone the GitHub repos for [web](#web), [publishing](#publishing) and/or [CMD](#cmd). Unless otherwise stated, services are run using `$ make debug`

[Web](#web) will be enough strictly to work on website content types other than filterable datasets (e.g. bulletins, articles, timeseries, datasets). 

[Publishing](#publishing) gives you the ability to update, preview and publish content - required for changes to any part of the publishing system. 

[CMD](#cmd) apps will support the filterable dataset journey, and would mean you have ever possible service running.

### Web

* [babbage](https://github.com/ONSdigital/babbage) - use `$ ./run.sh`
* [zebedee](https://github.com/ONSdigital/zebedee) - use `$ ./run-reader.sh`
* [sixteens](https://github.com/ONSdigital/sixteens) - use `$ ./run.sh`
* [dp-frontend-router](https://github.com/ONSdigital/dp-frontend-router)
* [dp-frontend-renderer](https://github.com/ONSdigital/dp-frontend-renderer)
* [dp-frontend-homepage-controller](https://github.com/ONSdigital/dp-frontend-homepage-controller)
* [dp-frontend-cookie-controller](https://github.com/ONSdigital/dp-frontend-cookie-controller)
* [dp-frontend-dataset-controller](https://github.com/ONSdigital/dp-frontend-dataset-controller)

### Publishing

Services cloned in [web](#web) that must be run with alternate commands:
* [babbage](https://github.com/ONSdigital/babbage) - use `$ ./run-publishing.sh`
* [zebedee](https://github.com/ONSdigital/zebedee) - use `$ ./run.sh`

New services for publishing:
* [florence](https://github.com/ONSdigital/florence) - use `$ make debug ENCRYPTION_DISABLED=true`
* [The-Train](https://github.com/ONSdigital/The-Train) - use `$ ./run.sh`

All other services listed in [web](#web) are also required for the publishing stack, as they are used for the preview functionality.


### CMD

Services cloned in web/publishing that must be run with alternate commands:
* [dp-frontend-router](https://github.com/ONSdigital/dp-frontend-router) - use `$ make debug DATASET_ROUTES_ENABLED=true`
* [florence](https://github.com/ONSdigital/florence) - use `$ make debug ENABLE_DATASET_IMPORT=true ENCRYPTION_DISABLED=true`

All other services listed in [web](#web) AND [publishing](#publishing) should also be run with the CMD stack to get the full journey - to import data, publish it and then test the public journey.

If you already have content, and you just want to run the web journey, you'll need the [dataset](#dataset-journey), [filter](#filter-journey) and [web](#web) services.

The publishing journey for CMD requires all of the services listed, including the [import](#import-services) and [publishing](#publishing) services. To configure [dataset](#dataset-journey) and [filter](#filter-journey) services for use in publishing, export `ENABLE_PRIVATE_ENDPOINTS=true` for every service.

#### Dataset journey:
* [dp-api-router](https://github.com/ONSdigital/dp-api-router)
* [dp-dataset-api](https://github.com/ONSdigital/dp-dataset-api)
* [dp-frontend-filter-dataset-controller](https://github.com/ONSdigital/dp-frontend-filter-dataset-controller)
* [dp-frontend-geography-controller](https://github.com/ONSdigital/dp-frontend-geography-controller)

#### Import services:
* [dp-recipe-api](https://github.com/ONSdigital/dp-recipe-api)
* [dp-import-api](https://github.com/ONSdigital/dp-import-api)
* [dp-import-tracker](https://github.com/ONSdigital/dp-import-tracker)
* [dp-dimension-extractor](https://github.com/ONSdigital/dp-dimension-extractor) - use `$ make debug ENCRYPTION_DISABLED=true`
* [dp-dimension-importer](https://github.com/ONSdigital/dp-dimension-importer)
* [dp-observation-extractor](https://github.com/ONSdigital/dp-observation-extractor) - use `$ make debug ENCRYPTION_DISABLED=true`
* [dp-observation-importer](https://github.com/ONSdigital/dp-observation-importer)
* [dp-hierarchy-builder](https://github.com/ONSdigital/dp-hierarchy-builder)
* [dp-search-builder](https://github.com/ONSdigital/dp-search-builder)
* [dp-publishing-dataset-controller](https://github.com/ONSdigital/dp-publishing-dataset-controller)

#### Filter journey:
* [dp-search-api](https://github.com/ONSdigital/dp-search-api)
* [dp-code-list-api](https://github.com/ONSdigital/dp-code-list-api)
* [dp-hierarchy-api](https://github.com/ONSdigital/dp-hierarchy-api)
* [dp-filter-api](https://github.com/ONSdigital/dp-filter-api)
* [dp-dataset-exporter](https://github.com/ONSdigital/dp-dataset-exporter)
* [dp-dataset-exporter-xlsx](https://github.com/ONSdigital/dp-dataset-exporter-xlsx)
* [dp-download-service](https://github.com/ONSdigital/dp-download-service)


## Configuration

The below environment variables need to be set to run the stack locally:

If using Zsh then enter the variables in your `~/.zshenv:

```zsh
{var_name}={value}
```

If using Bash then enter the variables in your `~/.bash_profile`:

```bash
export {var_name}={value}
```

Where `{var_name}` is to be replaced by a variable name like the ones below, and `{value}` with the value of the variable - if it is a string value then surround with quotes (`"`).

- `zebedee_root` should be a path to your zebedee content typically the directory the [dp-zebedee-content](https://github.com/ONSdigital/dp-zebedee-content) generation script points to when run.

- `TRANSACTION_STORE` = `$zebedee_root/zebedee/transactions`.

- `WEBSITE` = `$zebedee_root/zebedee/master`.

- `PUBLISHING_THREAD_POOL_SIZE` should be set to a sensible number like `10` general rule of thumb two threads for each core, minus 1.

- `ENABLE_PRIVATE_ENDPOINTS` should be set to `true` if wanting to run in publishing mode, alternatively either do not set or set the value to
 `false` for web mode.

- `ENABLE_PERMISSIONS_AUTH` should be set to `true` this ensures that all calls to APIs are from registered services or users. Alternatively set this to `false` if wanting to bypass this.

- `ENCRYPTION_DISABLED` should be set to `true` this will disable encryption making data readable for any debugging purposes.

- `DATASET_ROUTES_ENABLED` should be set to `true` this will enable the filterable dataset routes (the CMD journey).

- (Optional) `FORMAT_LOGGING` if set to `true` this will format logs.

## Running the apps

Most applications can be run using the `make debug` command, but deviations are all documented alongside each repo in the [cloning](#clone-the-services) guide.

#### Publishing
  - Florence will be available at `http://localhost:8081/florence/login`
  - The website will be available at `http://localhost:8081` (only available after a successful login into Florence)

  If you need to edit content in Florence, [try this guide](https://github.com/ONSdigital/florence/blob/develop/USAGE.md)

#### Web
  - The website will be available at `http://localhost:22000`

## Setup credentials
* In Zebedee `run.sh` remove the following line: export `SERVICE_AUTH_TOKEN="fc4089e2e12937861377629b0cd96cf79298a4c5d329a2ebb96664c88df77b67"`
* Service authentication token creation steps can be found in the [Zebedee repository](https://github.com/ONSdigital/zebedee/#service-authentication-with-zebedee)

Note that when the first login to a Florence account is detected a mandatory password update is required.
