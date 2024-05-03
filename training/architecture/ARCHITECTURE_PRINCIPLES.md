Architecture Principles
===========================

This training document breaks down our architectural standards and explains them in a bit more detail - providing context for how and why we focus on certain approaches.

## Pre-reading
Read our [Architecture Principles](https://github.com/ONSdigital/dp-standards/blob/main/ARCHITECTURE_STANDARDS.md#architecture-principles) and some of the referred resources to understand our high level approach

## Materials

Our approach to all things in Dissemination should find its roots in the [GDS Service Manual](https://www.gov.uk/service-manual) - particularly for architecture, this means the advice given in the [Technology](https://www.gov.uk/service-manual/technology) section. We don't expect everyone to read the whole service manual, but it can be a helpful resource when exploring new areas.

We are part of a community of open data publishers, so we have to ensure systems we build not only solve our problems but utilize common standards to help build consistent views of data across the web. [We've blogged](https://digitalblog.ons.gov.uk/2017/01/06/some-open-data-publishing-principles/) about this in the past, and it's worth understanding the [basics of open data](https://opendatahandbook.org/) as it often informs our approach.

We are responsible for our entire website estate - not just applications, but the databases and infrastructure that allows those applications to operate. This means our architectural considerations often involve deciding at what part of the stack to place a burden, and making sure that all parts of our stack are as robust as each other. We take an [Infrastructure as Code](https://www.hashicorp.com/resources/what-is-infrastructure-as-code) approach, ensuring that all of our infrastructure can be built through a few simple commands, and key values are configurable.


## Next steps

Learn more about how we put some of these princples into action through additional modules:
- [12 Factor App Principles](12_FACTOR_APP_PRINCIPLES.md)
- [Microservices](MICROSERVICES.md)


Further resources
----------------------------
- [Open Data Institute](https://theodi.org/)
- [Code of Practice for Statistics](https://code.statisticsauthority.gov.uk/)
- [GDS - Open source code](https://www.gov.uk/service-manual/service-standard/point-12-make-new-source-code-open)
- [GDS - Use open standards](https://www.gov.uk/service-manual/service-standard/point-13-use-common-standards-components-patterns)
