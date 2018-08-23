Getting started
===============

#### The Basics

1. Install the [prerequisites](#prerequisites)

2. Clone the GitHub repos for [web](#web), [publishing](#publishing) and/or [CMD](#cmd)

3. [Setup Zebedee content](#setup-zebedee-content)

4. Run the apps [web](#web), [publishing](#publishing) and/or [CMD](#cmd)
  - Florence will be available at `http://localhost:8081/florence/login`
  - The website will be available at `http://localhost:8081`

[Websysd](https://github.com/ONSdigital/dp/tree/master/websysd) can accomplish steps 1-4 for you.

If you need to edit content in Florence, [try this guide](https://github.com/ONSdigital/florence/tree/cmd-develop/USAGE.md)

#### For CMD:

__Additional app setup:__
6. If you want the feedback forms to work in the Beta banners, [follow this guidance](https://github.com/ONSdigital/dp-frontend-dataset-controller#feedback-service)

7. [Setup API keys in Zebedee](https://github.com/ONSdigital/zebedee#service-authentication-with-zebedee)

__Content:__ Your apps should start up/be functional at this point but you'll have no CMD content.

8. [Import code lists](https://github.com/ONSdigital/dp-code-list-scripts#import-to-a-new-development-environment)

9. [Import hierarchies](https://github.com/ONSdigital/dp-hierarchy-builder#getting-started)

__Florence steps:__
10. [Create datasets](https://github.com/ONSdigital/florence/tree/cmd-develop/USAGE.md#create-a-cmd-dataset)

11. [Import V4 file](https://github.com/ONSdigital/florence/tree/cmd-develop/USAGE.md#import-a-v4-file)

12. Publish content [as normal](https://github.com/ONSdigital/florence/tree/cmd-develop/USAGE.md#publish-a-collection)

-----
### Prerequisites

* [Java 8 JDK](http://www.oracle.com/technetwork/java/javase/downloads/jdk8-downloads-2133151.html)
* [Maven](https://maven.apache.org/)
* [Docker](https://www.docker.com/products/overview)
* [Node.js and npm](https://nodejs.org/en/), known working versions:
  - Node.js version 8.9.4
  - npm version 5.6.0
* [dp-compose](https://github.com/ONSdigital/dp-compose)
  - Elasticsearch 2.4.2
  - Highcharts
  - Postgres
* [go](https://golang.org/doc/install)

For CMD additionally install:

* [MongoDB](https://docs.mongodb.com/manual/tutorial/install-mongodb-on-os-x/#install-mongodb-community-edition-with-homebrew)
* [Neo4j](https://neo4j.com/download-center/#releases) - currently limited to 3.2.12
* [Kafka](https://kafka.apache.org/quickstart)

Elasticsearch will need to be on [version 5](https://www.elastic.co/guide/en/elasticsearch/reference/5.4/gs-installation.html) to work with CMD.

### Web

The website requires the following components:

* [babbage](https://github.com/ONSdigital/babbage)
  * use `run.sh`
* [florence](https://github.com/ONSdigital/florence)
* [zebedee](https://github.com/ONSdigital/zebedee)
  * use `run-reader.sh`
* [sixteens](https://github.com/ONSdigital/sixteens)
* [dp-frontend-router](https://github.com/ONSdigital/dp-frontend-router)
* [dp-frontend-renderer](https://github.com/ONSdigital/dp-frontend-renderer)

### Publishing

The publishing tool requires the following components:

* [babbage](https://github.com/ONSdigital/babbage)
  * use `run-publishing.sh`
* [florence](https://github.com/ONSdigital/florence)
* [ermintrude](https://github.com/ONSdigital/ermintrude)
* [zebedee](https://github.com/ONSdigital/zebedee)
  * use `run.sh`
* [sixteens](https://github.com/ONSdigital/sixteens)
* [The-Train](https://github.com/ONSdigital/The-Train)
* [project-brian](https://github.com/ONSdigital/project-brian)
* [dp-frontend-router](https://github.com/ONSdigital/dp-frontend-router)
* [dp-frontend-renderer](https://github.com/ONSdigital/dp-frontend-renderer)

#### Setup Zebedee content
Zebedee requires `zebedee_root` to be set and to exist before starting
the Zebedee CMS server, for example:

```bash
export zebedee_root=../zebedee-content
mkdir $zebedee_root
./run.sh
```

### CMD

* [dp-dataset-api](https://github.com/ONSdigital/dp-dataset-api)
* [dp-frontend-dataset-controller](https://github.com/ONSdigital/dp-frontend-dataset-controller)
* [dp-frontend-filter-dataset-controller](https://github.com/ONSdigital/dp-frontend-filter-dataset-controller)
* [dp-frontend-renderer](https://github.com/ONSdigital/dp-frontend-renderer)
* [dp-frontend-router](https://github.com/ONSdigital/dp-frontend-router)

* [dp-recipe-api](https://github.com/ONSdigital/dp-recipe-api)
* [dp-import-api](https://github.com/ONSdigital/dp-import-api)
* [dp-import-tracker](https://github.com/ONSdigital/dp-import-tracker)
* [dp-dimension-extractor](https://github.com/ONSdigital/dp-dimension-extractor)
* [dp-dimension-importer](https://github.com/ONSdigital/dp-dimension-importer)
* [dp-observation-extractor](https://github.com/ONSdigital/dp-observation-extractor)
* [dp-observation-importer](https://github.com/ONSdigital/dp-observation-importer)
* [dp-hierarchy-builder](https://github.com/ONSdigital/dp-hierarchy-builder)
* [dp-search-builder](https://github.com/ONSdigital/dp-search-builder)

* [dp-search-api](https://github.com/ONSdigital/dp-search-api)
* [dp-code-list-api](https://github.com/ONSdigital/dp-code-list-api)
* [dp-hierarchy-api](https://github.com/ONSdigital/dp-hierarchy-api)
* [dp-filter-api](https://github.com/ONSdigital/dp-filter-api)
* [dp-dataset-exporter](https://github.com/ONSdigital/dp-dataset-exporter)
* [dp-dataset-exporter-xlsx](https://github.com/ONSdigital/dp-dataset-exporter-xlsx)
* [dp-download-service](https://github.com/ONSdigital/dp-download-service)

Make sure other apps are on CMD branches, optionally try and run the [dp-api-router](https://github.com/ONSdigital/dp-api-router) locally as well.

The majority of these apps will run with `make debug`, except the XLSX exporter which requires `./run.sh`
