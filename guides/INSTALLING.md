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
* [Docker](https://www.docker.com/get-started) and **Docker Compose**

  - `$ brew cask install docker`
  - `$ brew install docker-compose`

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

Clone the GitHub repos for [web](#web-journey), [publishing](#publishing-journey) and/or [CMD](#cmd-journeys).

- [Web](#web-journey) - These apps make up the public facing website providing read only access to published content, and will be enough strictly to work on website content types other than filterable datasets (e.g. bulletins, articles, timeseries, datasets). 

- [Publishing](#publishing-journey) - The "publishing journey" gives you all the features of web together with an internal interface to update, preview and publish content. All content is encrypted and requires authentication.

- [CMD](#cmd-journeys) - apps will support the filterable dataset journey, and would mean you have every possible service running.

### Web Journey

* [babbage](https://github.com/ONSdigital/babbage)
  - `$ git clone git@github.com:ONSdigital/babbage`
* [zebedee](https://github.com/ONSdigital/zebedee)
  - `$ git clone git@github.com:ONSdigital/zebedee`
* [sixteens](https://github.com/ONSdigital/sixteens)
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

### Publishing Journey

All services listed in [web](#web-journey) are required for the publishing journey. They are used for the preview functionality.

* [florence](https://github.com/ONSdigital/florence)
  - `$ git clone git@github.com:ONSdigital/florence`

  Then work through the README.md in the florence directory.
* [The-Train](https://github.com/ONSdigital/The-Train)
  - `$ git clone git@github.com:ONSdigital/The-Train`

NOTE: The below services were cloned in [web](#web-journey), but have a separate additional instance when run in publishing:

* [babbage](https://github.com/ONSdigital/babbage)
* [zebedee](https://github.com/ONSdigital/zebedee)

### CMD Journeys

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
* [dp-dimension-extractor](https://github.com/ONSdigital/dp-dimension-extractor)
  - `$ git clone git@github.com:ONSdigital/dp-dimension-extractor`
* [dp-dimension-importer](https://github.com/ONSdigital/dp-dimension-importer)
  - `$ git clone git@github.com:ONSdigital/dp-dimension-importer`
* [dp-observation-extractor](https://github.com/ONSdigital/dp-observation-extractor)
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
`SERVICE_AUTH_TOKEN` | a value for `zebedee` to work
`NEW_HOMEPAGE_ENABLED` | set `true` to get new ONS homepage

After all the various steps, here's an example set of exports and their values that you might now have in your [startup file](#startup-file):

**# ONS services**

**export zebedee_root=~/Documents/website/zebedee-content/generated**

**export TRANSACTION_STORE=$zebedee_root/zebedee/transactions**

**export WEBSITE=$zebedee_root/zebedee/master**

**export PUBLISHING_THREAD_POOL_SIZE=10**

**export ENABLE_PRIVATE_ENDPOINTS=true**

**export ENABLE_PERMISSIONS_AUTH=true**

**export ENCRYPTION_DISABLED=true**

**export DATASET_ROUTES_ENABLED=true**

**export FORMAT_LOGGING=true**

**export SERVICE_AUTH_TOKEN="fc4089e2e12937861377629b0cd96cf79298a4c5d329a2ebb96664c88df77b67"**

**export NEW_HOMEPAGE_ENABLED=true**

Return to the [Getting Started](https://github.com/ONSdigital/dp/blob/master/guides/GETTING_STARTED.md) guide for next steps.

--------------

## Running the apps

Before running web or publishing, please make sure to run [dp-compose](https://github.com/ONSdigital/dp-compose) using the `./run.sh` command (in the dp-compose repo) to run the supporting services.

Most applications can be run using the `$ make debug` command, but deviations are all documented below:

#### Web

  - The website will be available at http://localhost:20000

    after you run up all of the following:
* [babbage](https://github.com/ONSdigital/babbage) - use: `$ ./run.sh`
* [zebedee](https://github.com/ONSdigital/zebedee) - use: `./run-reader.sh`
* [sixteens](https://github.com/ONSdigital/sixteens) - use: `./run.sh`
* [dp-frontend-router](https://github.com/ONSdigital/dp-frontend-router)
* [dp-frontend-renderer](https://github.com/ONSdigital/dp-frontend-renderer)
* [dp-frontend-homepage-controller](https://github.com/ONSdigital/dp-frontend-homepage-controller)
* [dp-frontend-cookie-controller](https://github.com/ONSdigital/dp-frontend-cookie-controller)
* [dp-frontend-dataset-controller](https://github.com/ONSdigital/dp-frontend-dataset-controller)

#### Publishing

Run all of the services in Web, BUT change the commands used to run babbage and zebedee to:

* [babbage](https://github.com/ONSdigital/babbage) - use: `$ ./run-publishing.sh`
* [zebedee](https://github.com/ONSdigital/zebedee) - use: `$ ./run.sh`

and also run the following:

* [florence](https://github.com/ONSdigital/florence) - use: `$ make debug ENCRYPTION_DISABLED=true`
* [The-Train](https://github.com/ONSdigital/The-Train) - use: `$ ./run.sh`
* [dp-api-router](https://github.com/ONSdigital/dp-api-router)

Florence will be available at http://localhost:8081/florence/login

The website will be available at http://localhost:8081 (only available after a successful login into Florence - see the README.md in the florence directory)

  If you need to edit content in Florence, [try this guide](https://github.com/ONSdigital/florence/blob/develop/USAGE.md)

#### CMD

Services cloned in web/publishing that must be run with alternate commands:
* [dp-frontend-router](https://github.com/ONSdigital/dp-frontend-router) - use: `$ make debug DATASET_ROUTES_ENABLED=true`
* [florence](https://github.com/ONSdigital/florence) - use: `$ make debug ENABLE_DATASET_IMPORT=true ENCRYPTION_DISABLED=true`

Other commands to be run with alternative commands:
* [dp-dimension-extractor](https://github.com/ONSdigital/dp-dimension-extractor) - use: `$ make debug ENCRYPTION_DISABLED=true`
* [dp-observation-extractor](https://github.com/ONSdigital/dp-observation-extractor) - use `$ make debug ENCRYPTION_DISABLED=true`

All other services listed in [web](#web-journey) AND [publishing](#publishing-journey) should also be run with this CMD stack to get the full journey - to import data, publish it and then test the public journey.

If you already have content, and you just want to run the web journey, you'll need the [dataset](#dataset-journey), [filter](#filter-journey) and [web](#web-journey) services.

The publishing journey for CMD requires all of the services listed, including the [import](#import-services) and [publishing](#publishing-journey) services. To configure [dataset](#dataset-journey) and [filter](#filter-journey) services for use in publishing, export `ENABLE_PRIVATE_ENDPOINTS=true` for every service.

Return to the [Getting Started](https://github.com/ONSdigital/dp/blob/master/guides/GETTING_STARTED.md) guide for next steps.

--------------

## Setup credentials

To run florence, you will need to update the environment variable `SERVICE_AUTH_TOKEN` in your [startup file](#startup-file) by following:

* Service authentication token creation steps can be found in the [Zebedee repository](https://github.com/ONSdigital/zebedee/#service-authentication-with-zebedee)

You will need to restart your terminal for the environment variable change to take effect.

Note that when the first login to a Florence account is detected a mandatory password update is required.
