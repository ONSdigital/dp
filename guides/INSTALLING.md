# Install, configure and run services

As listed step-by-step in the [Getting Started](https://github.com/ONSdigital/dp/blob/main/guides/GETTING_STARTED.md) guide. **You must follow the steps in the Getting Started guide** to ensure steps not documented here are not missed.

--------------

## Prerequisites

In the below, the installation of each app is typically one of:

* use the `brew` command where provided, or
* use the link to the website to follow the installation instructions, or
* follow the link to the Github repo, where you should clone the repo and follow the instructions in the `README.md` file to install/run (within the repo directory)

**Note:** when indicating a command that should be run in your terminal, we use the `$` prefix to indicate your shell prompt.

--------------

Software | Install | Notes
-------- | ------- | ----- |
[Java 8 JDK (OpenJDK)](https://openjdk.java.net/install/) | `$ brew install openjdk@8` | Append `export PATH="/usr/local/opt/openjdk@8/bin:$PATH"` to your shell profile (eg. `.zshrc`) and restart your terminal.
[Maven](https://maven.apache.org/)                        | `$ brew install maven`
[Docker](https://www.docker.com/get-started)              | `$ brew install --cask docker`
 Docker Compose                                           | `$ brew install docker-compose`
[Cypher Shell](https://neo4j.com/docs/operations-manual/current/tools/cypher-shell/)                                | `$ brew install cypher-shell` | deprecated (not needed if using Neptune over Neo4j)
[nvm](https://github.com/nvm-sh/nvm)         		  | Follow the [git install instructions](https://github.com/nvm-sh/nvm?tab=readme-ov-file#git-install) | Required to allow easy switching between node/npm versions depending on usage within app
[Go](https://go.dev/dl/) | `$ brew install go` |  The Go installation is processor architecture specific. For the newer Apple M1 processor the ARM installation is required. This is managed by Homebrew. However, if installing manually this is something to be aware of. [Go direct Download](https://go.dev/dl/) |
[GoConvey](https://github.com/smartystreets/goconvey#installation)  |
[GhostScript](https://www.ghostscript.com/download.html)  |   `$ brew install ghostscript`                          | Required for [Babbage](https://github.com/onsdigital/babbage)
[Vault](https://www.vaultproject.io/intro/getting-started/install.html) | `$ brew install hashicorp/tap/vault`      | Required for running Florence.
[jq](https://stedolan.github.io/jq/)                      | `$ brew install jq`                                     | A handy JSON tool (for debugging website content and much more)
[yq](https://github.com/mikefarah/yq)                     |  `$ brew install yq`                                    | A handy YAML tool
[dp-compose](https://github.com/ONSdigital/dp-compose)    | `$ git clone git@github.com:ONSdigital/dp-compose`      | See [`dp-compose` README](https://github.com/ONSdigital/dp-compose#dp-compose) for configuration of Docker Desktop resources

[dp-compose](https://github.com/ONSdigital/dp-compose) runs the following services:

* Services for the Website
  * Elasticsearch 2.4.2
  * Elasticsearch 7 (on non-standard port)
  * Highcharts
  * Postgres
  * MongoDB
  * Kafka (plus required Zookeeper dependency)
* Services for CMD
  * Elasticsearch 6 (on non-standard port)
  * [Neptune](https://github.com/ONSdigital/dp/blob/main/guides/NEPTUNE.md#migrating-from-neo4j-to-aws-neptune))

Return to the [Getting Started](https://github.com/ONSdigital/dp/blob/main/guides/GETTING_STARTED.md) guide for next steps.

--------------

## Clone the services

Clone the GitHub repos for [web](#web-journey), [publishing](#publishing-journey) and/or [CMD](#cmd-journeys) (Customise My Data).

* [Web](#web-journey) - These apps make up the public-facing website providing **read-only access** to published content, and will be enough strictly to work on website content types other than filterable datasets (e.g. bulletins, articles, timeseries, datasets).

* [Publishing](#publishing-journey) - The "publishing journey" gives you all the features of web together with an internal interface to update, preview and publish content. All content is encrypted and requires authentication.

* [CMD](#cmd-journeys) - apps will support the filterable dataset journey, and would mean you have every possible service running.

### Web Journey

* [babbage](https://github.com/ONSdigital/babbage)
* [zebedee](https://github.com/ONSdigital/zebedee)
* [sixteens](https://github.com/ONSdigital/sixteens)
* [dp-frontend-router](https://github.com/ONSdigital/dp-frontend-router)
* [dp-frontend-homepage-controller](https://github.com/ONSdigital/dp-frontend-homepage-controller)
* [dp-frontend-cookie-controller](https://github.com/ONSdigital/dp-frontend-cookie-controller)
* [dp-frontend-dataset-controller](https://github.com/ONSdigital/dp-frontend-dataset-controller)
* [dp-frontend-feedback-controller](https://github.com/ONSdigital/dp-frontend-feedback-controller)

  ```bash
  git clone git@github.com:ONSdigital/babbage
  git clone git@github.com:ONSdigital/zebedee
  git clone git@github.com:ONSdigital/sixteens

  git clone git@github.com:ONSdigital/dp-frontend-router

  git clone git@github.com:ONSdigital/dp-frontend-homepage-controller
  git clone git@github.com:ONSdigital/dp-frontend-cookie-controller
  git clone git@github.com:ONSdigital/dp-frontend-dataset-controller

  git clone git@github.com:ONSdigital/dp-frontend-feedback-controller
  ```

### Publishing Journey

All services listed in the [web journey](#web-journey) are required for the publishing journey. They are used for the preview functionality.

* [florence](https://github.com/ONSdigital/florence)
* [The-Train](https://github.com/ONSdigital/The-Train)
* [dp-api-router](https://github.com/ONSdigital/dp-api-router)
* [dp-image-api](https://github.com/ONSdigital/dp-image-api)
* [dp-image-importer](https://github.com/ONSdigital/dp-image-importer)
* [dp-upload-service](https://github.com/ONSdigital/dp-upload-service)
* [dp-download-service](https://github.com/ONSdigital/dp-download-service)

  ```bash
  git clone git@github.com:ONSdigital/florence
  git clone git@github.com:ONSdigital/The-Train
  git clone git@github.com:ONSdigital/dp-api-router
  git clone git@github.com:ONSdigital/dp-image-api
  git clone git@github.com:ONSdigital/dp-image-importer
  git clone git@github.com:ONSdigital/dp-upload-service
  git clone git@github.com:ONSdigital/dp-download-service
  ```

### CMD Journeys

All the services in the [web] and [publishing] journeys, as well as:

#### Dataset journey:

* [dp-dataset-api](https://github.com/ONSdigital/dp-dataset-api)
* [dp-frontend-filter-dataset-controller](https://github.com/ONSdigital/dp-frontend-filter-dataset-controller)

  ```bash
  git clone git@github.com:ONSdigital/dp-dataset-api
  git clone git@github.com:ONSdigital/dp-frontend-filter-dataset-controller
  ```

#### Import services:

* [dp-recipe-api](https://github.com/ONSdigital/dp-recipe-api)
* [dp-import-api](https://github.com/ONSdigital/dp-import-api)
* [dp-upload-service](https://github.com/ONSdigital/dp-upload-service)
* [dp-import-tracker](https://github.com/ONSdigital/dp-import-tracker)
* [dp-dimension-extractor](https://github.com/ONSdigital/dp-dimension-extractor)
* [dp-dimension-importer](https://github.com/ONSdigital/dp-dimension-importer)
* [dp-observation-extractor](https://github.com/ONSdigital/dp-observation-extractor)
* [dp-observation-importer](https://github.com/ONSdigital/dp-observation-importer)
* [dp-hierarchy-builder](https://github.com/ONSdigital/dp-hierarchy-builder)
* [dp-hierarchy-api](https://github.com/ONSdigital/dp-hierarchy-api)
* [dp-dimension-search-builder](https://github.com/ONSdigital/dp-dimension-search-builder)
* [dp-publishing-dataset-controller](https://github.com/ONSdigital/dp-publishing-dataset-controller)

  ```bash
  git clone git@github.com:ONSdigital/dp-recipe-api
  git clone git@github.com:ONSdigital/dp-import-api
  git clone git@github.com:ONSdigital/dp-upload-service
  git clone git@github.com:ONSdigital/dp-import-tracker
  git clone git@github.com:ONSdigital/dp-dimension-extractor
  git clone git@github.com:ONSdigital/dp-dimension-importer
  git clone git@github.com:ONSdigital/dp-observation-extractor
  git clone git@github.com:ONSdigital/dp-observation-importer
  git clone git@github.com:ONSdigital/dp-hierarchy-builder
  git clone git@github.com:ONSdigital/dp-hierarchy-api
  git clone git@github.com:ONSdigital/dp-dimension-search-builder
  git clone git@github.com:ONSdigital/dp-publishing-dataset-controller
  ```

[Documentation of the import process](https://github.com/ONSdigital/dp-import#dp-import)

[Sequence diagram of cmd import process](https://github.com/ONSdigital/dp-import/tree/main/docs/import-sequence#readme)

#### Filter journey:
If you have already setup the import journey, you will have the Hierarchy API already. It's still fine to copy the command set below, just be aware that if you hit 1 error for `destination path already exists` that is expected.

* [dp-dimension-search-api](https://github.com/ONSdigital/dp-dimension-search-api)
* [dp-code-list-api](https://github.com/ONSdigital/dp-code-list-api)
* [dp-hierarchy-api](https://github.com/ONSdigital/dp-hierarchy-api)
* [dp-filter-api](https://github.com/ONSdigital/dp-filter-api)
* [dp-dataset-exporter](https://github.com/ONSdigital/dp-dataset-exporter)
* [dp-dataset-exporter-xlsx](https://github.com/ONSdigital/dp-dataset-exporter-xlsx)

  ```bash
  git clone git@github.com:ONSdigital/dp-dimension-search-api
  git clone git@github.com:ONSdigital/dp-code-list-api
  git clone git@github.com:ONSdigital/dp-hierarchy-api
  git clone git@github.com:ONSdigital/dp-filter-api
  git clone git@github.com:ONSdigital/dp-dataset-exporter
  git clone git@github.com:ONSdigital/dp-dataset-exporter-xlsx
  ```

Return to the [Getting Started](https://github.com/ONSdigital/dp/blob/main/guides/GETTING_STARTED.md) guide for next steps.

### Cantabular

[dp-compose](https://github.com/ONSdigital/dp-compose/tree/main) contains a few stacks for Cantabular services, including the [Cantabular import journey](https://github.com/ONSdigital/dp-compose/tree/main/cantabular-import#cantabular-import-journey) and [Cantabular metadata publishing](https://github.com/ONSdigital/dp-compose/tree/main/cantabular-metadata-pub)

Both of these stacks rely on variations of an `scs.sh` script, which provides support in cloning, updating and running all the necessary repos for these journeys.

[See more information and diagrams](https://docs.google.com/document/d/13U5kM3ZwmfNXdy7dq-RX0nZ0hQe5wzgSlDSOvZRP8eU)

--------------

## Configuration

### Startup file

Some commands require changes to be made to your shell - e.g.

* to your `PATH` or
* to add environment variables - these commands take the form `export VAR_NAME=value`

and should be appended to the startup file for your shell:

* for the shell `zsh`, the startup file is `~/.zshrc`
* for the `bash` shell, the startup file is `~/.bashrc`

When the startup files are updated, to load the new changes into your shell, either:

* open a new terminal window, or
* `$ exec $SHELL -l`

### Environment variables

You should put the below _env vars_ in your [startup file](#startup-file).

Variable name | note
--- | ---
`zebedee_root` | path to your zebedee content, typically the directory the [dp-zebedee-content](https://github.com/ONSdigital/dp-zebedee-content) generation script points to when run
`ENABLE_PRIVATE_ENDPOINTS` | set `true` when running services in publishing, unset for web mode
`ENABLE_PERMISSIONS_AUTH`| set `true` to ensure that calls to APIs are from registered services or users
`ENCRYPTION_DISABLED` | set `true` to disable encryption, making data readable for any debugging purposes
`DATASET_ROUTES_ENABLED` | `true` will enable the filterable dataset routes (the CMD journey) in some services
`FORMAT_LOGGING` | if `true` then `zebedee` will format its logs
`SERVICE_AUTH_TOKEN` | a value for `zebedee` to work

After all the various steps, here's an example set of exports and their values that you might now have in your [startup file](#startup-file):

```bash
# Dissemination services
export zebedee_root=~/Documents/website/zebedee-content/generated
export ENABLE_PRIVATE_ENDPOINTS=true
export ENABLE_PERMISSIONS_AUTH=true
export ENCRYPTION_DISABLED=true
export DATASET_ROUTES_ENABLED=true
export FORMAT_LOGGING=true
export SERVICE_AUTH_TOKEN="fc4089e2e12937861377629b0cd96cf79298a4c5d329a2ebb96664c88df77b67"

export TRANSACTION_STORE=$zebedee_root/zebedee/transactions
export WEBSITE=$zebedee_root/zebedee/master
export PUBLISHING_THREAD_POOL_SIZE=10

# For CMD services
export GRAPH_DRIVER_TYPE=neptune
export GRAPH_ADDR=wss://localhost:8182/gremlin
export NEPTUNE_TLS_SKIP_VERIFY=true

```

Return to the [Getting Started](https://github.com/ONSdigital/dp/blob/main/guides/GETTING_STARTED.md) guide for next steps.

--------------

## Running the apps

Run [dp-compose](https://github.com/ONSdigital/dp-compose) using the `$ ./run.sh` command (in the dp-compose repo) to run the supporting services. As well as Vault, e.g. `$ vault server -dev`.

Most applications can be run using the `$ make debug` command, but deviations are all documented below:

### Web

Run all the services in the [web journey](#web-journey)

* [babbage](https://github.com/ONSdigital/babbage) - use: `$ ./run.sh`
* [zebedee](https://github.com/ONSdigital/zebedee) - use: `$ ./run-reader.sh`
* [sixteens](https://github.com/ONSdigital/sixteens) - use: `$ ./run.sh`
* [dp-frontend-router](https://github.com/ONSdigital/dp-frontend-router)
* [dp-frontend-homepage-controller](https://github.com/ONSdigital/dp-frontend-homepage-controller)
* [dp-frontend-cookie-controller](https://github.com/ONSdigital/dp-frontend-cookie-controller)
* [dp-frontend-dataset-controller](https://github.com/ONSdigital/dp-frontend-dataset-controller)
* [dp-frontend-feedback-controller](https://github.com/ONSdigital/dp-frontend-feedback-controller)

The website will be available at http://localhost:20000

### Publishing

Run all of the services in the [web journey](#web-journey), **but** change the commands used to run babbage and zebedee to:

* [babbage](https://github.com/ONSdigital/babbage) - use: `$ ./run-publishing.sh`
* [zebedee](https://github.com/ONSdigital/zebedee) - use: `$ ./run.sh`

and also run the following:

* [florence](https://github.com/ONSdigital/florence) - use: `$ make debug ENCRYPTION_DISABLED=true`
* [The-Train](https://github.com/ONSdigital/The-Train) - use: `$ ./run.sh`
* [dp-api-router](https://github.com/ONSdigital/dp-api-router)

If you also want to run Florence with the ability to edit images on the homepage (for the Featured Content section), you will need to additionally run:

* [dp-image-api](https://github.com/ONSdigital/dp-image-api)
* [dp-image-importer](https://github.com/ONSdigital/dp-image-importer) - use: `make debug ENCRYPTION_DISABLED=true`
* [dp-upload-service](https://github.com/ONSdigital/dp-upload-service) - use `make debug ENCRYPTION_DISABLED=true`
* [dp-download-service](https://github.com/ONSdigital/dp-download-service)  - use: `make debug ENCRYPTION_DISABLED=true`

Florence will be available at [http://localhost:8081/florence/login](http://localhost:8081/florence/login).

The website will be available at [http://localhost:8081](http://localhost:8081) after a successful login into florence. Login details are in the [florence repository](https://github.com/ONSdigital/florence/blob/develop/USAGE.md).

### CMD

All of the services in the [web](#web-journey), [publishing](#publishing-journey) and [CMD](#cmd-journeys) journeys need to be run for the full CMD journey to work. This journey includes importing data, publishing it and testing the public journey.

> You will want to make sure you have access to the Neptune test instance as well, if you want the entire CMD journey to be accessible. Details on how to set this up can be found [here](https://github.com/ONSdigital/dp/blob/main/guides/NEPTUNE.md).

Use the following alternative commands:

* [florence](https://github.com/ONSdigital/florence) - use: `$ make debug ENCRYPTION_DISABLED=true`
* [dp-frontend-router](https://github.com/ONSdigital/dp-frontend-router) - use: `$ make debug DATASET_ROUTES_ENABLED=true`
* For every service in [dataset](#dataset-journey) and [filter](#filter-journey)- use: `make debug ENABLE_PRIVATE_ENDPOINTS=true`
* [dp-dimension-extractor](https://github.com/ONSdigital/dp-dimension-extractor) - use: `$ make debug ENCRYPTION_DISABLED=true`
* [dp-observation-extractor](https://github.com/ONSdigital/dp-observation-extractor) - use `$ make debug ENCRYPTION_DISABLED=true`

#### CMD Web

If you already have content, and you just want to run the web journey, you'll need the [dataset](#dataset-journey), [filter](#filter-journey) and [web](#web-journey) services. Again, use the commands:

* [florence](https://github.com/ONSdigital/florence) - use: `$ make debug ENCRYPTION_DISABLED=true`
* [dp-frontend-router](https://github.com/ONSdigital/dp-frontend-router) - use: `$ make debug`
* unset `ENABLE_PRIVATE_ENDPOINTS`

Return to the [Getting Started](https://github.com/ONSdigital/dp/blob/master/guides/GETTING_STARTED.md) guide for next steps.

--------------

## Setup credentials

To run florence, you will need to update the environment variable `SERVICE_AUTH_TOKEN` in your [startup file](#startup-file). Steps for creating the service authentication token can be found in the [Zebedee repository](https://github.com/ONSdigital/zebedee/#service-authentication-with-zebedee).

You will need to restart your terminal for the environment variable change to take effect.

Note that when the first login to a Florence account is detected a mandatory password update is required.
