Tagging a release
===========================

In order to tag a release you first need to checkout the branch you aim to tag (usually master)
* Run `git tag -s <tag> -m <message>`
	* `<tag>` is the same as the branch name (e.g. `release/1.7.3` or `beta/1.0.0`)
	* `-s` ensures the tag is signed
	* provide a message to describe the changes at a high level
* `git push —tags`

Once pushed, go to the Github releases page for the repository in question:
* Choose `Edit` for the tag you've just created
* Copy the tag name and make human friendly (capitalise, remove /) to create a release title
* Add relevant release notes
* Publish release

