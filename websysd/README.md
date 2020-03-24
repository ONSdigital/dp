# Websyd

## Warning
websysd assumes that many services are located in the GOPATH. Since most have migrated to Go modules, this is no longer
the case and these scripts will need to be updated in order to work.

## Original documentation…
You can use [websysd](https://github.com/ONSdigital/dp/blob/master/websysd) to run the dp services from a centralised location:

If you are running the publishing stack, then you must set "ENABLE_PRIVATE_ENDPOINTS=true" and "ENCRYPTION_DISABLED=true” if encryption is not setup.

If you are running the publishing stack in terminals [dp-dimension-extractor](https://github.com/ONSdigital/dp-dimension-extractor), [dp-observation-extractor](https://github.com/ONSdigital/dp-observation-extractor) and [florence](https://github.com/ONSdigital/florence) should have "ENCRYPTION_DISABLED=true” if encryption is not setup. For [dp-dataset-api](https://github.com/ONSdigital/dp-dataset-api), [dp-filter-api](https://github.com/ONSdigital/dp-filter-api) and [dp-search-api](https://github.com/ONSdigital/dp-search-api) you require having "ENABLE_PRIVATE_ENDPOINTS=true".

## Installation

To install a dp-environment on Mac OS from scratch you will first need git access to all the dp services.

Once you have this you will be able to run:

`./dp-run.sh -i` to install all dependencies and services.

If you believe you already have the dependencies installed and only wish to clone
the services run:

`./dp-run.sh -c`

Before you run any services, ensure you have AWS access and set the following environment vars:

- export AWS_ACCESS_KEY_ID=<ACCESS_KEY>
- export AWS_SECRET_ACCESS_KEY=<SECRET_KEY>

replacing the ACCESS/SECRET keys with your own.

Run the following to pull all the repositories on the current branch

`./dp-run.sh -p`

## Running websysd

To run all the services simply run:

`./dp-run.sh`

This will preload all data stores with required data and start all services on websysd. To view the status of the services, navigate your browser to:

http://localhost:7050
