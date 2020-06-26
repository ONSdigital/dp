Library choices
==================

This is a partially-complete list of libraries used for different areas by Digital Publishing.

The aims of this document are to define standards around what libraries are being used, provide a place to center conversation about changing libraries and be a reference point to the existance of internal libraries.

## Table of Contents
1. [Internal Library Standards](#Internal-Library-Standards)
2. [HTTP and Middleware](#HTTP-and-middleware)
3. [Config & Logging](#config-&-logging)
4. [Database Drivers](#database-drivers)
5. [Testing](#testing)
6. [Misc.](#misc.)

# Internal Library Standards

We create internal libraries where we want our interaction with a 3rd-party library to be
standardized across all of our services. Often this is as simple as adding our health 
checking ability and surfacing the behaviour that's relevant to us, in other cases it 
provides us the opportunity to set defaults consistently across our micro services.

Our internal libraries typically begin with `dp-` in the repo name, and *must* reside in the
`ONSdigital` Github organization. 
  * Go note: It is acceptable for files to live at the top level of these repos and take the single word name after `dp-` as the package name. e.g. `dp-kafka` uses the `package kafka` definiton

These client libraries all serve that high level wrapping purpose:
  * [dp-api-clients-go](https://github.com/ONSdigital/dp-api-clients-go) - which houses all of our API clients
  * [dp-elasticsearch](https://github.com/ONSdigital/dp-elasticsearch)
  * [dp-nomad](https://github.com/ONSdigital/dp-nomad)
  * [dp-vault](https://github.com/ONSdigital/dp-vault)
  * [dp-mongodb](https://github.com/ONSdigital/dp-mongodb)
  * [dp-s3](https://github.com/ONSdigital/dp-s3)
  * [dp-kafka](https://github.com/ONSdigital/dp-kafka)

The underlying drivers may be highlighted in further sections of this document, but where a library exists it should be used and updated rather than importing drivers directly to services.


# HTTP and Middleware

Internal libraries
* [dp-net](https://github.com/ONSdigital/dp-net) wraps the Go `net/http` library with our preferred defaults
* [dp-cookies](https://github.com/ONSdigital/dp-cookies)
* [dp-healthcheck](https://github.com/ONSdigital/dp-healthcheck)
* [dp-authorisation](https://github.com/ONSdigital/dp-authorisation)

Internal forks of community libraries
* ~~[restolino](https://github.com/ONSdigital/restolino) is our legacy Java REST framework~~ **DEPRECATED**

3rd party libraries
  * [spark](http://sparkjava.com/) as our preferred java framework
  * [jetty](https://www.eclipse.org/jetty/) for servers in java
  * [alice](https://github.com/justinas/alice) for middleware chaining
  * [gorilla/mux](http://github.com/gorilla/mux) for route definitions
  
  
# Config & Logging
Internal libraries
  * [log.go](https://github.com/ONSdigital/log.go) for go logging
  * [dp-logging](https://github.com/ONSdigital/dp-logging) for java logging

 3rd party libraries
  * [envconfig](https://github.com/kelseyhightower/envconfig) for environment based configuration in Go

  
# Database Drivers

Internal libraries
  * [s3crypto](https://github.com/ONSdigital/s3crypto) to encrypt/decrypt a file stream to S3
  * [dp-graph](https://github.com/ONSdigital/dp-graph) wraps all graph behaviour behind a common interface
  * dp-mongo, dp-s3 and dp-kafka all as mentioned at the top of this page

Internal forks of community drivers
  * [neo4j-bolt](https://github.com/ONSdigital/golang-neo4j-bolt-driver)
  * [gremgo-neptune](https://github.com/ONSdigital/gremgo-neptune)
  * [graphson](https://github.com/ONSdigital/graphson) to interpret gremlin responses

3rd party drivers
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


# Testing

### Unit testing frameworks

3rd party libraries
* [junit](https://junit.org/junit5/) for java
* [goconvey](https://github.com/smartystreets/goconvey) for go
* [JestJS](https://jestjs.io/) for JS

### Mocking

Internal libraries
  * [dp-mocking](https://github.com/ONSdigital/dp-mocking) used by api-clients-go

3rd party libraries
  * [mockito](https://github.com/mockito/mockito) in Java
  * [MockJS](https://www.npmjs.com/package/mockjs) in JS
  * [moq](https://github.com/matryer/moq) in Go

### Integration, Automation and Accessibility

3rd party libraries
  * [BrowserStack](https://www.browserstack.com/) for cross browser compatibility testing
  * [Puppeteer](https://github.com/GoogleChrome/puppeteer) for automation testing of internal services

# Misc.

Internal libraries
  * [go-ns](https://github.com/ONSdigital/go-ns) contains many library packages that are slowly being broken out into their own libraries
  * [dp-bolt](https://github.com/ONSdigital/dp-bolt) is a high level wrapper around neo4j's bolt driver - this has been superceded by dp-graph
  * [dp-ssqs](https://github.com/ONSdigital/dp-ssqs) provides a high level wrapper for AWS SQS message consumers, but does not currently meet DP library standards

Internal forks of community libraries
  * [dp-cryptolite-java](https://github.com/ONSdigital/dp-cryptolite-java) provides encryption for our Java publishing services.

