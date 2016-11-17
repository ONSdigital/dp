Using GPG
=========

We use GPG to sign our commits.

* Unsigned commits will not be tested, built or deployed.
* Commits signed by unknown GPG keys will not be tested, built or deployed.

### Getting started

* Ask for a [keybase.io](https://keybase.io) invite if you don't already have an account
* Get your public GPG key added to CI
* Configure Git to sign all commits

### Setting up GPG signing on Mac OS

* Install GPG Suite - https://gpgtools.org/
* Follow the [keybase-gpg-github](https://github.com/pstadler/keybase-gpg-github) guide
* Configure git:
  * `git config --global commit.gpgsign true`
  * `git config --global user.signingkey <YOUR_GPG_KEY_ID>`
* Add `no-tty` to the end of `~/.gnupg/gpg.conf` (for IntelliJ)
* Commit something and push to GitHub
* Check commit history on GitHub, should have 'Verified' badge
