# Websyd

You can use websysd to run your dp services locally. To start your services
simply run:

```
./websysd_darwin_amd64
```

from this directory. Navigate to http://localhost:7050 to view microservices and
their logs.

## Pre-requisites

To start websyd you will need the following installed on your machine:

- Java
- Maven
- Go
- Postgres
- Kafka

In the case of postgres and kafka, you should also ensure that those services
are running.

You will also need to configure postgres by running the following commands:

- Run ```brew install postgres```
- Run ```brew services start postgres```
- Run ```createuser dp -d -w```
- Run ```createdb --owner dp FilterJobs```
- Run ```psql -U dp FilterJobs -f scripts/InitDatabase.sql```

You will need the following services cloned:

### Go services:

Note this must be cloned under your $GOPATH/src/github.com/ONSdigital:

- dp-frontend-router
- dp-frontend-renderer
- dp-frontend-datset-controller
- dp-frontend-filter-dataset-controller
- dp-filter-api
- dp-code-list-api
- dp-hierarchy-api
- dp-import-api
- dp-recipe-api
- florence

### Other services:

Note this must be cloned at the top level directory ($HOME):

- dp-compose
- zebedee
- babbage
- sixteens

## Adding services

To add a new Go service to websysd, simply extend the workspace.json file to include
the new service in the same format as the other Go services.

To add a service in any other language you will need to add a new bash file to start
that service as well as extending workspace.json.
