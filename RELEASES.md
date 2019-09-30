Release process
===============

#### Notes

* Some ops work does not need to go through `PO sign off` or `Ready for release` - for example, terraform configuration which has already been applied to **develop** or **production** environments.
* Major features which are being tested in sandpit do not join this process until they are merged to `develop`/`cmd-develop`

## Development 'done'

* Pull request from feature branch into `develop`
* Code is peer-reviewed and has passed in CI
* Code issues identified in pull request are fixed
* Pull request approved
* Feature branch merged into `develop`
* Developer tests feature in `develop` environment
* Developer moves story from team board to release board `PO sign off` column

## 'PO sign off' column

* Product owner reviews changes in `develop` environment
* Bug fixes applied to feature branch and retested
* Once signed off, product owner moves story to `Ready for release` column

## 'Ready for release' column

* Release branch (e.g. `release/1.7.3` or `beta/1.0.0`) taken from `develop`
* Bug fixes applied to release branch
* Release branch bug fixes regularly merged back to `develop`
* Pull request from release branch into `master`
* [Release tag created](TAGS.md) from `master` and release branch deleted

## Deployment

* `master` needs a manual deployment to **production** from the release tag this can be done in CI
* To access CI for deployments ask a member of the Dev team.
