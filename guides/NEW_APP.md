Creating a new app
==================

The following steps will guide you through adding a new app to our system that is built by the CI system and deployed to the live environment.

**NOTE:** If you're creating a new library go to [new library](NEW_LIBRARY.md)

Prerequisites
-------------

To follow this guide you will need:

* A GitHub account with access to the ONSdigital organisation

* A login for the Concourse CI environment

* [dp-cli](https://github.com/ONSdigital/dp-cli)

* Git:

  `$ brew install git`

* Terraform: follow the instructions in the [Terraform's Prerequisites section of dp-setup](https://github.com/ONSdigital/dp-setup/blob/awsb/terraform/README.md#prerequisites) to install the specified version. Do not install Terraform directly from Homebrew, as it will install the latest version by default

* Fly CLI (can be downloaded from the links in the corner of the Concourse UI)

Getting started
---------------

These steps will create a new app that is built by CI and deployed to the develop and production environments. If any of the following GitHub links 404, likely you need to be added to the ONSdigital organisation.

1. Determine what language and technologies you will use
1. Pick a local dev port for your app and add it to the [local ports list](PORTS.md)
1. Use [`dp-cli` to generate the new repo with boilerplate files](https://github.com/ONSdigital/dp-cli/tree/main/project_generation/COMPLETE_PROJECT_SETUP.md)
1. [Add the new app to the environment](https://github.com/ONSdigital/dp-setup#adding-a-new-app)
1. [Add configurations for your new app](https://github.com/ONSdigital/dp-configs#adding-a-new-app). *Note: The Concourse CI pipeline `pipeline-generator` will auto-generate the pipeline for the new app*
1. Check the [Concourse UI](https://github.com/ONSdigital/dp/blob/main/training/platform-services/PLATFORM.md#concourse) and/or [Nomad UI](https://github.com/ONSdigital/dp/blob/main/training/platform-services/PLATFORM.md#nomad) to ensure your new app has successfully deployed to develop. See [here](https://github.com/ONSdigital/dp-setup/tree/develop/scripts#ansible-vault-helpers) for info on getting any acl tokens
