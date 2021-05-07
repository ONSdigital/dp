Architecture Principles
=================
These principles are an evolution of the [GDS Architecture approach](https://docs.google.com/presentation/d/1jj8mndjx0ZYXo0ZS9V9AZRlh0-enuZVYF1L44Myj80Q/edit#slide=id.g1155300453_2_10)

The system exists to serve users
---------
* Simplicity for users is worth some complexity in the system
* Users need available and resilient services
* User experience should degrade gracefully

Be good citizens of the web
---------
* Don’t break the internet (or peoples bookmarks)
* All changes should be transparent
* Breaking changes must be versioned
* Make use of open standards ([as outlined by GDS](https://www.gov.uk/government/publications/open-standards-principles/open-standards-principles#why-you-should-use-open-standards))

Open source by default
---------
* Don’t use proprietary formats
* Open by default - APIs included
* Avoid vendor lock in where possible, but not at the expense of the operational burden of not using managed services

We are not unique
---------
* Don’t reinvent the wheel
* Consume 3rd party APIs for data that isn’t ours to own

Innovate Responsibly™
---------
* Start small and question scope
* One step at a time (iteration)
* Security from the start
* [Eat our own dog food](https://en.wikipedia.org/wiki/Eating_your_own_dog_food) - consume our own APIs

Single responsibilities - the systems and the teams
---------
* [Loosely coupled microservices](https://microservices.io/microservices/2021/01/04/loosely-coupled-services.html)
* All apps should be aligned to [12 Factor principles](../training/architecture/12_FACTOR_APP_PRINCIPLES.md)
* An app must have sole ownership of its data
* There must be a single source of truth for any data

Design for operability
---------
* [A little duplication is better than a little dependency](https://www.youtube.com/watch?v=PAAkCSZUG1c&t=568s)
* [Clear is better than clever](https://dave.cheney.net/2019/07/09/clear-is-better-than-clever)
* A new feature is worthless if it’s not supportable

Repeatable for reliability
---------
* Choose the right kind of tests for the job
* Nothing should exist that can’t be rebuilt simply
