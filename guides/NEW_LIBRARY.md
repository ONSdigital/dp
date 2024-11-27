Creating a new library
======================

The following steps will guide you through adding a new **library** to our system that is built by the CI system and deployed to the live environment.

**NOTE:** If you're creating a new app go to [new app](NEW_APP.md)

Prerequisites
-------------

To follow this guide you will need:

* A Github account with access to the ONSdigital organisation

* A login for the concourse CI environment

* [dp-cli](https://github.com/ONSdigital/dp-cli)

* Git:

  `$ brew install git`

* Fly CLI (can be downloaded from the links in the corner of the [Concourse UI](https://concourse.dp-ci.aws.onsdigital.uk/) UI)

Getting started
---------------

These steps will create a new app that is built by CI and deployed to the develop and production environments. If any of the following Github links 404, likely you need to be added to the ONSdigital organisation.

1. Determine what language and technologies you will use

2. Use [`dp-cli` to generate the new repo with boilerplate files](https://github.com/ONSdigital/dp-cli/tree/master/project_generation/COMPLETE_PROJECT_SETUP.md)

3. [Add a manifest file for your app](https://github.com/ONSdigital/dp-configs/blob/main/manifests/README.md)

   - If you are creating a Go library, use either `library` or `library-v2` manifest.
   - If you are creating a JavaScript library, use the `npm-pkg` manifest.

4. Check the Concourse UI and/or Nomad UI to ensure your new app has successfully deployed to develop
