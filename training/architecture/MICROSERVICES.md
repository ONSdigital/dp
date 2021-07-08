Microservice architecture
===========================

We use a microservice architecture in our approach to new system components as it provides a flexible, iterative and scalable way to solve new problems. Its worth understanding some basic theory around what a microservice architecture means, but also to get familiar with the service classes and terms we use when discussing these concepts.

### Pre-reading

To understand some of our pre-existing structures, and some of the problems we're trying to move past, it may be helpful to review our [core services](../services/CORE_APPS.md) training materials.


### Prerequisites

Understand a [basic definition of microservices](https://smartbear.com/solutions/microservices/)

### Materials

[This presentation](https://docs.google.com/presentation/d/1bXZw2rJTBa7spXVQNuA0lVQSN-2jZiz9khn6hJsvWwc/edit#slide=id.ge2b2dc2785_0_359) walks through our different service types and where they are commonly deployed. It also highlights some considerations and standards that are common across all of our service types, such as healthchecking.

As microservices mean we often have to create new repositories, we have created a [project generation](https://github.com/ONSdigital/dp-cli/tree/main/project_generation) tool that allows us to start from a boilerplated version of the relevant kind of repository. We also have a [hello world](https://github.com/ONSdigital/dp-hello-world) repo where we document our baseline application structure for reference, review, and as a place to center conversations about standards.

### Next steps

[12 factor app](12_FACTOR_APP_PRINCIPLES.md) principles are a common set of considerations for building applications that can be deployed in a clean containerised way - and is very common in a microservice paradigm.

[API](API.md) training will explain in more detail what our API microservices are responsible for. It will also link to relevant standards around API design and development.


Further resources
----------------------------

* [Microservices.io](https://microservices.io/patterns/microservices.html)