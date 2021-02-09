12 Factor App Principles
===========================

In Digital Publishing we build, deploy and manage applications which run as services. The twelve-factor methodology helps us follow best practices and build applications which are portable and resilient. This module will introduce you to the Twelve-factor principles and provide you with examples of how we use them in our team. 

Note: The examples refer primarily to Go apps because that's what I know about. Please add more details or alternative examples which might be useful.

## Pre reading

- [About version control](https://git-scm.com/book/en/v2/Getting-Started-About-Version-Control)

## Pre requisites


## Materials

### 1. Codebase
We create one GitHub repository per application, for example the [dp-bulletin-api](https://github.com/ONSdigital/dp-bulletin-api) repository contains our codebase for that app. 
Whenever we _deploy_ the dp-bulletin-api, we use a version of the code in the dp-bulletin-api repository. The `production` environment will contain a version deployed from the `master` or `main` branch of the repository, which may be different from the `develop` branch.

Find out more about how we [collaborate on Digital Publishing repos](https://github.com/ONSdigital/dp/blob/master/guides/CONTRIBUTING.md).
### 2. Dependencies
- Apps are containerised
- Use languages that let us keep track of dependencies, e.g. the [go.mod file](https://blog.golang.org/using-go-modules) in Go.
### 3. Config
We use the [envconfig package](https://github.com/kelseyhightower/envconfig) to use environment variables for configuration settings. In [this example](https://github.com/ONSdigital/dp-bulletin-api/blob/develop/config/config.go), the `config/config.go` contains a default value for `BindAddr` of `":24200"`; to overwrite it, you can set the environment variable `BIND_ADDR` to a different value.

### 4. Backing services
### 5. Build, release, run
### 6. Processes
### 7. Port binding
### 8. Concurrency
### 9. Disposability
### 10. Dev/prod parity
### 11. Logs
### 12. Admin processes

## Next steps


Further resources
----------------------------
- Using [dp-cli to generate a repository](https://github.com/ONSdigital/dp-cli/blob/master/project_generation/COMPLETE_PROJECT_SETUP.md)
- [Pull Request guidance](https://github.com/ONSdigital/dp/blob/master/training/culture-and-process/PULL_REQUEST_GUIDANCE.md)
- [envconfig](https://github.com/kelseyhightower/envconfig) package
