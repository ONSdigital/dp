# Tagging a release

## Pre-reading

- [Understanding semantic versioning](../guides/VERSIONING.md) and how Dissemination implement this on our git repositories.

## How to tag a release/commit in git

In order to tag a release you first need to checkout the branch you aim to tag (usually `main` or `master`)

* Run `git tag -s <tag> -m <message>`
	* `<tag>` is the semver value in the release branch name (e.g. `v1.7.0`, see [version control](VERSIONING.md))
	* `-s` ensures the tag is signed
	* provide a message to describe the changes at a high level
* Run `git push --follow-tags`
	* `--follow-tags` flag will only push tags that match your commits to the shared remote repository, this will prevent any local tags being pushed to remote that are not referencing a commit which has been pushed. This flag can be set as default by adding to `push.followTags` to your git config locally, [see documentation here](https://git-scm.com/docs/git-push#Documentation/git-push.txt---follow-tags)


Once pushed, go to the Github **releases** page for the repository in question, which can be found by clicking on the `<> Code` tab and clicking on `Releases` on the right hand side:

* Select `Tags`
* Click on the `v1.7.0` tag that you just created
* Choose `Create release from tag`
* Copy the release name and make human-friendly (capitalise, remove `/`) to create a release title (e.g. `release/1.7.0 -> Release 1.7.0`)
* Add relevant release notes
* Hit **Publish release**
