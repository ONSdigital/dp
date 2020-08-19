The Complete Getting Started Guide
===============
# Welcome to the team!

We're working on some guides to outline our team culture, and some of our development principles, but the big thing to remember is there are no stupid questions! If anything in these guides is not clear, please let us know so we can improve them for future team members.

# Set up your mac for development work
1. [Install applications](MAC_SETUP.md#install-applications)
2. Set up your [git config](MAC_SETUP.md#git-account-and-configuration)
3. Set up your [GPG keys](https://github.com/ONSdigital/dp-ci/blob/master/gpg-keys/developers/README.md)* 
4. Set up your [AWS credentials](AWS_CREDENTIALS.md)
5. Install the [dp-cli](https://github.com/ONSdigital/dp-cli) tool
6. Run `dp remote allow develop` to prove all your previous setup is correct
7. Add [environment keys](https://github.com/ONSdigital/dp-ci/tree/master/gpg-keys#adding-a-gpg-key-to-your-keyring) to keychain. You'll need to ask someone for passphrases first.

*If you cannot access this link, something has gone wrong in the previous step.

# Install, configure and run services
1. [Install the prerequisites](INSTALLING.md#prerequisites)
2. [Clone the repositories](INSTALLING.md#clone-the-services)
3. Follow this separate guide to [setup base content](https://github.com/ONSdigital/dp-zebedee-content#dp-zebedee-content)
4. [Setup base environment configuration](INSTALLING.md#configuration)
5. [Run the apps](INSTALLING.md#running-the-apps)
6. [Update configuration](INSTALLING.md#setup-credentials) for your credentials

We've made a few attempts at simplifying these steps, but haven't converged on one way of doing things yet. The alternatives are detailed in our [developer setup](DEV_SETUP.md) guide. 

### Usage guides
1. (Optional) [Enable feedback form](https://github.com/ONSdigital/dp-frontend-dataset-controller#feedback-service)
2.  [Publishing system guide](https://github.com/ONSdigital/florence/blob/develop/USAGE.md)
3. [CMD import guide](#cmd-import-steps)

# Set up for infrastructure & CI work

In linking to other guides, there may be mentions of tasks that have already been completed as part of [setting up your mac](#set-up-your-mac-for-development-work), such as AWS or environment key setup. You do not need to perform these tasks again.

### Infrastructure
All infrastructure configuration is managed in [dp-setup](https://github.com/ONSdigital/dp-setup) and includes the following pieces of set up:
1. Install the [prerequisites for terraform](https://github.com/ONSdigital/dp-setup/blob/develop/terraform/README.md#prerequisites)
2. Install the [prerequisites for ansible](https://github.com/ONSdigital/dp-setup/blob/develop/ansible/README.md#prerequisites)

### Continuous Integration (CI)
All CI configuration is managed in [dp-ci](https://github.com/ONSdigital/dp-ci) and includes the following pieces of set up:
1. Check this [list of ansible prerequisites](https://github.com/ONSdigital/dp-ci/tree/master/ansible#prerequisites) and complete any steps not already performed
2. Install the `fly` command by going to the Concourse UI and selecting the Apple from the CLI list in the bottom right of the screen.

-----

#### CMD import steps

1. [Import code lists and hierarchies](https://github.com/ONSdigital/dp-cli#dp-cli) using the `dp import cmd` command
2. [Create datasets](https://github.com/ONSdigital/florence/tree/develop/USAGE.md#create-a-cmd-dataset-page)
3. [Import a file](https://github.com/ONSdigital/florence/tree/develop/USAGE.md#import-a-v4-file)
4. [Publish a collection](https://github.com/ONSdigital/florence/tree/develop/USAGE.md#publish-a-collection)
