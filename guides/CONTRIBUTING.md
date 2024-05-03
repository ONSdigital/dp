# Contributing to Dissemination repos

## Development work

* As a first step, GPG key setup is required in order to contribute:
  instructions on setup can be found in the [GPG Guide](https://github.com/ONSdigital/dp-operations/blob/main/guides/gpg.md)
* We use [git-flow](https://www.atlassian.com/git/tutorials/comparing-workflows/gitflow-workflow) -
  create a feature or fix branch from `develop`, e.g. `feature/<new-feature>`, `fix/<new-fix>`. Alternatively, create a hotfix branch from master e.g. `hot-fix/<new-hotfix>`.
* Pull requests must use the pull request template found in each repo. Each section must contain the following:
  * 'What' - a succinct, clear summary of what the user-need is that is driving this feature change,
  * 'How to review' - a step-by-step guide or set of guides that detail how a reviewer can test and verify the changes made
  * 'Who can review' - a list of the most appropriate people to review the pull request,
    e.g. if on a frontend application, a member of the frontend team should review; likewise for backend and platform.
    If the changes are critical or are likely to have severe side-effects then the Tech Lead should review.
* Ensure your branch contains logical atomic commits before sending a pull request - follow the [alphagov Git styleguide](https://github.com/alphagov/styleguides/blob/master/git.md)
* You may rebase your branch after feedback if it's to include relevant updates from the `develop` branch. We prefer a rebase here to a merge commit as we
  prefer a clean and straight history on the `develop` branch, with discrete merge commits for features
* It is advised to squash commits before pushing to a remote repository

----

## Releasing

* For instructions on how to release work, read the [Releases guidelines](RELEASES.md)

----

## Access to Environments

### Note: if you get a 404 error for any of the below links, then you need to be added to the `ONSdigital` organisation in GitHub

* For direct access to the environments
  * Setup your [AWS credentials](AWS_CREDENTIALS.md)
  * Configure your [SSH access](https://github.com/ONSdigital/dp-operations/blob/main/guides/ssh-access.md)
  * [Ansible](https://github.com/ONSdigital/dp-ci/tree/master/ansible#prerequisites) is required for provisioning to environments
