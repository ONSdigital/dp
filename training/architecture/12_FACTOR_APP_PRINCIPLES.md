12 Factor App Principles
===========================

In Dissemination we build, deploy and manage applications which run as services. The twelve-factor methodology helps us follow best practices and build applications which are portable and resilient. This module will introduce you to the Twelve-factor principles and provide you with examples of how we use them in our team.

Note: The examples refer primarily to Go apps. Please add more details or alternative examples which might be useful.

## Pre-reading
- Introduction and Background to [the twelve-factor app](https://12factor.net/)
- [About version control](https://git-scm.com/book/en/v2/Getting-Started-About-Version-Control)
- [What is Continuous Integration (CI) and Continuous Delivery (CD)?](https://www.redhat.com/en/topics/devops/what-is-ci-cd)


## Prerequisites
See Prerequisites in each section.

## Materials

In this section you can find out how we apply each principle in Dissemination:

### 1. Codebase
Prerequisite: [Codebase](https://12factor.net/codebase)

> One codebase tracked in revision control, many deploys

The app should be tracked using version control. In dissemination, We create one GitHub repository per application, for example the [dp-image-api](https://github.com/ONSdigital/dp-image-api) repository contains our codebase for that app.
Each _deploy_ of the `dp-image-api` will use a version of the codebase. The version is defined by an annotated tag which references a specific commit in the `dp-image-api` repository. The `production` environment will contain a version deployed from the `master`/`main` branch of the repository, which may be different from the `develop` branch.

Find out more:
- How we [collaborate on Dissemination repos](../../guides/CONTRIBUTING.md).
- [Semantic Versioning](https://semver.org/)

### 2. Dependencies
Prerequisite: [Dependencies](https://12factor.net/dependencies)

> Explicitly declare and isolate dependencies

A 12-factor app will need to have ownership of its dependencies by defining the version/release of a library ([what's the difference between an app and a library?](https://stackoverflow.com/questions/1270729/difference-between-library-and-application-code)). This allows for better control of how the application functions instead of always inheriting the latest version of a library which might not be backward compatible.

We use dependency managers to be able to maintain them all: [Maven](https://maven.apache.org/guides/introduction/introduction-to-dependency-mechanism.html) for Java apps and [modules](https://blog.golang.org/using-go-modules) for Go apps.

### 3. Config
Prerequisite: [Config](https://12factor.net/config)

> Store config in the environment

The configuration we want for a given app may vary for each environment (develop and production) and local setup. In order to change the configuration without modifying code, we set it using environment variables which overwrite default values. In [this example](https://github.com/ONSdigital/dp-image-api/blob/develop/config/config.go), the `config/config.go` contains a default value for `BindAddr` of `":24700"`; to overwrite it, you can set the environment variable `BIND_ADDR` to a different value.

We aim to provide open source software; if we were to hard code our configuration in the service, we might disclose information which could raise security issues. In order to protect information, it is good practice to make sure any secrets, passwords are set using config and never hard coded into an app.

The configuration for each app exists in one place so that we can collaborate and manage them easily, for example: `config/config.go` in our go applications.

In our Go apps we use the [envconfig package](https://github.com/kelseyhightower/envconfig) to help us implement this principle.


### 4. Backing services
Prerequisite: [Backing services](https://12factor.net/backing-services)

> Treat backing services as attached resources

Our apps rely on the use of other resources, such as APIs (e.g. other apps of ours; and third-party resources, like vault) and databases (e.g. MongoDB).
A 12-factor app should be able to switch out any of its resources by modifying the configuration, e.g. with no code changes.

### 5. Build, release, run
Prerequisite: [Build, release, run](https://12factor.net/build-release-run)

> Strictly separate build and run stages

We have created pipelines that build, release and run our apps.
1. Build: based on a commit, the code is converted into an executable bundle i.e. a build.
2. Release: the [dp-configs](https://github.com/ONSdigital/dp-configs) configuration is applied to the build from the previous stage.
3. Run: the app is deployed to the environment (`develop` or `production`) using the selected release.

Strict separation means that any change has to go through the stages, rather than being applied directly at runtime.

We use [Semantic Versioning](https://semver.org/) to give each release a unique id. This allows us to easily rollback to earlier releases.

Find out more:
- [Concourse documentation](../platform-services/PLATFORM.md#concourse)
- Head to [concourse.dp-ci.aws.onsdigital.uk/](https://concourse.dp-ci.aws.onsdigital.uk/) to see our pipelines.


### 6. Processes
Prerequisite: [Processes](https://12factor.net/processes)

> Execute the app as one or more stateless processes

A 12-factor app is stateless, i.e. it does not save client data from a session to be used in the next session with that client. This means that restarting the app does not result in different outcomes when processing information or requests.

In dissemination, data that should persist is stored in a stateful backing service (typically a database like Neptune or MongoDB, but also a message queue, such as kafka).

### 7. Port binding
Prerequisite: [Port binding](https://12factor.net/port-binding)

> Export services via port binding

A 12-factor app is self-contained and services are exported by port binding. In dissemination, we follow this principle, exporting [HTTP and TCP](https://www.extrahop.co.uk/company/blog/2018/tcp-vs-http-differences-explained/) as services to listen to requests and events, respectively.

Often, apps in dissemination become backing services for each other by providing their URL in the configuration. You can see examples of this by decrypting some secrets in [dp-configs](https://github.com/ONSdigital/dp-configs/tree/master/secrets). We keep track of [port allocations](https://github.com/ONSdigital/dp-setup/tree/awsb/PORTS.md).

### 8. Concurrency
Prerequisite: [Concurrency](https://12factor.net/concurrency)

> Scale out via the process model

The applications must be able to scale horizontally. This means creating multiple instances of a single application, allowing them to share the work load.

To be able to achieve this, the app must share-nothing (following the process model in [section 6](#6-processes)).

In Dissemination nearly all of our apps are compliant with this. A notable exception is [Zebedee](https://github.com/ONSdigital/zebedee) (CMS), as the data processed and served by it exists as files on disk.

### 9. Disposability
Prerequisite: [Disposability](https://12factor.net/disposability)

> Maximize robustness with fast startup and graceful shutdown

A 12 factor app should maximise robustness with fast start up and gracefully shutting down:

Fast start up is important to enable rapid deployment of new code or config and being able to scale the app horizontally or vertically quickly (this can be based on varying demand). We use [Nomad](https://www.nomadproject.io/) to manage our running services. We can easily stop, move and restart apps. If needed, we can also re-deploy apps in Concourse.

A Graceful Shutdown is important to enable an app to finish processing data before shutting down, to protect data integrity as well as responding to the requesting service or user. To achieve this the application must be able to stop handling new requests when receiving a SIGTERM and finish existing processes before shutting down.

Find out more:
- [Example code for graceful shutdown](https://github.com/ONSdigital/dp-image-api/blob/1059237022da1caa8dea37816e37508067b1d156/service/service.go#L124).

### 10. Dev/prod parity
Prerequisite: [Dev/prod parity](https://12factor.net/dev-prod-parity)

> Keep development, staging, and production as similar as possible

A 12-factor app is continuously-deployed, keeping differences between the `development` and `production` environments to a minimum. In order to do this, we need minimise the:

- Time gap: when we write code, we _can_ deploy within hours.
- Personnel gap: developers who write the code are involved in the deployment and monitoring.
- Tools gap: `develop` and `production` environments are similar. However, we limit the resources in `develop` to minimise costs.

### 11. Logs
Prerequisite: [Logs](https://12factor.net/logs)

> Treat logs as event streams

Logs should be written to `STDOUT` and the environment will decide how to stream these events using other services like [fluentd](https://github.com/fluent/fluentd).

- We have logging libraries that follow our [logging standards](https://github.com/ONSdigital/dp-standards/blob/main/LOGGING_STANDARDS.md#logging-standards). Using structured logs means we can aggregate the log events.
- The logs are streamed, centralised and aggregated in [Kibana](https://www.elastic.co/kibana)

### 12. Admin processes
Prerequisite: [Admin processes](https://12factor.net/admin-processes)

> Run admin/management tasks as one-off processes

One-off administrative tasks should be one-off processes shipped with application code. Then the process (admin task) should be run remotely, for example by SSHing to the environment.

In dissemination, we sometimes write these scripts in a separate git repo [dp-data-tools](https://github.com/ONSdigital/dp-data-tools) and hence it is not deployable; instead, we use flags to allow us to inject the correct configuration for each environment.

## Next steps


Further resources
----------------------------
- Using [dp-cli to generate a repository](https://github.com/ONSdigital/dp-cli/blob/main/project_generation/COMPLETE_PROJECT_SETUP.md)
- [Pull Request guidance](../culture-and-process/PULL_REQUEST_GUIDANCE.md)
- How we [collaborate on Dissemination repos](../../guides/CONTRIBUTING.md).
- [Semantic Versioning](https://semver.org/)
- [envconfig](https://github.com/kelseyhightower/envconfig) package
- [Concourse](https://concourse-ci.org/docs.html#docs) and [Concourse tutorial](https://concoursetutorial.com/)
- [Nomad - Getting started](https://learn.hashicorp.com/collections/nomad/get-started)
