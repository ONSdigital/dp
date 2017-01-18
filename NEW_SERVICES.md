Creating new services
=====================

If you're creating a new service, these are some things you'll need to do:

- Create a new GitHub repo
  - use [dp-repo-template](https://github.com/ONSdigital/dp-repo-template) - see [development standards](DEV_STANDARDS.md)
  - set the default branch to `develop`
  - add the `Digital Publishing` and `Data Discovery` teams as collaborators
  - protect both the `master` and `develop` branches, with these settings:
    - Protect this branch
    - Require pull request reviews before merging
    - Include administrators
    - Restrict who can push to this branch:
      - `ONSdigital/data-discovery`
      - `ONSdigital/digitalpublishing`
  - setup the GitHub webhooks - see https://github.com/ONSdigital/dp-ci
- Setup the project for CI
  - Add Jenkinsfile to the project repo
  - Add CodeDeploy scripts to the project repo
  - Add Dockerfile to the project repo
  - Add a new ECR repo to dp-ci terraform
  - Add configs to the dp-ci repo
  - Add the project repo to the Jenkins ansible script
