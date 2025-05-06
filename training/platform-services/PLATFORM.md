Platform Services
===========================

Platform services allow us to develop, run and manage applications without the complexity of building and maintaining the infrastructure typically associated with developing and launching an app. This documentation gives information on the platform services that we use to assist in deploying and monitoring applications.

## Prerequisites

You need to complete the [set up your Mac for development work](/guides/MAC_SETUP.md). If this has not been completed, please click on the link and follow the instructions. Afterwards, return back to this page.

Run `dp remote allow sandbox` and `dp remote allow ci` to ensure you can access concourse and the sandbox environment.
Please note that if you have prod access, you can run `dp remote allow prod` to access the prod environment.

[Cloud Technologies and Infrastructure](../services/INFRASTRUCTURE.md) - You should have an understanding of AWS and the infrastructure as the platform services are very closely related to this.

## Structure
![platform structure](images/platform-structure.png)

## Contents

- Provisioning: [Terraform](PROVISIONING.md#terraform) & [Ansible](PROVISIONING.md#ansible)
- Deployment and Running: [Concourse](DEPLOYMENT_AND_RUNNING.md#concourse), [Consul](DEPLOYMENT_AND_RUNNING.md#consul), [Nomad](DEPLOYMENT_AND_RUNNING.md#nomad) & [Vault](DEPLOYMENT_AND_RUNNING.md#vault)
- Monitoring and Alerting: [Kibana](MONITORING_AND_ALERTING.md#kibana), [Grafana](MONITORING_AND_ALERTING.md#grafana) & [Prometheus](MONITORING_AND_ALERTING.md#prometheus)

## Next steps

1. Gain experience and familiarity with the platform services
2. Please update this file if important information is missing regarding the platform services

[//]: # (Reference Links and Images)
   [setup-ons-mac]: <https://github.com/ONSdigital/dp/blob/main/guides/GETTING_STARTED.md#set-up-your-mac-for-development-work>
