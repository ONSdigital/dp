Getting started
===============

#### Overview

Note: [Websysd](https://github.com/ONSdigital/dp/tree/master/websysd) can accomplish steps 1-4 for you. Alternatively see [iTerm2 steps](#iTerm2-setup)

1. Install the [prerequisites](#prerequisites)

2. Clone the GitHub repos for [web](#web), [publishing](#publishing) and/or [CMD](#cmd). To get everything running follow the [CMD](#cmd) steps

3. [Setup Zebedee content](#setup-zebedee-content)

4. Run the apps:

  - [web](#web)
  - [publishing](#publishing)
  - and optionally [CMD](#cmd).

5. [Configure the apps](#configure)

If you need to edit content in Florence, [try this guide](https://github.com/ONSdigital/florence/blob/develop/USAGE.md)

Once setup is complete, see the [Contributing guidelines](CONTRIBUTING.md)

#### For CMD:

__Additional app setup:__

6. If you want the feedback forms to work in the Beta banners, [follow this guidance](https://github.com/ONSdigital/dp-frontend-dataset-controller#feedback-service)

7. [Setup API keys in Zebedee](https://github.com/ONSdigital/zebedee/tree/develop#service-authentication-with-zebedee)

__Content:__ Your apps should startup and be functional at this point, but you'll have no CMD content.

8. [Import code lists](https://github.com/ONSdigital/dp-code-list-scripts#import-to-a-new-development-environment)

9. [Import hierarchies](https://github.com/ONSdigital/dp-hierarchy-builder#getting-started)

__Florence steps:__

10. [Create datasets](https://github.com/ONSdigital/florence/tree/develop/USAGE.md#create-a-cmd-dataset-page)

11. [Import V4 file](https://github.com/ONSdigital/florence/tree/develop/USAGE.md#import-a-v4-file)

12. Publish content [as normal](https://github.com/ONSdigital/florence/tree/develop/USAGE.md#publish-a-collection)

-----
### Prerequisites

* [Java 8 JDK (OpenJDK)](https://openjdk.java.net/install/)
* [Maven](https://maven.apache.org/)
* [Docker](https://www.docker.com/get-started)
* [Cypher Shell](https://neo4j.com/docs/operations-manual/current/tools/cypher-shell/) - installed with `brew install cypher-shell`
* [Node.js and npm](https://nodejs.org/en/), known working versions:
  - Node.js version 10.15.3
  - npm version 6.9.0
* [dp-compose](https://github.com/ONSdigital/dp-compose)
  - Elasticsearch 2.4.2
  - Highcharts
  - Postgres
* [go v1.13](https://golang.org/doc/install)
* [GoConvey](https://github.com/smartystreets/goconvey#installation)
* [Govendor](https://github.com/kardianos/govendor)
* [GhostScript](https://www.ghostscript.com/download.html) - Required for [Babbage](https://github.com/onsdigital/babbage)
* [Vault](https://www.vaultproject.io/intro/getting-started/install.html) - This will be required for running Florence.

* [jq](https://stedolan.github.io/jq/) - installed with `brew install jq`
  - This isn't essential, but it's a useful tool for debugging website content.

For CMD additionally install:

* [MongoDB](https://docs.mongodb.com/manual/tutorial/install-mongodb-on-os-x/#install-mongodb-community-edition-with-homebrew)
* [Neo4j](https://neo4j.com/download-center/#releases) - currently limited to 3.2.12
* [Kafka v0.10.2.1](https://kafka.apache.org/downloads#0.10.2.1) - [Docs for version specific commands](https://kafka.apache.org/0102/documentation.html) 

Elasticsearch will need to be on [version 5](https://www.elastic.co/guide/en/elasticsearch/reference/5.4/gs-installation.html) to work with CMD.

### Web

The website requires the following components:

* [babbage](https://github.com/ONSdigital/babbage)
  * use `$ ./run.sh`
* [zebedee](https://github.com/ONSdigital/zebedee)
  * use `$ ./run-reader.sh`
* [sixteens](https://github.com/ONSdigital/sixteens)
  * use `$ ./run.sh`
* [dp-frontend-router](https://github.com/ONSdigital/dp-frontend-router)
  * use `$ make debug`
* [dp-frontend-renderer](https://github.com/ONSdigital/dp-frontend-renderer)
  * use `$ make debug`

### Publishing

The publishing tool requires the following components:

* [babbage](https://github.com/ONSdigital/babbage)
  * use `$ ./run-publishing.sh`
* [florence](https://github.com/ONSdigital/florence)
  * use `$ make debug ENCRYPTION_DISABLED=true`
* [zebedee](https://github.com/ONSdigital/zebedee)
  * use `$ ./run.sh`
* [sixteens](https://github.com/ONSdigital/sixteens)
  * use `$ ./run.sh`
* [The-Train](https://github.com/ONSdigital/The-Train)
  * use `$ ./run.sh`
* [dp-frontend-router](https://github.com/ONSdigital/dp-frontend-router)
  * use `$ make debug`
* [dp-frontend-renderer](https://github.com/ONSdigital/dp-frontend-renderer)
  * use `$ make debug`

#### Setup Zebedee content

To setup Zebedee content, follow the guide in [dp-zebedee-content](https://github.com/ONSdigital/dp-zebedee-content#dp-zebedee-content)

### CMD

To set up the publishing process for CMD, you will need to run all of the following *and* the [publishing](#publishing) services, note that some services require feature flags enabled such as the below:

* [dp-frontend-router](https://github.com/ONSdigital/dp-frontend-router)
  * use `$ make debug DATASET_ROUTES_ENABLED=true`
* [florence](https://github.com/ONSdigital/florence)
  * use `$ make debug ENABLE_DATASET_IMPORT=true ENCRYPTION_DISABLED=true`
*

If you already have content, and you just want to run the web journey, you'll need the [dataset](#dataset-journey), [filter](#filter-journey) and [web](#web) services.

#### Dataset journey:
* [dp-dataset-api](https://github.com/ONSdigital/dp-dataset-api)
  * use `$ make debug`
* [dp-frontend-dataset-controller](https://github.com/ONSdigital/dp-frontend-dataset-controller)
  * use `$ make debug`
* [dp-frontend-filter-dataset-controller](https://github.com/ONSdigital/dp-frontend-filter-dataset-controller)
  * use `$ make debug`
* [dp-frontend-geography-controller](https://github.com/ONSdigital/dp-frontend-geography-controller)
  * use `$ make debug`

#### Import services:
* [dp-recipe-api](https://github.com/ONSdigital/dp-recipe-api)
  * use `$ make debug`
* [dp-import-api](https://github.com/ONSdigital/dp-import-api)
  * use `$ make debug`
* [dp-import-tracker](https://github.com/ONSdigital/dp-import-tracker)
  * use `$ make debug`
* [dp-dimension-extractor](https://github.com/ONSdigital/dp-dimension-extractor)
  * use `$ make debug ENCRYPTION_DISABLED=true`
* [dp-dimension-importer](https://github.com/ONSdigital/dp-dimension-importer)
  * use `$ make debug`
* [dp-observation-extractor](https://github.com/ONSdigital/dp-observation-extractor)
  * use `$ make debug ENCRYPTION_DISABLED=true`
* [dp-observation-importer](https://github.com/ONSdigital/dp-observation-importer)
  * use `$ make debug`
* [dp-hierarchy-builder](https://github.com/ONSdigital/dp-hierarchy-builder)
  * use `$ make debug`
* [dp-search-builder](https://github.com/ONSdigital/dp-search-builder)
  * use `$ make debug`
* [dp-publishing-dataset-controller](https://github.com/ONSdigital/dp-publishing-dataset-controller)
  * use `$ make debug`

#### Filter journey:
* [dp-search-api](https://github.com/ONSdigital/dp-search-api)
  * use `$ make debug`
* [dp-code-list-api](https://github.com/ONSdigital/dp-code-list-api)
  * use `$ make debug`
* [dp-hierarchy-api](https://github.com/ONSdigital/dp-hierarchy-api)
  * use `$ make debug`
* [dp-filter-api](https://github.com/ONSdigital/dp-filter-api)
  * use `$ make debug`
* [dp-dataset-exporter](https://github.com/ONSdigital/dp-dataset-exporter)
  * use `$ make debug`
* [dp-dataset-exporter-xlsx](https://github.com/ONSdigital/dp-dataset-exporter-xlsx)
  * use `$ make debug`
* [dp-download-service](https://github.com/ONSdigital/dp-download-service)
  * use `$ make debug`

[dp-api-router](https://github.com/ONSdigital/dp-api-router) locally as well.
  * use `$ make debug`


### iTerm2 setup
  - Any terminal can be used however there are advantages of using iTerm2 such as badges, profiles, arrangements, multiple panes, ability to send commands to all panes.
  - iTerm2 can be installed here: https://www.iterm2.com/
  - Setup the pane arrangements and profiles by following this tutorial https://blog.andrewray.me/how-to-create-custom-iterm2-window-arrangments/
  - It can be advantageous to give each profile a badge, instructions to complete this can be found here: https://www.iterm2.com/documentation-badges.html

#### Configure
* In Zebedee `run.sh` remove the following line: export `SERVICE_AUTH_TOKEN="fc4089e2e12937861377629b0cd96cf79298a4c5d329a2ebb96664c88df77b67"`
* Service authentication token creation steps can be found in the [Zebedee repository](https://github.com/ONSdigital/zebedee/#service-authentication-with-zebedee) 

Note that when the first login to a Florence account is detected a mandatory password update is required.

## Reaching end-user experience

### Publishing
  - Florence will be available at `http://localhost:8081/florence/login`
  - The website will be available at `http://localhost:8081` (only available after a successful login into Florence)

### Web
  - The website will be available at `http://localhost:22000`

### Useful resources
* [Splash-page](https://github.com/ONSdigital/dp-setup/tree/develop/splash-page) (if you get a 404 error following this link, you need to be added to the ONSdigital organisation in GitHub), a list of digital publishing's repositories, environments, and platform-tools
* [dp-cli](https://github.com/ONSdigital/dp-cli) is a useful tool to grant access to environments and SSH into them when working offsite as well as other tasks
