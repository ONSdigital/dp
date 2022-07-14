# Release process

## Notes

* **PR** refers to a [Github Pull Request](../training/culture-and-process/PULL_REQUEST_GUIDANCE.md)
* **Story** refers to the Trello ticket for the work. **Columns** are where stories move within Trello.
* In some repos, the `main` branch may still have the **deprecated branch name**: `master` (replace in the below, as necessary)
* Some **ops work** does not need to go through `PO sign off` (though a tech lead may sign it off) or `Ready for release`
  * for example, terraform configuration which has already been applied to the *sandbox* or *production* environments.

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
    * in [_consul_](https://consul.dp.aws.onsdigital.uk/ui/eu/services) (for the expected commit/version and health)
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
  * the release PR must include the release notes e.g. the features, bugs that have been made
  * any issues (e.g. in CI) are applied to release branch
  * release branch fixes regularly merged back to `develop`
* **Merge** into `main` branch
* [Create Release tag](TAGS.md) from `main` branch
  * tag must be published as a release, with bullet point list of changes in the release notes. These should be obtained from the release pull request description
  * release branch can be deleted

## Deployment

* The `main` branch requires **manual deployment** to the *production* environment - this can be done in CI
  * to access `production-ship-it`, in CI, ask a member of the dev team
  * ensure the app has been shipped as expected
    * in CI (`production-ship-it` job was successful for the expected release)
    * in _consul_ (for the expected version and health)
* Any issues arising:
  * major issues should prompt you to rollback to the previous version and re-work the original (or a new) feature branch; see [dp operation docs on how to rollback an application to previous version](https://github.com/ONSdigital/dp-operations/blob/main/guides/rollback-app.md#rolling-back-an-application-to-a-previous-deployment)
  * minor issues are fixed in `hotfix/my_fix` branches (which are PR'd back into the `main` branch)
  * merge any hotfixes back into the `develop` branch, too
* Developer moves the story to the `Done` column :tada:
