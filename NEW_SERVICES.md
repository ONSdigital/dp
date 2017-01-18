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
  - Add required files to the project repo
    - Jenkinsfile
    - CodeDeploy scripts
    - Dockerfile
  - Add required entries to the [dp-ci](https://github.com/ONSdigital/dp-ci) repo
    - Add a new ECR repo to terraform
    - Add application config files
    - Add to the Jenkins ansible script
