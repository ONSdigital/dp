# Contributing guidelines

## Git workflow

- We use trunk based development - create a feature branch from `main`, e.g. `feature/new-feature`
- Pull requests must contain a succinct, clear summary of what the user need is driving this document change
- Ensure your branch contains logical atomic commits following our [commit standards](https://github.com/ONSdigital/dp-standards/blob/main/COMMIT_STANDARDS.md) before submitting a pull request
- You may rebase your branch after feedback if it's to include relevant updates from the main branch. We prefer a rebase here to a merge commit as we prefer a clean and straight history on develop with discrete merge commits for changes

## Pre-commit

We encourage the use of [pre-commit](https://pre-commit.com/) locally, this reduces the amount of common mistakes and engineers can just focus on reviewing the actual changes.

```sh
# Install pre-commit
brew install pre-commit
# This will enable pre-commit to run every time you git commit
pre-commit install
```

If you want to do an adhoc run of pre-commit

⚠️ This will only run on files that are staged in git

```sh
pre-commit run --all-files
```

If you want to commit and skip checks

```sh
git commit --no-verify -m "some message"
```

## Markdown linting

This respository uses [markdown linting](https://github.com/DavidAnson/markdownlint-cli2) for all PRs.

To install:

```sh
    brew install markdownlint-cli2
```

To run:

```sh
    markdownlint-cli2 **/*.md
```
