The Complete Getting Started Guide
===============
# Welcome to the team!

We're working on some guides to outline our team culture, and some of our development principles, but the big thing to remember is there are no stupid questions! If anything in these guides is not clear, please let us know so we can improve them for future team members.

# Set up your mac for development work
1. [Install applications](MAC_SETUP.md#install-applications)
2. Set up your [git config](MAC_SETUP.md#git-account-and-configuration)
3. Set up your [GPG keys](https://github.com/ONSdigital/dp-ci/blob/main/gpg-keys/developers/README.md)
4. Set up your [AWS account access](AWS_ACCOUNT_ACCESS.md)
5. Install the [dp-cli](https://github.com/ONSdigital/dp-cli) tool
6. Run `dp remote allow sandbox` to prove all your previous setup is correct. The output should include `[dp] allowing access to sandbox`. Ask for help if you encounter errors or warnings.
7. Add [environment keys](https://github.com/ONSdigital/dp-operations/blob/main/guides/gpg.md#adding-a-gpg-key-to-your-keyring) to keychain. First, you'll need to ask one of the *senior developers* for the `sandbox` passphrase or ask your *Tech Lead* for the `prod` passphrase.

Note: If you cannot access this link, something has gone wrong in the previous step.

# Install, configure and run services
1. [Install the prerequisites](INSTALLING.md#prerequisites)
2. [Clone the repositories](INSTALLING.md#clone-the-services)
3. [Setup base content](https://github.com/ONSdigital/dp-zebedee-content#dp-zebedee-content)
4. [Setup base environment configuration](INSTALLING.md#configuration)
5. [Run the apps](INSTALLING.md#running-the-apps)
6. [Update configuration](INSTALLING.md#setup-credentials) for your credentials

We've made a few attempts at simplifying these steps, but haven't converged on one way of doing things yet. The alternatives are detailed in our [developer setup](DEV_SETUP.md) guide.

### Usage guides
1. (Optional) [Enable feedback form](https://github.com/ONSdigital/dp-frontend-dataset-controller#feedback-service)
2.  [Publishing system guide](https://github.com/ONSdigital/florence/blob/develop/USAGE.md)
3. [CMD import guide](#cmd-import-steps)

# Set up for infrastructure & CI work

:warning: All steps in this section link to private repositories. If the links don't work for you, revisit the steps in [setting up your mac](#set-up-your-mac-for-development-work) to ensure you're a member of our Github organization.

For all work in this area you will first need to:
1. [Install ansible](https://github.com/ONSdigital/dp-operations/blob/main/guides/ansible.md#install-ansible)
2. [Configure access to servers](https://github.com/ONSdigital/dp-operations/blob/main/guides/ansible.md#configure-access-to-servers)

### Infrastructure
All infrastructure configuration is managed in [dp-setup](https://github.com/ONSdigital/dp-setup) and includes the following pieces of set up:
1. Install the [prerequisites for terraform](https://github.com/ONSdigital/dp-setup/blob/develop/terraform/README.md#prerequisites)
2. Install any additional [dependencies for ansible](https://github.com/ONSdigital/dp-setup/blob/develop/ansible/README.md#prerequisites)

### Continuous Integration (CI)
All CI configuration is managed in [dp-ci](https://github.com/ONSdigital/dp-ci) and includes the following pieces of set up:
1. Install any additional [dependencies for ansible](https://github.com/ONSdigital/dp-ci/tree/main/ansible#prerequisites)
2. Install the `fly` command by going to the Concourse UI and selecting the Apple from the CLI list in the bottom right of the screen.

-----

#### CMD import steps

1. [Import recipes](https://github.com/ONSdigital/dp-recipe-api) using the `import-recipes` script
    - If you're running a local graph database like Neo4j you'll also need to [import code lists](https://github.com/ONSdigital/dp-code-list-scripts) for the recipes you'll be using and may need to [import hierarchies](https://github.com/ONSdigital/dp-hierarchy-builder).
    - In most cases this is not currently needed, as we use a shared instance of Neptune running in the Development AWS account.
2. [Create datasets](https://github.com/ONSdigital/florence/tree/develop/USAGE.md#create-a-cmd-dataset-page)
3. [Import a file](https://github.com/ONSdigital/florence/tree/develop/USAGE.md#import-a-v4-file)
4. [Publish a collection](https://github.com/ONSdigital/florence/tree/develop/USAGE.md#publish-a-collection)
