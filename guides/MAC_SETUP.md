# Set up MacBook for development work

This document is part of the [Getting Started](https://github.com/ONSdigital/dp/blob/main/guides/GETTING_STARTED.md) guide.

## Install Applications

1. (Strongly recommended) Install [homebrew](https://brew.sh/)

    - a macOS package manager - useful for installing (and keeping up-to-date) apps and utilities on macOS

2. (Optional) Install [iTerm2](https://www.iterm2.com/downloads.html) (or your terminal of choice).

    - used for the [command-line](https://en.wikipedia.org/wiki/Command-line_interface) (a good proportion of our work)

    - either install from the above website, or use `brew` to install:

      ```sh
      brew install --cask iterm2
      ```

    - the `$` in the example commands represents your shell prompt (`$` is for `bash`, but you might see `%` if using `zsh`) and is not typed

3. Install your preferred browser(s)

    - the team typically uses some combination of [Chrome](https://www.google.co.uk/chrome) and [Firefox](https://www.mozilla.org/en-GB/firefox)
    - the choice is yours
    - either install from the above websites, or use `brew` to install:

      ```sh
      brew install --cask google-chrome
      ```

      ```sh
      brew install --cask firefox
      ```

4. Install [Slack](https://slack.com/intl/en-gb/downloads/mac?geocode=en-gb)

    - the primary communication app (text, voice, video, screen-sharing) within the team/department (and beyond)
    - as usual, choose the above download link, or use `brew`:

      ```sh
      brew install --cask slack
      ```

5. (Recommended) Install [Visual Studio Code](https://code.visualstudio.com/docs/setup/mac)

    - an [IDE](https://en.wikipedia.org/wiki/Integrated_development_environment) / graphical file editor
    - has many extensions for most file types and can do `go` debugging
    - if using `brew` for apps:

      ```sh
      brew install --cask visual-studio-code
      ```

6. (Recommended) MAC Software Update

    - In top left corner, Click `Apple Icon`, Click `About This Mac`, Click `Software Update...`

7. Jira

    - the primary Project Organizer within the team/department
    - you can access it via your browser

8. (Optional / Recommended) ZSH  ... the Z shell
    - this is an extended version of the `bash` shell
    - if using `brew` for apps:

      ```sh
      $ brew install zsh
      $ sudo nano /etc/shells
      ```

      Add the following to the end of the list and save the file:

      ```sh
      /usr/local/bin/zsh
      ```

      Use ```chsh``` to change the Shell line to be: `Shell: /usr/local/bin/zsh`. Save your changes and exit.

      and close/re-open terminal.

      (optional / Recommended) add `Oh My Zsh` with:

      ```sh
      % sh -c "$(curl -fsSL https://raw.githubusercontent.com/robbyrussell/oh-my-zsh/master/tools/install.sh)"
      ```

      then to further customize, refer to Step 5 onwards, in [this guide](https://www.freecodecamp.org/news/how-to-configure-your-macos-terminal-with-zsh-like-a-pro-c0ab3f3c1156/).

The MacBook is yours to configure in whatever way will help you be the most productive and comfortable. Feel free to install other software you find useful, like password managers, music players, note-taking apps.

Return to the [Getting Started](https://github.com/ONSdigital/dp/blob/main/guides/GETTING_STARTED.md) guide for next steps.

--------------

## Git account and configuration

We use [git](https://book.git-scm.com/) for version control manangement (code and documentation). You will need an up-to-date version of `git`, we recommend running:

```sh
brew install git
```

Our code - in git repositories - is shared via [Github](https://github.com), so you'll need a Github account. You may use an existing account or set one up that is specific to ONS, if you like.

1. Create a [github account](https://github.com/), if you don't already have one.

2. Give your Github username to a Tech Lead, so that they can add you to the [ONSdigital](https://github.com/ONSdigital) organisation on Github.

3. Ensure you have [enabled Multi-Factor Authentication](https://docs.github.com/en/github/authenticating-to-github/configuring-two-factor-authentication) on your Github account.

4. Configure `git` locally by running the following commands:

- `$ git config --global user.name "<YOUR_NAME>"`
- `$ git config --global user.email <YOUR_EMAIL_IN_GITHUB>`

--------------

### Making `git` work over `SSH`

In your home directory, run these commands:

- `$ ssh-keygen` and write down the passphrase you enter (somewhere _very_ safe)
- `$ cat ~/.ssh/id_rsa.pub`

Copy the displayed public key to your clipboard.

Then follow [instructions to add the SSH key in the clipboard to your github account](https://help.github.com/en/github/authenticating-to-github/adding-a-new-ssh-key-to-your-github-account)

Return to the [Getting Started](https://github.com/ONSdigital/dp/blob/main/guides/GETTING_STARTED.md) guide for next steps.

#### Testing SSH Connection

To check that you've set up your Git SSH on your terminal, open your favoured terminal and enter:

- `ssh -T git@github.com`

You will be asked to type *yes* if you're ready to connect. You should be prompted with:

- `Hi [YOUR GIT USERNAME]! You've successfully authenticated, but GitHub does not provide shell access.`

You can then close your terminal. If it's unsuccessful check you've used the correct email address (email used for Github account) when creating and saving your key.

#### Background

SSH (secure shell) is an encryption system for network communications (usually between client-server).

> Contrast with GPG (GNU Privacy Guard): an encryption system for encrypting files.

To protect sensitive data, both SSH and GPG use [cryptography](https://en.wikipedia.org/wiki/Public-key_cryptography):

- **encryption** to protect sensitive data, where that data is either:
  - _in transit,_ in the case of `ssh` (over a network connection)
  - or, _at-rest,_ in the case of `gpg` (secure a file on a disc)

- **authentication** to prove the provenance of the data
  - when logging into a server using `ssh` - your ssh key proves that you are who you claim to be
  - when sending someone a document, GPG can be used to digitally sign that document to prove that it comes from you

In all cases, the receiver of the data must already have your **public key**. So:

- we give our ssh public keys
  - to servers that we wish to login to
  - to our github settings, so that we can authenticate to github (e.g. `git push`)
- we give our GPG public keys
  - to systems (e.g. github) that read files we wish to sign (e.g. git commits are effectively files, which we sign) to prove that they came from us
  - to colleagues, so that they may send us files securely

In dissemination, we use `ssh` to login to remote machines (including our servers in the cloud), and also to secure
our communication with Github.

## OSS Index account and configuration

For Java and NodeJS projects, we use [OSS Index](https://ossindex.sonatype.org/) for auditing vulnerabilities.

This requires an account and a corresponding API token to use their services.

To get setup:

- [create an account](https://ossindex.sonatype.org/user/register) with OSS Index with your ONS email
- append the following variables to your shell startup file (e.g. `~/.zshrc`):
  - `OSSINDEX_USERNAME` should be set to the email address you signed up to OSS Index with
  - `OSSINDEX_TOKEN` is your API token which can be retrieved from the profile page in OSS Index

## Maven: Local Setup for ossindex:audit

To run mvn `ossindex:audit` or `make audit` successfully in Java projects, you must configure Maven to authenticate with the OSS Index API using your credentials.

Even though you’ve set OSSINDEX_USERNAME and OSSINDEX_TOKEN in your shell, Maven does not read environment variables directly for this plugin. Instead, it uses credentials defined in your Maven settings.xml file.

### Step 1: Confirm Your Environment Variables (Already Done?)

Ensure these are set in your shell profile (e.g. ~/.zshrc, ~/.bashrc):

```sh
export OSSINDEX_USERNAME="yourname@ons.gov.uk"
export OSSINDEX_TOKEN="your-api-token-from-ossindex"
```

### Step 2: Configure Maven settings.xml

Create or edit your Maven settings file:

```sh
~/.m2/settings.xml
```

Add the following configuration:

```sh
<settings xmlns="http://maven.apache.org/SETTINGS/1.0.0">
  <servers>
    <server>
      <id>ossindex</id>
      <username>${env.OSSINDEX_USERNAME}</username>
      <password>${env.OSSINDEX_TOKEN}</password>
    </server>
  </servers>
</settings>
```

This securely references your environment variables — your token is never stored in plain text in this file.

Make sure the `<id>ossindex</id>` matches exactly — the plugin looks for this server ID by default.

### Step 3: Confirm Java Project is configured properly to use the `ossindex-maven-plugin`

In your project’s pom.xml, configure the ossindex-maven-plugin to use this authentication:

```sh
<plugin>
  <groupId>org.sonatype.ossindex.maven</groupId>
  <artifactId>ossindex-maven-plugin</artifactId>
  <version>3.2.0</version>
  <configuration>
    <authId>ossindex</authId>
  </configuration>
</plugin> 
```  

The `authId` value in `pom.xml` should match that of `server.id` in `~/.m2/settings.xml`.

### Step 4: Reload Shell & Test

Reload your shell config:

```sh
source ~/.zshrc  # or ~/.bashrc
```

Now run the audit, If you see a scan result (not a 401 error), you're set!
