# Library choices

This is an incomplete list of libraries used for different areas by Dissemination.

This document aims to:

* define the reference libraries we use,
* provide a focus for the conversation about changing libraries and
* keep a list of our internal libraries.

## Table of Contents

* [Table of Contents](#table-of-contents)
* [Internal Library Standards](#internal-library-standards)
* [HTTP and Middleware](#http-and-middleware)
* [Config and Logging](#config-and-logging)
* [Database Drivers](#database-drivers)
* [Frontend](#frontend)
* [Testing](#testing)
  * [Unit testing frameworks](#unit-testing-frameworks)
  * [Component testing frameworks](#component-testing-frameworks)
  * [Mocking](#mocking)
  * [Integration, Automation and Accessibility](#integration-automation-and-accessibility)
* [Misc](#misc)

## Internal Library Standards

We create internal libraries where we want our interaction with a 3rd-party library to be
standardized across all of our services. Often this is as simple as adding our health
checking ability and surfacing the behaviour that's relevant to us, in other cases it
provides us the opportunity to set defaults consistently across our microservices.

Our internal libraries typically begin with `dp-` in the repo name, and *must* reside in the `ONSdigital` Github organization.

> [!NOTE]
> Go: It is acceptable for files to live at the top level of these repos and take the single word name after `dp-` as the package name. e.g. `dp-kafka` uses the `package kafka` definition

These client libraries all serve that high-level wrapping purpose around 3rd party services:

* [dp-elasticsearch](https://github.com/ONSdigital/dp-elasticsearch)
* [dp-nomad](https://github.com/ONSdigital/dp-nomad)
* [dp-vault](https://github.com/ONSdigital/dp-vault)
* [dp-mongodb](https://github.com/ONSdigital/dp-mongodb)
* [dp-s3](https://github.com/ONSdigital/dp-s3)
* [dp-kafka](https://github.com/ONSdigital/dp-kafka)
* [dp-otel-go](https://github.com/ONSdigital/dp-otel-go)
* [dp-graph](https://github.com/ONSdigital/dp-graph)
* [dp-slack-client-java](https://github.com/ONSdigital/dp-slack-client-java)

As well as libraries we use for communication between internal services:

* [dp-api-clients-go](https://github.com/ONSdigital/dp-api-clients-go) which the majority of our API clients. More recent applications house their own API clients (see [dp-topic-api](https://github.com/ONSdigital/dp-topic-api/tree/develop/sdk))
* [dp-dataset-api-java-client](https://github.com/ONSdigital/dp-dataset-api-java-client)
* [dp-static-files-api-client-java](https://github.com/ONSdigital/dp-static-files-api-client-java)
* [dp-upload-service-client-java](https://github.com/ONSdigital/dp-upload-service-client-java)
* [dp-image-api-client-java](https://github.com/ONSdigital/dp-image-api-client-java)
* [dp-reporter-client](https://github.com/ONSdigital/dp-reporter-client) which publishes kafka message throughout the data import process
* [dp-zebedee-sdk-go](https://github.com/ONSdigital/dp-zebedee-sdk-go)

The underlying drivers may be highlighted in further sections of this document, but where a library exists it should be used and updated rather than importing drivers directly to services.

## HTTP and Middleware

Internal libraries:

* [dp-net](https://github.com/ONSdigital/dp-net) wraps the Go `net/http` library with our preferred defaults
* [dp-cookies](https://github.com/ONSdigital/dp-cookies)
* [dp-healthcheck](https://github.com/ONSdigital/dp-healthcheck)
* [dp-authorisation](https://github.com/ONSdigital/dp-authorisation)
* [dp-frontend-cache-helper](https://github.com/ONSdigital/dp-frontend-cache-helper)
* [dp-cache](https://github.com/ONSdigital/dp-cache)

Internal forks of community libraries:

* ~~[restolino](https://github.com/ONSdigital/restolino) is our legacy Java REST framework~~ **DEPRECATED**
* ~~[httpino](https://github.com/ONSdigital/httpino) provides a straightforward HTTP client.~~ **DEPRECATED**

3rd party libraries:

* [spark](http://sparkjava.com/) as our preferred java framework
* [jetty](https://www.eclipse.org/jetty/) for servers in java
* [alice](https://github.com/justinas/alice) for middleware chaining
* [gorilla/mux](http://github.com/gorilla/mux) for route definitions

## Config and Logging

Internal libraries:

* [log.go](https://github.com/ONSdigital/log.go) for go logging
* [dp-logging](https://github.com/ONSdigital/dp-logging) for java logging

 3rd party libraries:

* [envconfig](https://github.com/kelseyhightower/envconfig) for environment based configuration in Go

## Database Drivers

Internal libraries:

* ~~[s3crypto](https://github.com/ONSdigital/s3crypto) to encrypt/decrypt a file stream to S3~~ **DEPRECATED**
* [java-s3crypto](https://github.com/ONSdigital/java-s3crypto) to encrypt/decrypt a file stream to S3 but now in Java
* [dp-graph](https://github.com/ONSdigital/dp-graph) wraps all graph behaviour behind a common interface
* dp-mongo, dp-s3 and dp-kafka all as mentioned at the top of this page

Internal forks of community drivers:

* [neo4j-bolt](https://github.com/ONSdigital/golang-neo4j-bolt-driver)
* [gremgo-neptune](https://github.com/ONSdigital/gremgo-neptune)
* [graphson](https://github.com/ONSdigital/graphson) to interpret gremlin responses

3rd party drivers:

* Java
  * [hibernate](https://hibernate.org/) for java database access - specifically used for Postgres access currently
* Go
  * [mgo](https://github.com/globalsign/mgo) for mongodb
  * [sarama](https://github.com/Shopify/sarama) for kafka
  * [sarama-cluster](https://github.com/bsm/sarama-cluster) for kafka clustering
  * [aws-sdk](https://github.com/aws/aws-sdk-go) for S3 and other AWS needs
* JS
  * [Firebase](https://firebase.google.com/docs/reference/js)
  * [MiniMongo](https://www.npmjs.com/package/minimongo)

## Frontend

Internal libraries

* ~~[sixteens](https://github.com/ONSdigital/sixteens) legacy SCSS/JS library for Babbage~~ **DEPRECATED**
* [dp-design-system](https://github.com/ONSdigital/dp-design-system) SCSS/JS wrapper / extension library for the [ONS design system](https://github.com/ONSdigital/design-system)
* [dp-renderer](https://github.com/ONSdigital/dp-renderer) go templating library for frontend applications

## Testing

### Unit testing frameworks

3rd party libraries:

* [junit](https://junit.org/junit5/) for java
* [goconvey](https://github.com/smartystreets/goconvey) for go
* [JestJS](https://jestjs.io/) for JS

### Component testing frameworks

Internal libraries

* [dp-component-test](https://github.com/ONSdigital/dp-component-test) provides a framework for BDD component testing using [godog](https://github.com/cucumber/godog)

### Mocking

Internal libraries:

* [dp-mocking](https://github.com/ONSdigital/dp-mocking) used by api-clients-go

3rd party libraries:

* [mockito](https://github.com/mockito/mockito) in Java
* [MockJS](https://www.npmjs.com/package/mockjs) in JS
* [moq](https://github.com/matryer/moq) in Go

### Integration, Automation and Accessibility

3rd party libraries:

* [BrowserStack](https://www.browserstack.com/) for cross browser compatibility testing
* [Puppeteer](https://github.com/GoogleChrome/puppeteer) for automation testing of internal services

## Misc

Internal libraries:

* ~~[go-ns](https://github.com/ONSdigital/go-ns) contains many library packages that are slowly being broken out into their own libraries~~ **DEPRECATED**
* ~~[dp-bolt](https://github.com/ONSdigital/dp-bolt) is a high level wrapper around neo4j's bolt driver - this has been superceded by dp-graph.~~ **DEPRECATED**

Internal forks of community libraries:

* ~~[dp-cryptolite-java](https://github.com/ONSdigital/dp-cryptolite-java) provides encryption for our Java publishing services.~~ **DEPRECATED**
* ~~[resourece-utils](https://github.com/ONSdigital/resource-utils) provides management of Java resources~~ **DEPRECATED**
* ~~[encrypted-file-upload](https://github.com/ONSdigital/encrypted-file-upload) encryption for file uploads~~ **DEPRECATED**
* [dp-ssqs](https://github.com/ONSdigital/dp-ssqs) provides a high level wrapper for AWS SQS message consumers, but does not currently meet DP library standards
