# Install, configure and run services

As listed step-by-step in the [Getting Started](https://github.com/ONSdigital/dp/blob/master/guides/GETTING_STARTED.md) guide. **You must follow the steps in the Getting Started guide** to ensure steps documented there are not missed.

---

## Prerequisites

For the below applications, the installation of each app is typically one of:

- use the given `brew` command where provided, or
- use the link to the website to follow the installation instructions, or
- follow the link to the Github repo, where you should clone the repo and follow the instructions in the `README.md` file to install/run (within the repo directory)

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

**Note:** when indicating a command that should be run in your terminal, we use the `$` prefix to indicate your shell prompt.

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

---

### Applications

* [Java 8 JDK (OpenJDK)](https://openjdk.java.net/install/) - install via
  - `$ brew cask install adoptopenjdk8`
* [Maven](https://maven.apache.org/)
  - `$ brew install maven`
* [Docker](https://www.docker.com/get-started)
  - on the website, download **Docker Desktop**, install and *run* the Docker application to complete installation.
* [Cypher Shell](https://neo4j.com/docs/operations-manual/current/tools/cypher-shell/)
  - `$ brew install cypher-shell`
* [Node.js and npm](https://nodejs.org/en/), install the LTS version with:
  - `$ brew install node@12`
  - append the following to your [startup file](#startup-file)
    - `export PATH="/usr/local/opt/node@12/bin:$PATH"`

* [dp-compose](https://github.com/ONSdigital/dp-compose)

  Clone the repo:

  `$ git clone git@github.com:ONSdigital/dp-compose`

  then see the `dp-compose` README file for configuration of Docker Desktop resources

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
  - `$ brew install go@1.13`
* [GoConvey](https://github.com/smartystreets/goconvey#installation)
* [GhostScript](https://www.ghostscript.com/download.html) - Required for [Babbage](https://github.com/onsdigital/babbage)
  - `brew install ghostscript`
* [Vault](https://learn.hashicorp.com/tutorials/vault/getting-started-install)
  - `$ brew install hashicorp/tap/vault`
* [jq](https://stedolan.github.io/jq/) - a handy JSON tool
  - `$ brew install jq`
* [yq](https://github.com/mikefarah/yq) - a handy YAML tool
  - `$ brew install yq`

Return to the [Getting Started](https://github.com/ONSdigital/dp/blob/master/guides/GETTING_STARTED.md) guide for next steps.

---

## Clone and run each service

Clone the GitHub repos for [web](#web), [publishing](#publishing) and/or [CMD](#cmd).

- [Web](#web) will be enough to work on website content other than filterable datasets (e.g. bulletins, articles, timeseries, datasets).

- [Publishing](#publishing) gives you the ability to update, preview and publish content - required for changes to any part of the publishing system.

- [CMD](#cmd) services support the import, publish and filterable dataset journeys.

Once you have cloned the services, you should run them.

The first services that should be started are those of [dp-compose](https://github.com/ONSdigital/dp-compose) using `$ ./run.sh` (in your `dp-compose` directory) to start the supporting services.

Unless otherwise stated, other services are started by changing into the repo directory and running:

`$ make debug`

---

Unless otherwise stated, services are run using `$ make debug`

### Web

* [babbage](https://github.com/ONSdigital/babbage) - use

  - `$ git clone git@github.com:ONSdigital/babbage`
  - `$ cd babbage`
  - `$ ./run.sh`


* [zebedee](https://github.com/ONSdigital/zebedee) - use `$ ./run-reader.sh`
  - `git clone git@github.com:ONSdigital/zebedee`

* [sixteens](https://github.com/ONSdigital/sixteens) - use `$ ./run.sh`
  - `git clone git@github.com:ONSdigital/sixteens`

* [dp-frontend-router](https://github.com/ONSdigital/dp-frontend-router)
  - `git clone git@github.com:ONSdigital/dp-frontend-router`

* [dp-frontend-renderer](https://github.com/ONSdigital/dp-frontend-renderer)
  - `git clone git@github.com:ONSdigital/dp-frontend-renderer`

* [dp-frontend-homepage-controller](https://github.com/ONSdigital/dp-frontend-homepage-controller)
  - `git clone git@github.com:ONSdigital/dp-frontend-homepage-controller`

* [dp-frontend-cookie-controller](https://github.com/ONSdigital/dp-frontend-cookie-controller)
  - `git clone git@github.com:ONSdigital/dp-frontend-cookie-controller`

* [dp-frontend-dataset-controller](https://github.com/ONSdigital/dp-frontend-dataset-controller)
  - `git clone git@github.com:ONSdigital/dp-frontend-dataset-controller`

The website will be available at http://localhost:22000

### Publishing

Some repos are common to both [web](#web) and publishing, so require a different startup command when in publishing:

* [babbage](https://github.com/ONSdigital/babbage)
  - `$ ./run-publishing.sh`
* [zebedee](https://github.com/ONSdigital/zebedee)
  - In Zebedee repo's `run.sh` file, remove the following line:
    - `export SERVICE_AUTH_TOKEN="fc4089e2e12937861377629b0cd96cf79298a4c5d329a2ebb96664c88df77b67"`
  - Service authentication token creation steps can be found in the [Zebedee repository](https://github.com/ONSdigital/zebedee/#service-authentication-with-zebedee)
  - `$ ./run.sh`

The following services run only in _publishing_:

* [florence](https://github.com/ONSdigital/florence)
  - `$ git clone git@github.com:ONSdigital/florence`
  - `$ make debug ENCRYPTION_DISABLED=true`
  - Florence will be available at http://localhost:8081/florence/login
  - The website will be available at http://localhost:8081 (only available after a successful login into Florence)
  - Note that when the first login to Florence, there will be a mandatory password change
* [The-Train](https://github.com/ONSdigital/The-Train)
  - `git clone git@github.com:ONSdigital/The-Train`
  - `$ ./run.sh`

All services listed in [web](#web) are also required for the publishing stack, as they are used for preview functionality.

If you need to edit content in Florence, [try this guide](https://github.com/ONSdigital/florence/blob/develop/USAGE.md)

### CMD

(CMD aka "customise my data")

The end-to-end process (for CMD data) comprises of three steps:

1. _import journey_ - the uploading of a dataset (or CSV file) into the database(s)
1. the _publishing journey_ - adding that dataset to a collection and publishing it onto the website
1. and, finally, the _web or filter journey_ - wherein the public filters and downloads a customised version of the published dataset

The full CMD stack (i.e. the services required for the CMD end-to-end process) comprises:

- all [web](#web) services
- all [publishing](#publishing) services
- and all the services listed below

Sometimes, you have already imported the content into the databases and have published it, so
you may want to start a subset for the web journey. If so, you'll need

- [web](#web) services
- the [dataset](#dataset-journey)
- [filter](#filter-journey)

The publishing journey for CMD requires all of the services listed, including the [import](#import-services) and [publishing](#publishing) services. To configure [dataset](#dataset-journey) and [filter](#filter-journey) services for use in publishing,

  `$ export ENABLE_PRIVATE_ENDPOINTS=true`

for every service.

#### Frontend services

The below services (you should already have cloned these in web and publishing sections, above), must be started with alternate commands for CMD:

- [dp-frontend-router](https://github.com/ONSdigital/dp-frontend-router)
  - `$ make debug DATASET_ROUTES_ENABLED=true`
- [florence](https://github.com/ONSdigital/florence)
  - `$ make debug ENABLE_DATASET_IMPORT=true ENCRYPTION_DISABLED=true`

#### Dataset journey

*  [dp-api-router](https://github.com/ONSdigital/dp-api-router)
* [dp-dataset-api](https://github.com/ONSdigital/dp-dataset-api)
* [dp-frontend-filter-dataset-controller](https://github.com/ONSdigital/dp-frontend-filter-dataset-controller)
* [dp-frontend-geography-controller](https://github.com/ONSdigital/dp-frontend-geography-controller)

#### Import services

* [dp-recipe-api](https://github.com/ONSdigital/dp-recipe-api)
* [dp-import-api](https://github.com/ONSdigital/dp-import-api)
* [dp-import-tracker](https://github.com/ONSdigital/dp-import-tracker)
* [dp-dimension-extractor](https://github.com/ONSdigital/dp-dimension-extractor)
  - `$ make debug ENCRYPTION_DISABLED=true`
* [dp-dimension-importer](https://github.com/ONSdigital/dp-dimension-importer)
* [dp-observation-extractor](https://github.com/ONSdigital/dp-observation-extractor)
  - `$ make debug ENCRYPTION_DISABLED=true`
* [dp-observation-importer](https://github.com/ONSdigital/dp-observation-importer)
* [dp-hierarchy-builder](https://github.com/ONSdigital/dp-hierarchy-builder)
* [dp-search-builder](https://github.com/ONSdigital/dp-search-builder)
* [dp-publishing-dataset-controller](https://github.com/ONSdigital/dp-publishing-dataset-controller)

#### Filter journey

* [dp-search-api](https://github.com/ONSdigital/dp-search-api)
* [dp-code-list-api](https://github.com/ONSdigital/dp-code-list-api)
* [dp-hierarchy-api](https://github.com/ONSdigital/dp-hierarchy-api)
* [dp-filter-api](https://github.com/ONSdigital/dp-filter-api)
* [dp-dataset-exporter](https://github.com/ONSdigital/dp-dataset-exporter)
* [dp-dataset-exporter-xlsx](https://github.com/ONSdigital/dp-dataset-exporter-xlsx)
* [dp-download-service](https://github.com/ONSdigital/dp-download-service)

Return to the [Getting Started](./GETTING_STARTED.md) guide for next steps.
