Gitflow
=======

A useful resource outlining Gitflow can be found [here](https://www.atlassian.com/git/tutorials/comparing-workflows/gitflow-workflow). This should provide a good understaning of the Gitflow model.

The below will use the term `main` for the live branch - some repositories may still use the legacy term `master`.

## What is Gitflow?

Gitflow is essentially a branching model of Git, and the model used at ONS Digital.

It allows developers to manage source code for regular releases, and provides a framework for developers to be able to work on features separately from one another without any unwanted clashes.

At its basic level, the Gitflow model will use two permanent branches - `main` and `develop`. The `main` branch is used for the production code, and the `develop` branch is used to check code before going to `main`.

Each of the features developers are working on also get their own branch. Usually they will branch off from `develop`, not from `main` (however, in some cases this may not be true, and `feature` branches come directly from the `main` branch). Once the the work on the feature is done, it is then merged into `develop`.

## Gitflow Branches

### Main

This is the live branch of the project and contains the last release version of source code in production.

### Develop

This branch serves as a branch for integrating different features planned for an upcoming release. It generally where developers merge `feature` branches into.

### Feature

Feature branches are used to develop new features for the upcoming release. The `feature` branch will exist whilst the feature is in development, but will eventually be merged back into `develop`.

### Release

Release branches support preparation of a new production release. Creating this branch starts the next release cycle. Once it's ready to ship, the `release` branch gets merged into `main` and tagged with a version number. See [here](https://github.com/ONSdigital/dp/blob/main/guides/VERSIONING.md) for more information on version control at ONS Digital.
