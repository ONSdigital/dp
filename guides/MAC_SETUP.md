Set up MacBook for development work
====

This document is part of the [Getting Started](https://github.com/ONSdigital/dp/blob/master/guides/GETTING_STARTED.md) guide.

Install Applications
----

1. Install your preferred browser - the team use a combination of [Chrome](https://www.google.co.uk/chrome) and [Firefox](https://www.mozilla.org/en-GB/firefox) mostly, but the choice is yours.

2. Install [Slack](https://slack.com/intl/en-gb/downloads/mac?geocode=en-gb) - used for much of the communication (text, voice, video, screen-sharing) within the team/department (and beyond).

3. (Optional) Install [iTerm2](https://www.iterm2.com/downloads.html) or your terminal of choice.

4. (Strongly recommended) Install [homebrew](https://brew.sh/) - a package manager - useful for installing (and keeping up-to-date) apps on macOS.

The MacBook is yours to configure in whatever way will help you be the most productive and comfortable. Feel free to install other software you find useful, like password managers, music players, note-taking apps.

Return to the [Getting Started](https://github.com/ONSdigital/dp/blob/master/guides/GETTING_STARTED.md) guide for next steps.

--------------

Git account and configuration
----

We use [git](https://book.git-scm.com/) for version control manangement (code and documentation). You will need an up-to-date version of `git`, we recommend running:

```sh
brew install git
```

Our code - in git repositories - is shared via [Github](https://github.com), so you'll need a Github account. You may use an existing account or set one up that is specific to ONS, if you like.

1. Create a [github account](https://github.com/), if you don't already have one.

2. Give your Github username to a Tech Lead, so that they can add you to the [ONSdigital](https://github.com/ONSdigital) organisation on Github.

3. Ensure you have [enabled Multi-Factor Authentication](https://docs.github.com/en/github/authenticating-to-github/configuring-two-factor-authentication) on your Github account.

4. Configure `git` locally by running the following commands:
    * `git config --global user.name "<YOUR_NAME>"`
    * `git config --global user.email <YOUR_EMAIL_IN_GITHUB>`
    * `git config --global github.user <YOUR_GITHUB_USERNAME>`

Return to the [Getting Started](https://github.com/ONSdigital/dp/blob/master/guides/GETTING_STARTED.md) guide for next steps.
