# Pull Request Guidance

General guidance on creating a pull request

## Pre-reading

- It is useful to have [an understanding of GIT](https://git-scm.com/about) and branches as a way of working
- Have a look at pull requests in the `#dev_code_review` channel on Slack
- There is also this documentation on our [contributing guidelines](https://github.com/ONSdigital/dp/blob/main/guides/CONTRIBUTING.md#development-work) which covers some aspects of creating pull requests.
- There is also an option to create a *draft* pull request, these can be created to have a discussion about a feature before it is ready for a code review.

## Prerequisites

- In order to contribute you need to first [set up a GPG key](https://github.com/ONSdigital/dp-operations/blob/main/guides/gpg.md) and configure git to sign all commits.

## Steps

To merge a feature branch into the `develop` (or `master`/`main` branches) it has to have an approved pull request.

As recommended best practice before creating a pull request it can be a good idea to ensure the feature branch is up to date with the latest version of the default branch, this can help to manage conflicts. Also sometimes smaller pull requests can be easier to work with and approve.

### Creating the PR

A pull request can be created in a couple of ways, once the changes on a local feature branch have been commited and pushed.

- After pushing a branch locally a `compare and pull request` button pops up in the GitHub remote repository. A link to navigate to this page to create a pull request is in the terminal. Follow this url, to create a pull request, but just check the branch (either develop or master/main) that the pull request is being created for.
- Using the GitHub UI, navigate to the repository and the pull requests tab and use the `new pull request` button on the top right of the screen. [This GitHub documentation](https://docs.github.com/en/free-pro-team@latest/github/collaborating-with-issues-and-pull-requests/creating-a-pull-request) provide more detail on creating a pull request.

When creating a pull request the following structure is used to provide information to the reviewer:

- What: A description of the changes and why
- How to review: If required steps to enable the reviewer to do a code review
- Who can review

### Asking for code review

Once a pull request has been created put a link to the pull request in the `#dev-code-review` channel on slack. At the top of the channel are the reactions used to communicate when the pull request is being reviewed, if there are any comments and when it has been approved. After any comments on the pull requests have been addressed, this channel can be used to re request a review.

### Merging

Once the pull request has been approved it can be merged. :warning: **To merge branches the command-line has to be used, rather than the GitHub UI.** :warning:
This is because the commits have to be GPG signed/verified *by you* which is not possible via the Github UI.

## Next steps

- [Understand the release process](../../guides/RELEASES.md)
