# Renaming an App

When renaming a service from A to B, we want to keep A’s configuration available until we are sure that B is working.
Once B is up, we can tidy up and remove references to A.

Note that some event driven apps, will not have 'magic ports' so the routing and 'ports' changes aren't applicable as
they would be with APIs.

## Move OLD_SERVICE_NAME to NEW_SERVICE_NAME (DEVELOP)

- Issue PR to dp-setup:

  - Add a new policy to vault-policies and dp-deployer policy

  - If applicable, add routing templates to dp-setup, adding new stanzas for the new name. Leave the old ones in place
    for now (backwards compatible). The port number should be a new one (+1 from the previous).

    (The files which need changes can be found in the guide for
    [Adding a New App](https://github.com/ONSdigital/dp-setup#adding-a-new-app) )

  - If applicable, add the new port to PORTS.md.

- Issue PR to dp:

  - Update the OLD_SERVICE_NAME to NEW_SERVICE_NAME associated with the port in the PORTS.md in dp

- Change branch protection in GitHub for develop branch so that the checks for the OLD_SERVICE_NAME are not required (
  this is achieved by going to settings of the repo in GitHub)

- Issue PR to service being renamed:

  - Update the log.Namespace in the service

  - Update any other occurrences.

- Rename the repository in GitHub - be sure to `@developers` in `#development-general` to let them know when this is
  happening, and remind them they can ensure their local clones are pointing to the right repo by using
  `git remote set-url origin <new_url>`

- Issue PR to dp-configs:

  - Create a copy of the old secrets files in dp-configs for develop and name them with the NEW_SERVICE_NAME.

  - Update URL route of service in the secrets.

  - Update the copied [manifest file](https://github.com/ONSdigital/dp-configs/tree/master/manifests) to reflect the
    name change and repo move.

  ⚠️ Ensure all changes have been merged into develop before continuing.

- Apply the dp-setup ansible for DEVELOP for the vault policies ONLY (step 4 in
  [Adding a New App](https://github.com/ONSdigital/dp-setup#adding-a-new-app))

- If the service has a port (eg. api), search for the service URL (e.g. "SEARCH_URL") in the organization in GitHub to
  see where the service is being used and update those accordingly with NEW_SERVICE_NAME.

  ⚠️ Make a note of the PRs (changes) as these need to be released later.

- Issue PR to dp-configs:

  - Update develop secrets of services which use OLD_SERVICE_NAME with the new port numbering

- If applicable, run terraform apply in [dp-ci](https://github.com/ONSdigital/dp-ci/tree/master/terraform)

- Create a new pipeline in concourse for the newly named service (step 4 in [Manually Generating the Pipelines
  ](https://github.com/ONSdigital/dp-ci/tree/master/pipelines/pipeline-generator#manually-generating-the-pipelines))

- Ship the service in develop through concourse

- Update api-router config to point to new port (if the service is a backend one, e.g. API)

- Update frontend-router config to point to new port (if the service is a frontend one, e.g. controller)

- Apply the dp-setup ansible for DEVELOP for the consul-template (step 4 in
  [Adding a New App](https://github.com/ONSdigital/dp-setup#adding-a-new-app))

- Check if develop environment is still working as expected. Then PO-sign off this with Tech Lead

## Move OLD_SERVICE_NAME to NEW_SERVICE_NAME (PRODUCTION)

- Release the changes made in dp-setup

- Release the changes made in NEW_SERVICE_NAME

- Rename secrets files in dp-configs for production AND update URL route of service in the secrets

- Change branch protection in GitHub for master branch for NEW_SERVICE_NAME (this is achieved by going to settings of
  the repo in GitHib)

- Release the changes made in dp-api-router or dp-frontend-router

- Apply the dp-setup ansible for PROD for the vault policies ONLY (step 4 in
  [Adding a New App](https://github.com/ONSdigital/dp-setup#adding-a-new-app))

- Ship the service in production through concourse

- Apply the dp-setup ansible for PROD for the consul-template (step 4 in
  [Adding a New App](https://github.com/ONSdigital/dp-setup#adding-a-new-app))

## Tidy up - DEVELOP

- Remove dp-setup/ansible/files/vault-policies of OLD_SERVICE_NAME

- Remove OLD_SERVICE_NAME from dp-setup/ansible/files/vault-policies/dp-deployer.hcl

- Remove OLD_SERVICE_NAME from dp-setup/ansible/templates/consul-template (web-nginx.http.conf.tpl.j2 for web) and (
  publishing-nginx.http.conf.tpl.j2 for publishing)

- Remove OLD_SERVICE_NAME from dp-setup/splash-page/app.js and add NEW_SERVICE_NAME

- Apply dp-setup ansible for the consul template and vault policies in develop

  STEP 4 - https://github.com/ONSdigital/dp-setup#adding-a-new-app

- Check if develop environment is still working as expected

- Issue PR to dp-configs develop secrets: remove OLD_SERVICE_NAME secrets

- Delete secrets from vault web interface develop itself (Go to the secrets and click on delete)

- Stop the OLD_SERVICE_NAME running in nomad develop (it will be automatically removed by the garbage collector in the
  next few days)

- Check if develop environment is still working and then PO-sign off this with Tech Lead

## Tidy Up - PRODUCTION

- Release changes of dp-setup and check if production environment is still working as expected

- Delete secrets from vault web interface production itself (Go to the secrets and click on delete)

- Stop the OLD_SERVICE_NAME running in nomad production (it will be automatically removed by the garbage collector in
  the next few days)

- Update concourse to not include OLD_SERVICE_NAME (do pipeline changes
  https://github.com/ONSdigital/dp-ci/tree/master/pipelines/branch-manager)

- Update references to OLD_SERVICE_NAME to NEW_SERVICE_NAME in [the dp repo](https://github.com/ONSdigital/dp)

- Remove AWS ECR repository of the OLD_SERVICE_NAME (ask Tech Lead to remove as they should have the permission to do
  so)

## Post-Rename Tasks

After the service itself is renamed, there may be additional steps required depending on the app type.

### Kafka

Once the manifest file has been renamed, an event-driven (kafka-using) app will require:

- a new [client certificate](https://github.com/ONSdigital/dp-setup/tree/develop/csr/private#creating-a-client-certificate)
  (for the new name)
- kafka will need to
  [be updated using the topic-manager](https://github.com/ONSdigital/dp-setup/tree/develop/scripts/kafka#topic-manager)
  to allow the new certificate access to its (possibly changed) topics
- if a consumer, the name of the consumer group will need to be changed (see below)
- if a producer, the topic name may need to be renamed (this is less likely, but see below)

#### Rename the kafka consumer group

The convention is for the kafka consumer group to be the same as the app name so this will need
changing. It is important to note that this may cause messages to be replayed so this step can only be done if the
handler is idempotent.

- Update the secrets in dp-configs
- Update the default consumer group config in the app itself along with any references eg. in the README.
- Release to develop and watch for any spike in traffic or any problems due to lack of idempotency.
- If develop was successful, release the app to production.

#### Rename the kafka producer topic

Depending on the nature of the name change, one or more kafka topics may need to be renamed. Renaming these needs to be
done in a specific order so that messages aren't lost. To do this, release the producers before releasing the consumers.

##### In develop

- Rename the kafka topic in the manifest and apply it to develop
  [using the topic-manager](https://github.com/ONSdigital/dp-setup/tree/develop/scripts/kafka#topic-manager)

  This creates a new topic in kafka with the new name, but it does not delete the old topic.

- Rename the secrets in dp-configs for develop for both the producer app and any consumers.

- Rename the default config (eg. in `config.go`) in the **producer** app and merge the PR.

  ⚠️ Ensure that the topic-manager step is complete and secrets have been synced first or this step will fail.

- Rename the default config (eg. in `config.go`) in all **consumer** apps and merge the PRs.

  ⚠️ Ensure that the producer app step is complete first or this step will fail.

- Check that messages are processing correctly on develop.

##### In production

- Repeat the dp-setup process above in production

- Rename the secrets in dp-configs for production for both the producer app and any consumers.

- Release the **producer** app to production

  ⚠️ Ensure that the topic-manager step is complete and secrets have been synced first or this step will fail.

- Release all **consumer** apps to production.

  ⚠️ Ensure that the producer app step is complete first or this step will fail.

- Check that messages are processing correctly on production.
