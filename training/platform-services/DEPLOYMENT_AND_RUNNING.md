Deployment and Running
===========================

## Concourse

Visit <https://concourse.dp-ci.aws.onsdigital.uk/> to access Concourse

[<img src="images/concourse-home.png" title="Concourse Homepage" alt="concourse ons" width=49.7% height=49.7%/>][concourse-ons] [<img src="images/concourse-pipeline.png" title="Concourse Pipeline" alt="concourse pipeline" width=49.7% height=49.7%/>][concourse-ons]

Concourse lists all of the applications where each application has its own [pipeline][concourse-pipeline].
When a [PR][git-pr] of an application is said to be merged into `develop` or `master`/`main`, Concourse runs the pipeline of the application which does unit testing, checks if a [docker image][docker-image] can be built for it, bundles it for release and ships it to that release whilst ensuring that it has deployed everything completely.

Furthermore, Concourse helps us to identify if an application has been fully deployed with any new changes implemented and can provide information on the version of the application which is currently running in the `develop` and `production` environments.

Please note that when deploying changes to the `production` environment (the live website), you need to **manually** deploy the application by doing the following:

1. Go to [Concourse][concourse-ons]
2. Login to concourse as `dp` user (if you do not have the login, please ask anyone in the team)
3. Navigate to the application
4. Click on `production-ship-it`
5. <div>Click on the <img src="images/concourse-build.png" title="Concourse Build" width=4% height=4%></div>

More information about Concourse can be found on their website - <https://concourse-ci.org/docs.html>

## Hashistack

The next following 3 services: [Consul](#consul), [Nomad](#nomad) and [Vault](#vault) are all part of [HashiCorp][hashicorp-home] which are sometimes referred as the Hashistack

## Consul

Visit the following to access Consul

- SANDBOX:&nbsp;&nbsp;<https://consul.dp.aws.onsdigital.uk/>
- STAGING:&nbsp;&nbsp;<https://consul.dp-staging.aws.onsdigital.uk/>
- PROD:&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<https://consul.dp-prod.aws.onsdigital.uk>

[<img src="images/consul-dev-services.png" alt="consul develop home" title="Consul Develop Home"/>][consul-dev-services]

> Consul is a service mesh solution providing a full featured control plane with service discovery, configuration, and segmentation functionality. Each of these features can be used individually as needed, or they can be used together to build a full service mesh

Consul is mainly used to check the health of applications which Consul retrieves by calling the health check endpoints of the application (`/health`). This information is then used to monitor cluster health, and it is used by the service discovery components to route traffic away from unhealthy hosts

All apps that run in production are required to have a health check. The health check must have an HTTP health endpoint to return the health status, even if they are event-driven and comply with the health check specification which can found in [dp/standards/HEALTH_CHECK_SPECIFICATION.md][health-check-spec]

More information about Consul can be found on their website - <https://www.consul.io/docs/intro>

## Nomad

Visit the following to access Nomad

- SANDBOX:&nbsp;&nbsp;<https://nomad.dp.aws.onsdigital.uk/>
- STAGING:&nbsp;&nbsp;<https://nomad.dp-staging.aws.onsdigital.uk/>
- PROD:&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<https://nomad.dp-prod.aws.onsdigital.uk>

[<img src="images/nomad-dev-home.png" title="Nomad Develop Home" width=49.7% height=49.7%/>][nomad-dev-home] [<img src="images/nomad-dev-job.png" title="Nomad Develop Job Page" width=49.7% height=49.7%/>][nomad-dev-home]

> Nomad is a simple and flexible workload orchestrator to deploy and manage containers and non-containerized applications across on-prem and clouds at scale

Nomad enables us to deploy our applications by bundling computing resources of [AWS][aws]. In Nomad, each application is considered as a [job][nomad-job] which can have multiple [task groups][nomad-task-group] (`web` or `publishing`). The mapping of these tasks in a job is done using allocations, which is used to declare that a set of tasks in a job should be run on a particular node within the [AWS Auto-scaling Group (ASG)][asg]. In more detail, the nomad server talks to its nomad clients, which each sit on a node in an ASG, and the allocations are assigned to a nomad client, which then manages the allocation for its lifecycle on that node. Nomad would schedule these allocations which is detailed more [here][nomad-scheduling].

With Nomad, we can understand whether a particular allocation or the entire service itself is failing. We can look at a specific task depending on the journey (e.g. looking at the publishing task group of the application if the publishing journey was carried out). We can check the logs for specific allocations for debugging. Please note that the logs that we can see in Nomad are limted, hence we can use [Kibana](MONITORING_AND_ALERTING.md#kibana) to get more information.

Another common use of Nomad is to stop and start a service (do not do this in production) to capture new updates and secrets for the application. On the other hand, we can tweak the nomad plan's CPU or memory in any environment to trigger a rolling re-deploy which will pick up new secrets which is a safer alternative method to carry out in the production environment.

More information about Nomad can be found on their website - <https://www.nomadproject.io/>

## Vault

Visit the following to access Vault

- SANDBOX:&nbsp;&nbsp;<https://vault.dp.aws.onsdigital.uk/>
- STAGING:&nbsp;&nbsp;<https://vault.dp-staging.aws.onsdigital.uk/>
- PROD:&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<https://vault.dp-prod.aws.onsdigital.uk>

**To sign into Vault, you require a `token` which can be retrieved from the [helper script in dp-setup][helper-script-dp-setup]**

[<img src="images/vault-dev-login.png" title="Vault Develop Login" width=49.7% height=49.7%/>][vault-develop] [<img src="images/vault-dev-home.png" title="Vault Develop Home" width=49.7% height=49.7%/>][vault-develop]

> Vault secures, stores and tightly controls access to tokens, passwords, certificates, encryption keys for protecting secrets and other sensitive data

Vault manages secrets and stores the secrets in [dp-configs][dp-configs-secrets] in a secure way and makes them available to the service. The secrets can include information such as how to communicate with other services and database, or feature flags controlling what behaviour is active. In general, Vault is used for the following two functionalities:

1. Storing application configs (secrets) from [dp-configs][dp-configs-secrets]
2. Storing per-file encryption keys for florence uploads.

Some apps that deal with S3 uploads (florence) need a vault client in the app, (hence additional dp-vault related config is informed in the app config), while other apps only retrieve configs. Please note that all apps need their [vault policies][vault-policy] defined in [dp-setup/ansible/files/vault-policies][dp-setup-vault-policies]. Policies provide a declarative way to grant or forbid access to certain paths and operations in Vault. That’s why if you look at the vault policies, most will only have one path while some (such as [florence vault policy][florence-vault-policy] which deals with S3 uploads) will have multiple.

More information about Vault can be found on their website - <https://www.vaultproject.io/>

[//]: # (Reference Links and Images)
[asg]: <https://docs.aws.amazon.com/autoscaling/ec2/userguide/AutoScalingGroup.html>
[aws]: <https://aws.amazon.com/>
[concourse-ons]: <https://concourse.dp-ci.aws.onsdigital.uk>
[concourse-pipeline]: <https://concourse-ci.org/pipelines.html#pipelines>
[consul-dev-services]: <https://consul.dp.aws.onsdigital.uk/ui/eu/services>
[docker-image]: <https://jfrog.com/knowledge-base/a-beginners-guide-to-understanding-and-building-docker-images/>
[dp-configs-secrets]: <https://github.com/ONSdigital/dp-configs/tree/master/secrets>
[dp-setup-vault-policies]: <https://github.com/ONSdigital/dp-setup/tree/awsb/ansible/files/vault-policies>
[florence-vault-policy]: <https://github.com/ONSdigital/dp-setup/blob/awsb/ansible/files/vault-policies/florence.hcl>
[git-pr]: <https://docs.github.com/en/free-pro-team@latest/github/collaborating-with-issues-and-pull-requests/about-pull-requests#:~:text=Pull%20requests%20let%20you%20tell,merged%20into%20the%20base%20branch.>
[hashicorp-home]: <https://www.hashicorp.com>
[health-check-spec]: <https://github.com/ONSdigital/dp-standards/blob/main/HEALTH_CHECK_SPECIFICATION.md#health-check-specification>
[helper-script-dp-setup]: <https://github.com/ONSdigital/dp-setup/tree/develop/scripts>
[nomad-dev-home]: <https://nomad.dp.aws.onsdigital.uk>
[nomad-job]: <https://learn.hashicorp.com/tutorials/nomad/get-started-jobs>
[nomad-scheduling]: <https://www.nomadproject.io/docs/internals/scheduling/scheduling>
[nomad-task-group]: <https://www.nomadproject.io/docs/job-specification>
[vault-develop]: <https://vault.dp.aws.onsdigital.uk/>
[vault-policy]: <https://www.vaultproject.io/docs/concepts/policies>
