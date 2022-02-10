Creating a new app
==================

The following steps will guide you through adding a new app to our system that is built by the CI system and deployed to the live environment.

**NOTE:** If you're creating a new library go to [new library](NEW_LIBRARY.md)

Prerequisites
-------------

To follow this guide you will need:

* A Github account with access to the ONSdigital organisation

* A login for the concourse CI environment

* [dp-cli](https://github.com/ONSdigital/dp-cli)

* Git:

  `$ brew install git`

* Terraform, in document `dp-setup/terraform/README.md` install the version specified in the Prerequisites section : [dp-setup](https://github.com/ONSdigital/dp-setup/blob/develop/terraform/README.md#prerequisites) 

  If for example that specifies version 0.13, then do:
  
  `$ brew install terraform@0.13`

* Fly CLI (can be downloaded from the links in the corner of the concourse UI)

Getting started
---------------

These steps will create a new app that is built by CI and deployed to the develop and production environments.  If any of the following Github links 404, likely you need to be added to the ONSdigital organisation.

1. Determine what language and technologies you will use
1. Pick a local dev port for your app and add it to the [local ports list](PORTS.md)
1. Use [`dp-cli` to generate the new repo with boilerplate files](https://github.com/ONSdigital/dp-cli/tree/main/project_generation/COMPLETE_PROJECT_SETUP.md)
1. [Add the new app to the environment](https://github.com/ONSdigital/dp-setup#adding-a-new-app)
1. [Add configurations for your new app](https://github.com/ONSdigital/dp-configs#adding-a-new-app)
1. [Get your pipeline boostrapped](https://github.com/ONSdigital/dp-ci/blob/670acd663353ae6b641e2ae39b879d46f4d60e64/pipelines/README.md#bootstrapping-the-pipelines) note: this will need to be done by someone with permissions to concourse gpg key  
1. Check the [Concourse UI](https://github.com/ONSdigital/dp/blob/main/training/platform-services/PLATFORM.md#concourse) and/or [Nomad UI](https://github.com/ONSdigital/dp/blob/main/training/platform-services/PLATFORM.md#nomad) to ensure your new app has successfully deployed to develop. See [here](https://github.com/ONSdigital/dp-setup/tree/develop/scripts#ansible-vault-helpers) for info on getting any acl tokens.
