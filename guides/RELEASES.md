# Release process

This is an overview of the sequence of steps that our tickets (sometimes called *tasks* or *stories*)
take from initial idea to completion.

## Notes

* **PR** refers to a [Github Pull Request](../training/culture-and-process/PULL_REQUEST_GUIDANCE.md)
  * not to be confused with the related *Peer Review* (when someone assesses a Pull Request)
* **Story** refers to the Jira ticket describing the work.
* Some **ops work** does not need to go through `PO sign off` (though a tech lead may sign it off) or `Ready for release`
  * for example, terraform configuration which has already been applied to the *sandbox* or *production* environments.
* **Columns** are where stories move within Jira - from left-to-right - towards the 'Done' column.
* The current environments and their base branches (from which deployments are made) are:

   |                                   | Sandbox env         | Staging env          | Prod env
   | - | - | - | -
   | env used for                      | initial integration | final stable testing | live production env
   | applications (base branch)        | `develop`           | `master`             | `master`
   | infrastructure `dp-setup` branch  | `awsb`              | `main`               | `main`
   | `dp-configs` (secrets, manifests) | `main` branch       | `main`               | `main`

* In some repos (apps, largely), the `main` branch may still have the **deprecated branch name** of `master` (replace in the below, as necessary)
* Some **ops/infrastructure work** (e.g. in `dp-setup`) does not need to go through `PO sign off` (though your Tech Lead should sign it off instead of the Product Owner)
  * for example, terraform configuration which has already been applied to the current environments
  * if in doubt, ask your Tech Lead, or the Platform Team

## 'Backlog' column

Tickets often start in this column, but may not be fully-formed and need clarification/prioritisation.

Once a ticket has been given enough detail for someone to be able to start working on it, it can move to the 'Ready' column.
More information may be required, but it is sufficiently detailed for most.

This move usually happens in planning sessions.

## 'Ready' column

This is the list of stories that are aimed to be completed in the current sprint.
The column is often sorted by priority (those at the top have higher priority).

Scrum teams may break this column into sections for the current (higher section) and future (lower sections) tickets.

During the sprint, team members take tickets from this column, add themselves to that ticket,
move the ticket to the 'In progress' column and begin the work therein.

## 'In progress' column

Once work/coding is complete:

* **Create a PR** - to have your feature branch merged into the `develop` branch
* Ensure the PR has passed in CI
* Request a peer-review on Slack (`#dev-code-review`)
* Move the ticket to the 'PR' column

## 'PR' column

* **Get approval:** get the PR peer-reviewed
  * if any code or other issues are identified, apply fixes to the feature branch
* **Merge** the approved feature branch into the `develop` branch
  * do **not** merge using Github :warning:
* **Ship it** - the updated `develop` branch should be auto-deployed to the *sandbox* environment
  * ensure that shipping was successful:
    * in [CI](https://concourse.dp-ci.aws.onsdigital.uk/) (`sandbox-ship-it` job was successful for the right commit)
    * in [*consul*](https://consul.dp.aws.onsdigital.uk/ui/eu/services) (for the expected commit/version and health)
* Developer moves the story to `Testing` column

## 'Testing' column

* Developer **tests** the feature in the *sandbox* environment
  * similar to approval, above: any fixes go into your feature branch for re-approval, re-merge, etc
    * if your feature branch was deleted (manually or automated by github settings), you can always reopen it via the github UI
* Developer moves the story to `PO sign off` (if required) or `Ready for release` column

## 'PO sign off' column

* Product owner **reviews changes** in the *sandbox* environment
  * bug fixes applied to feature branch (and re-approved, etc into the `develop` branch)
* **Sign off:** the product owner moves the story to `Ready for release` column

## 'Ready for release' column

* **Create Release branch** (e.g. `release/1.7.0` see [version control](VERSIONING.md) ) taken from `develop` branch
* **Create a PR** - to merge the release branch into the `main` branch
  * the release PR must include the release notes
    * a list of the features and bug fixes added since the last release
    * please ask the other contributors (commit authors) to add to the release notes for their commits/features/fixes
  * any issues (e.g. in CI) are applied to release branch
  * any release branch updates (from PR comments) should be regularly merged back to `develop` (or `awsb` for `dp-setup` repo)
* **Merge** into `main` branch but don't push it to Github (origin) yet, wait until the tag has been created in the next step and push both at the same time.
* [Create Release tag](TAGS.md) from `main` branch
  * tag must be published as a release, with bullet point list of changes in the release notes. These should be obtained from the release pull request description
  * release branch can be deleted

## Deployment

* The `main` branch requires **manual deployment** to the *production* and *staging* environments
  * for apps, this can be done in CI
    * to access `production-ship-it` or `staging-ship-it`, in CI, ask a member of the dev team
    * ensure the app has been shipped as expected
      * in CI (`production-ship-it` and `staging-ship-it` jobs were successful for the expected release)
    * in *consul* (for the expected version and health)
  * for dp-setup (infrastructure) changes, please ask each commit author (in the release) to apply their commits/changes to the appropriate environments
* Any issues arising:
  * major issues should prompt you to rollback to the previous version and re-work the original (or a new) feature branch; see [dp operation docs on how to rollback an application to previous version](https://github.com/ONSdigital/dp-operations/blob/main/guides/rollback-app.md#rolling-back-an-application-to-a-previous-deployment)
  * minor issues are fixed in `hotfix/my_fix` branches (which are PR'd back into the `main` branch)
  * merge any hotfixes back into the `develop` branch, too
* Developer moves the story to the `Done` column :tada:
