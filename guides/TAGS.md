Tagging a release
===========================

In order to tag a release you first need to checkout the branch you aim to tag (usually `master`)

* Run `git tag -s <tag> -m <message>`
	* `<tag>` is the semver value in the release branch name (e.g. `v1.7.0`, see [version control](VERSIONING.md))
	* `-s` ensures the tag is signed
	* provide a message to describe the changes at a high level
* `git push —-tags`

Once pushed, go to the Github **releases** page for the repository in question:

* Choose `Edit` for the tag you've just created
* Copy the release name and make human-friendly (capitalise, remove `/`) to create a release title (e.g. `release/1.7.0 -> Release 1.7.0`)
* Add relevant release notes
* Hit **Publish release**

