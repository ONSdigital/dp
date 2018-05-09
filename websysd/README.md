# Websyd

You can use websysd to run your dp services locally.

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

Note that you will need to wait around 30 seconds for the full stack to start correctly. The services will
cycle through a restart state until kafka has successfully started.

To view the status of the services, navigate your browser to:

http://localhost:7050