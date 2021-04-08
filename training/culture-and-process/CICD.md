CI/CD Process - Deploying Apps into Environments
===============================================

The information in this guide should provide understanding on what CI/CD is, how the CI/CD pipeline automates our software delivery process, and how to deploy apps into our environments.

## What is CI/CD?

Continuous integration (CI) and continuous delivery (CD) are a set of operating principles, and collection of practices that enable application development teams to deliver code changes more frequently and reliably. The implementation is also known as the CI/CD pipeline.

With CI, each change in code triggers an automated build-and-test sequence for the given project, providing feedback to the developer(s) who made the change. 

When you create a PR in Github, you will see that there are three CI checks being carried out - 

- CI audit tests 
- CI build
- CI unit tests

You will be unable to merge your PR if any of these checks fail.

Continuous Delivery includes infrastructure provisioning and deployment, which may be manual and consist of multiple stages. What’s important is that all these processes are fully automated, with each run fully logged and visible to the entire team.

A more in-depth look at the overall principles of CI/CD and its processes can be found [here](https://www.atlassian.com/continuous-delivery/principles/continuous-integration-vs-delivery-vs-deployment).

## CI/CD at the ONS

The CI/CD system used by ONS Digital is called [Concourse](https://concourse-ci.org/). A link to the Councourse Github repository can be found [here](https://github.com/concourse/concourse).

Concourse is an open source CI/CD tool that ONS Digital uses to configure and automate its pipelines.

You can access the ONS Digital Concourse app [here](https://concourse.onsdigital.co.uk/).



## Deploying your apps

This section is a continuation of the steps outlined in the "Ready for Release" section [here](culture-and-process/PULL_REQUEST_GUIDANCE.md). 

1. To deploy an app, you will need to have logged into ONS Digital Councourse via the link above. You can log in via your personal Github account, or email addresss
2. Once logged into Concourse, you should see the all of the pipelines that are currently running 
3. Find the correct pipeline for the repository you are releasing to. So, for example, if you have created a release for `dp-dataset-api`, you will need to find the corresponding pipeline. In this case it would be - https://concourse.onsdigital.co.uk/teams/main/pipelines/dp-dataset-api
4. If you have followed previous steps in the "Ready for Release" section linked above, and the changes have been pushed to the `master` branch, Concourse should start the build. You will see that the build has started running as it will be indicated by a yellow border.
5. Wait for concourse to finish the `master-unit`, `master-build`, and `release` jobs
6. Once these have been completed, click on the green `production-ship-it` task (at the far-right of the page), then at the top right of this page you will see a `+` symbol. Click this to begin shipping the code to production
7. Test the code in production and keep an eye on the `#production-alarm` and `#production-warning` slack channels for any issues