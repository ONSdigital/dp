AWS Account Access
===============
You should receive credentials as part of onboarding. If you do not have credentials yet, please get in touch with the platform team to get an account created

- The account must be registered to an ONS email address (ons.gov.uk or ext.ons.gov.uk)
- AWS accounts are now managed by the Cloud Ops team (in DST) and are requested in ServiceNow
- You should receive an email with a link to register

Note: **Do not add prod/sandbox/staging credentials in `~/.aws/credentials`.You must not have the AWS env vars set as this will break our tooling including dp cli, Terraform and Ansible.**

Working with multiple profiles
------------------------------

Our tooling, including dp cli, Terraform and Ansible, will select the correct profile automatically, provided your credentials are set as described above.

You **must not** set environment variables for `AWS_ACCESS_KEY_ID` and `AWS_SECRET_ACCESS_KEY`, as this will prevent our tooling accessing your credentials properly.

If you are running other commands however you will need to specify which profile to use. The following is the recommended approach that is safe to use with our tooling:

* When using the aws CLI to work with production, pass the following command line option.
   ```
   --profile=dp-prod` 
   ```
* When using other tools you can set the profile using:

   ```
   AWS_PROFILE= dp-staging ...
   ```

#### AWS Command Line Access

Follow the guide at [Configuring the AWS CLI](https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-sso.html) to use AWS Single Sign-On to set up CLI access with SSO. The details needed are:
```yaml
  SSO Start URL: https://ons.awsapps.com/start
  SSO Region: eu-west-2
```

Each environment you have access to will need to be configured. Make sure you name each profile the same as the environment (dp-sandbox, dp-staging, dp-prod, dp-ci) so that it works seamlessly with other DP tooling (dp-cli, terraform). A sample config is included at the end of this guide as a reference.

## Sample config for `~/.aws/config`:

```
[profile dp-sandbox]
sso_start_url = https://ons.awsapps.com/start
sso_account_id = 1234556253 #replace this with correct account id
sso_role_name = AdministratorAccess
sso_region = eu-west-2
region = eu-west-2
```