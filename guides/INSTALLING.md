Install, configure and run services
===============
As listed step-by-step in the [Getting Started](https://github.com/ONSdigital/dp/blob/master/guides/GETTING_STARTED.md) guide. **You must follow the steps in the Getting Started guide** to ensure steps not documented here are not missed.

--------------

## Prerequisites

In the below, the installation of each app is typically one of:
- use the `brew` command where provided, or
- use the link to the website to follow the installation instructions, or
- follow the link to the Github repo, where you should clone the repo and follow the instructions in the `README.md` file to install/run (within the repo directory)

**Note:** when indicating a command that should be run in your terminal, we use the `$` prefix to indicate your shell prompt.

---
* [Java 8 JDK (OpenJDK)](https://openjdk.java.net/install/) - install via
  - `$ brew cask install adoptopenjdk8`
* [Maven](https://maven.apache.org/)
  - `$ brew install maven`
* [Docker](https://www.docker.com/get-started)

  - on the website, download **Docker Desktop**, install and RUN the Docker application to complete installation.

* [Cypher Shell](https://neo4j.com/docs/operations-manual/current/tools/cypher-shell/)
  - `$ brew install cypher-shell`
* [Node.js and npm](https://nodejs.org/en/), install the LTS version with:

  - `$ brew install node@12`
  
  - append `export PATH="/usr/local/opt/node@12/bin:$PATH"` to the your `.zshrc` or `.basrc` file and restart your terminal.

* [dp-compose](https://github.com/ONSdigital/dp-compose)

  `$ git clone git@github.com:ONSdigital/dp-compose`

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
* [GoConvey](https://github.com/smartystreets/goconvey#installation)
* [GhostScript](https://www.ghostscript.com/download.html) - Required for [Babbage](https://github.com/onsdigital/babbage)

  - `$ brew install ghostscript`
* [Vault](https://www.vaultproject.io/intro/getting-started/install.html) - Required for running Florence.

  - `$ brew install hashicorp/tap/vault`

* [jq](https://stedolan.github.io/jq/) - a handy JSON tool (for debugging website content)
  - `$ brew install jq`

* [yq](https://github.com/mikefarah/yq) - a handy YAML tool
  - `$ brew install yq`

Return to the [Getting Started](https://github.com/ONSdigital/dp/blob/master/guides/GETTING_STARTED.md) guide for next steps.

---

## Clone the services

Clone the GitHub repos for [web](#web), [publishing](#publishing) and/or [CMD](#cmd).

- [Web](#web) will be enough strictly to work on website content types other than filterable datasets (e.g. bulletins, articles, timeseries, datasets). 

- [Publishing](#publishing) gives you the ability to update, preview and publish content - required for changes to any part of the publishing system. 

- [CMD](#cmd) apps will support the filterable dataset journey, and would mean you have ever possible service running.

Unless otherwise stated, within a services directory it is run using `$ make debug`

### Web

* [babbage](https://github.com/ONSdigital/babbage) - use: `$ ./run.sh`

  - `$ git clone git@github.com:ONSdigital/babbage`

* [zebedee](https://github.com/ONSdigital/zebedee) - use: `./run-reader.sh`
  - `$ git clone git@github.com:ONSdigital/zebedee`

* [sixteens](https://github.com/ONSdigital/sixteens) - use: `./run.sh`
  - `$ git clone git@github.com:ONSdigital/sixteens`

* [dp-frontend-router](https://github.com/ONSdigital/dp-frontend-router)
  - `$ git clone git@github.com:ONSdigital/dp-frontend-router`

* [dp-frontend-renderer](https://github.com/ONSdigital/dp-frontend-renderer)
  - `$ git clone git@github.com:ONSdigital/dp-frontend-renderer`

* [dp-frontend-homepage-controller](https://github.com/ONSdigital/dp-frontend-homepage-controller)
  - `$ git clone git@github.com:ONSdigital/dp-frontend-homepage-controller`

* [dp-frontend-cookie-controller](https://github.com/ONSdigital/dp-frontend-cookie-controller)
  - `$ git clone git@github.com:ONSdigital/dp-frontend-cookie-controller`

* [dp-frontend-dataset-controller](https://github.com/ONSdigital/dp-frontend-dataset-controller)
  - `$ git clone git@github.com:ONSdigital/dp-frontend-dataset-controller`

### Publishing

Services cloned in [web](#web) that must be run with alternate commands:
* [babbage](https://github.com/ONSdigital/babbage) - use `$ ./run-publishing.sh`
* [zebedee](https://github.com/ONSdigital/zebedee) - use `$ ./run.sh`

New services for publishing:
* [florence](https://github.com/ONSdigital/florence) - use: `$ make debug ENCRYPTION_DISABLED=true`
  - `$ git clone git@github.com:ONSdigital/florence`

* [The-Train](https://github.com/ONSdigital/The-Train) - use: `$ ./run.sh`
  - `$ git clone git@github.com:ONSdigital/The-Train`

All other services listed in [web](#web) are also required for the publishing stack, as they are used for the preview functionality.


### CMD

Services cloned in web/publishing that must be run with alternate commands:
* [dp-frontend-router](https://github.com/ONSdigital/dp-frontend-router) - use `$ make debug DATASET_ROUTES_ENABLED=true`
* [florence](https://github.com/ONSdigital/florence) - use `$ make debug ENABLE_DATASET_IMPORT=true ENCRYPTION_DISABLED=true`

All other services listed in [web](#web) AND [publishing](#publishing) should also be run with this CMD stack to get the full journey - to import data, publish it and then test the public journey.

If you already have content, and you just want to run the web journey, you'll need the [dataset](#dataset-journey), [filter](#filter-journey) and [web](#web) services.

The publishing journey for CMD requires all of the services listed, including the [import](#import-services) and [publishing](#publishing) services. To configure [dataset](#dataset-journey) and [filter](#filter-journey) services for use in publishing, export `ENABLE_PRIVATE_ENDPOINTS=true` for every service.

#### Dataset journey:
* [dp-api-router](https://github.com/ONSdigital/dp-api-router)
  - `$ git clone git@github.com:ONSdigital/dp-api-router`
* [dp-dataset-api](https://github.com/ONSdigital/dp-dataset-api)
  - `$ git clone git@github.com:ONSdigital/dp-dataset-api`
* [dp-frontend-filter-dataset-controller](https://github.com/ONSdigital/dp-frontend-filter-dataset-controller)
  - `$ git clone git@github.com:ONSdigital/dp-frontend-filter-dataset-controller`
* [dp-frontend-geography-controller](https://github.com/ONSdigital/dp-frontend-geography-controller)
  - `$ git clone git@github.com:ONSdigital/dp-frontend-geography-controller`

#### Import services:
* [dp-recipe-api](https://github.com/ONSdigital/dp-recipe-api)
  - `$ git clone git@github.com:ONSdigital/dp-recipe-api`
* [dp-import-api](https://github.com/ONSdigital/dp-import-api)
  - `$ git clone git@github.com:ONSdigital/dp-import-api`
* [dp-import-tracker](https://github.com/ONSdigital/dp-import-tracker)
  - `$ git clone git@github.com:ONSdigital/dp-import-tracker`
* [dp-dimension-extractor](https://github.com/ONSdigital/dp-dimension-extractor) - use: `$ make debug ENCRYPTION_DISABLED=true`
  - `$ git clone git@github.com:ONSdigital/dp-dimension-extractor`

* [dp-dimension-importer](https://github.com/ONSdigital/dp-dimension-importer)
  - `$ git clone git@github.com:ONSdigital/dp-dimension-importer`
* [dp-observation-extractor](https://github.com/ONSdigital/dp-observation-extractor) - use: `$ make debug ENCRYPTION_DISABLED=true`
  - `$ git clone git@github.com:ONSdigital/dp-observation-extractor`
* [dp-observation-importer](https://github.com/ONSdigital/dp-observation-importer)
  - `$ git clone git@github.com:ONSdigital/dp-observation-importer`
* [dp-hierarchy-builder](https://github.com/ONSdigital/dp-hierarchy-builder)
  - `$ git clone git@github.com:ONSdigital/dp-hierarchy-builder`
* [dp-search-builder](https://github.com/ONSdigital/dp-search-builder)
  - `$ git clone git@github.com:ONSdigital/dp-search-builder`
* [dp-publishing-dataset-controller](https://github.com/ONSdigital/dp-publishing-dataset-controller)
  - `$ git clone git@github.com:ONSdigital/dp-publishing-dataset-controller`

#### Filter journey:
* [dp-search-api](https://github.com/ONSdigital/dp-search-api)
  - `$ git clone git@github.com:ONSdigital/dp-search-api`
* [dp-code-list-api](https://github.com/ONSdigital/dp-code-list-api)
  - `$ git clone git@github.com:ONSdigital/dp-code-list-api`
* [dp-hierarchy-api](https://github.com/ONSdigital/dp-hierarchy-api)
  - `$ git clone git@github.com:ONSdigital/dp-hierarchy-api`
* [dp-filter-api](https://github.com/ONSdigital/dp-filter-api)
  - `$ git clone git@github.com:ONSdigital/dp-filter-api`
* [dp-dataset-exporter](https://github.com/ONSdigital/dp-dataset-exporter)
  - `$ git clone git@github.com:ONSdigital/dp-dataset-exporter`
* [dp-dataset-exporter-xlsx](https://github.com/ONSdigital/dp-dataset-exporter-xlsx)
  - `$ git clone git@github.com:ONSdigital/dp-dataset-exporter-xlsx`
* [dp-download-service](https://github.com/ONSdigital/dp-download-service)
  - `$ git clone git@github.com:ONSdigital/dp-download-service`

Return to the [Getting Started](https://github.com/ONSdigital/dp/blob/master/guides/GETTING_STARTED.md) guide for next steps.

--------------

## Configuration

### Startup file

Some commands require changes to be made to your shell - e.g.

- to your `PATH` or
- to add environment variables - these commands take the form `export VAR_NAME=value`

and should be appended to the startup file for your shell:

- for the shell `zsh`, the startup file is `~/.zshrc`
- for the `bash` shell, the startup file is `~/.bashrc`

When the startup files are updated, to load the new changes into your shell, either:

- open a new terminal window, or
- `$ exec $SHELL -l`

### Environment variables

You should put the below _env vars_ in your [startup file](#startup-file).

VAR_NAME | note
--- | ---
`zebedee_root` | path to your zebedee content, typically the directory the [dp-zebedee-content](https://github.com/ONSdigital/dp-zebedee-content) generation script points to when run
`ENABLE_PRIVATE_ENDPOINTS` | set `true` when running services in publishing, unset for web mode
`ENABLE_PERMISSIONS_AUTH`| set `true` to ensure that calls to APIs are from registered services or users
`ENCRYPTION_DISABLED` | set `true` to disable encryption, making data readable for any debugging purposes
`DATASET_ROUTES_ENABLED` | `true` will enable the filterable dataset routes (the CMD journey) in some services
`FORMAT_LOGGING` | if `true` then `zebedee` will format its logs

Return to the [Getting Started](https://github.com/ONSdigital/dp/blob/master/guides/GETTING_STARTED.md) guide for next steps.

--------------

## Running the apps

Before running web or publishing, please make sure to run [dp-compose](https://github.com/ONSdigital/dp-compose) using the `./run.sh` command (in the dp-compose repo) to run the supporting services

Most applications can be run using the `make debug` command, but deviations are all documented alongside each repo in the [cloning](#clone-the-services) guide.

#### Publishing
  - Florence will be available at `http://localhost:8081/florence/login`
  - The website will be available at `http://localhost:8081` (only available after a successful login into Florence)

  If you need to edit content in Florence, [try this guide](https://github.com/ONSdigital/florence/blob/develop/USAGE.md)

#### Web
  - The website will be available at `http://localhost:22000`

Return to the [Getting Started](https://github.com/ONSdigital/dp/blob/master/guides/GETTING_STARTED.md) guide for next steps.

--------------

## Setup credentials
* In Zebedee `run.sh` remove the following line: export `SERVICE_AUTH_TOKEN="fc4089e2e12937861377629b0cd96cf79298a4c5d329a2ebb96664c88df77b67"`
* Service authentication token creation steps can be found in the [Zebedee repository](https://github.com/ONSdigital/zebedee/#service-authentication-with-zebedee)

Note that when the first login to a Florence account is detected a mandatory password update is required.
