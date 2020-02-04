Release process
===============

#### Notes

* Some ops work does not need to go through `PO sign off` or `Ready for release` - for example, terraform configuration which has already been applied to **develop** or **production** environments.

## Development 'done'

* Pull request from feature branch into `develop`
* Code is peer-reviewed and has passed in CI
* Code issues identified in pull request are fixed
* Pull request approved
* Feature branch merged into `develop`
* Developer tests feature in `develop` environment
* Developer moves story to `PO sign off` column

## 'PO sign off' column

* Product owner reviews changes in `develop` environment
* Bug fixes applied to feature branch and retested
* Once signed off, product owner moves story to `Ready for release` column

## 'Ready for release' column

* Release branch (e.g. `release/1.7.0` see [version control][VERSIONING.md]) taken from `develop`
* Bug fixes applied to release branch
* Release branch bug fixes regularly merged back to `develop`
* Pull request from release branch into `master`
* [Release tag created](TAGS.md) from `master` and release branch deleted

## Deployment

* `master` needs a manual deployment to **production** from the release tag this can be done in CI
* To access CI for deployments ask a member of the Dev team.
