CI/CD Process - Deploying Apps into Environments
===============================================

The information in this guide should provide understanding on what CI/CD is, how the CI/CD pipeline automates our software delivery process, and how to deploy apps into our environments.

## What is CI/CD?

Continuous integration (CI) and continuous delivery (CD) are a set of operating principles, and collection of practices that enable application development teams to deliver code changes more frequently and reliably. The implementation is also known as the CI/CD pipeline.

Continuous integration most often refers to the build or integration stage of the software release process and entails both an automation component (e.g. a CI or build service) and a cultural component (e.g. learning to integrate frequently). The key goals of continuous integration are to find and address bugs quicker, improve software quality, and reduce the time it takes to validate and release new software updates.

With CI, each change in code triggers an automated build-and-test sequence for the given project, providing feedback to the developer(s) who made the change. 

When you create a PR in Github, you will see that there are three CI checks being carried out - 

- CI audit tests 
- CI build
- CI unit tests

You will be unable to merge your PR if any of these checks fail.

Continuous Delivery includes infrastructure provisioning and deployment, which may be manual and consist of multiple stages. What’s important is that all these processes are fully automated, with each run fully logged and visible to the entire team.

Continuous delivery can be referred to as an extension of Continuous Integration, which includes testing, packaging the code ready for a release and even testing after a release. The important part to note is that Continuous Delivery involves a manual step to deploy to production.

A more in-depth look at the overall principles of CI/CD and its processes can be found [here](https://www.atlassian.com/continuous-delivery/principles/continuous-integration-vs-delivery-vs-deployment).

## CI/CD at the ONS

At ONS Digital Publishing, We have a Continuous Delivery system with Git and Concourse at the heart of it. However, you will often hear it referred to as just ‘CI’.


Our CI is made up of the following - 

- Version Control/Source Control (Git + GitHub)
- Processes and practices
- Slack
- Concourse
- Docker Images
- AWS S3
- Nomad
- Consul

[Concourse](https://concourse-ci.org/) is a tool that helps us with the CI process, and is primarily used to configure and automate our pipelines. A link to the Councourse Github repository can be found [here](https://github.com/concourse/concourse).

You can access the ONS Digital Concourse app [here](https://concourse.onsdigital.co.uk/).

A more comprehensive look at how CI/CD works at ONS Digital Publishing can be found [here](https://docs.google.com/presentation/d/1_oPT6WBfWrNsIHONExp0xwUAlW_VElgu-eMxPPBR1W8/edit#slide=id.p).

## Deploying your apps

This section is a continuation of the steps outlined in the "Ready for Release" section [here](TRELLO_BOARD_FLOW.md). 

1. To deploy an app, you will need to have logged into ONS Digital Councourse via the link above. You can log in via your personal Github account, or email addresss
2. Once logged into Concourse, you should see the all of the pipelines that are currently running 
3. Find the correct pipeline for the repository you are releasing to. So, for example, if you have created a release for `dp-dataset-api`, you will need to find the corresponding pipeline. In this case it would be - https://concourse.onsdigital.co.uk/teams/main/pipelines/dp-dataset-api
4. If you have followed previous steps in the "Ready for Release" section linked above, and the changes have been pushed to the `master` branch, Concourse should start the build. You will see that the build has started running as it will be indicated by a yellow border.
5. Wait for concourse to finish the `master-unit`, `master-build`, `master-audit` and `release`,  jobs
6. Once these have been completed, click on the green `production-ship-it` task (at the far-right of the page), then at the top right of this page you will see a `+` symbol. Click this to begin shipping the code to production
7. Test the code in production and keep an eye on the `#production-alarm` and `#production-warning` slack channels for any issues