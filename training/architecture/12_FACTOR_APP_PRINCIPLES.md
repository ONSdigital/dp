12 Factor App Principles
===========================

In Digital Publishing we build, deploy and manage applications which run as services. The twelve-factor methodology helps us follow best practices and build applications which are portable and resilient. This module will introduce you to the Twelve-factor principles and provide you with examples of how we use them in our team. 

Note: The examples refer primarily to Go apps because that's what I know about. Please add more details or alternative examples which might be useful.

## Pre reading

- [About version control](https://git-scm.com/book/en/v2/Getting-Started-About-Version-Control)
- [What is Continuous Integration (CI) and Continuous Delivery (CD)?](https://www.redhat.com/en/topics/devops/what-is-ci-cd)


## Pre requisites
None.

## Materials
Familiarise yourself [The Twelve-Factor App Principles](https://12factor.net/).

In this section you can find out how we apply each principle in Digital Publishing:

### 1. Codebase
Pre-requisite: [Codebase](https://12factor.net/codebase)

We create one GitHub repository per application, for example the [dp-bulletin-api](https://github.com/ONSdigital/dp-bulletin-api) repository contains our codebase for that app.
Each _deploy_ of the dp-bulletin-api will use a version of the codebase. The version is defined by an annotated tag which references a specific commit in the dp-bulletin-api repository. The `production` environment will contain a version deployed from the `master`/`main` branch of the repository, which may be different from the `develop` branch.

Find out more about how we [collaborate on Digital Publishing repos](https://github.com/ONSdigital/dp/blob/master/guides/CONTRIBUTING.md).

### 2. Dependencies
Pre-requisite: [Dependencies](https://12factor.net/dependencies)
- Apps are containerised
- Use languages that let us keep track of dependencies, e.g. the [go.mod file](https://blog.golang.org/using-go-modules) in Go.

### 3. Config
Pre-requisite: [Config](https://12factor.net/config)

We use the [envconfig package](https://github.com/kelseyhightower/envconfig) to use environment variables for configuration settings. In [this example](https://github.com/ONSdigital/dp-bulletin-api/blob/develop/config/config.go), the `config/config.go` contains a default value for `BindAddr` of `":24200"`; to overwrite it, you can set the environment variable `BIND_ADDR` to a different value.

### 4. Backing services
Pre-requisite: [Backing services](https://12factor.net/backing-services)

### 5. Build, release, run
Pre-requisite: [Build, release, run](https://12factor.net/build-release-run)

Head to [concourse.onsdigital.co.uk/](https://concourse.onsdigital.co.uk/) to see our pipelines.

### 6. Processes
Pre-requisite: [Processes](https://12factor.net/processes)

### 7. Port binding
Pre-requisite: [Port binding](https://12factor.net/port-binding)

We export HTTP as a service by binding to a port (generally the `BIND_ADDR` environment variable). We keep track of [port allocations](https://github.com/ONSdigital/dp-setup/blob/develop/PORTS.md).

### 8. Concurrency
Pre-requisite: [Concurrency](https://12factor.net/concurrency)

### 9. Disposability
Pre-requisite: [Disposability](https://12factor.net/disposability)

- We use [Nomad](https://www.nomadproject.io/) to manage our running services. We can easily stop, move and restart apps. If needed, we can also re-deploy apps in Concourse.
- Gracefully shutting down services. [Example code](https://github.com/ONSdigital/dp-bulletin-api/blob/142a6adf7a2897221f648af2a9854c26d5830622/service/service.go#L71).

### 10. Dev/prod parity
Pre-requisite: [Dev/prod parity](https://12factor.net/dev-prod-parity)

- We _can_ deploy within hours
- We write code, deploy it and monitor it i.e. no personnel gap
- Our `develop` and `production` environments are similar. However, we limit the resources in `develop` to minimise costs.

### 11. Logs
Pre-requisite: [Logs](https://12factor.net/logs)

- We have logging libraries that follow our [logging standards](https://github.com/ONSdigital/dp/blob/master/standards/LOGGING_STANDARDS.md).
- The logs are streamed, centralised and aggregated in [Kibana](https://www.elastic.co/kibana)

### 12. Admin processes
Pre-requisite: [Admin processes](https://12factor.net/admin-processes)

### Where we could improve

## Next steps


Further resources
----------------------------
- Using [dp-cli to generate a repository](https://github.com/ONSdigital/dp-cli/blob/master/project_generation/COMPLETE_PROJECT_SETUP.md)
- [Pull Request guidance](https://github.com/ONSdigital/dp/blob/master/training/culture-and-process/PULL_REQUEST_GUIDANCE.md)
- [envconfig](https://github.com/kelseyhightower/envconfig) package
- [Concourse](https://concourse-ci.org/docs.html#docs) and [Concourse tutorial](https://concoursetutorial.com/)
- [Nomad - Getting started](https://learn.hashicorp.com/collections/nomad/get-started)