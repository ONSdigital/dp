Getting started
===============

1. Install the prerequisites
2. Clone the GitHub repos
3. Start the services you need for either Web or Publishing

#### Prerequisites

* [Java 8 JDK](http://www.oracle.com/technetwork/java/javase/downloads/jdk8-downloads-2133151.html)
* [Maven](https://maven.apache.org/)
* [Docker](https://www.docker.com/products/overview)
* [Node.js and npm](https://nodejs.org/en/), known working versions:
  - Node.js version 6.9.1
  - npm version 3.10.8
* [dp-compose](https://github.com/ONSdigital/dp-compose)

#### Web

The website requires the following components:

* [babbage](https://github.com/ONSdigital/babbage)
  * use `run.sh`
* [florence](https://github.com/ONSdigital/florence)
* [zebedee](https://github.com/ONSdigital/zebedee)
  * use `run-reader.sh`
* [sixteens](https://github.com/ONSdigital/sixteens)

#### Publishing

Zebedee requires `zebedee_root` to be set and to exist before starting
the Zebedee CMS server, for example:

```bash
export zebedee_root=../zebedee-content
mkdir $zebedee_root
./run.sh
```

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
