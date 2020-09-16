AWS credentials
===============

Our tooling uses AWS credential profiles to manage working with different AWS accounts. These AWS credentials are set in the `~/.aws/credentials` file as shown below.  **You must not have the AWS env vars set as this will break our tooling including dp cli, Terraform and Ansible.**

Credentials file example
------------------------

The `~/.aws/credentials` file profiles must be named the same as those shown below.  If you do not have access to production then just remove the `production` profile block.

```
[development]
aws_access_key_id=YOUR_AWS_ACCESS_KEY_ID
aws_secret_access_key=YOUR_AWS_SECRET_ACCESS_KEY
region=eu-west-1

[production]
aws_access_key_id=YOUR_SECRET_ACCESS_KEY_ID
aws_secret_access_key=YOUR_SECRET_ACCESS_KEY
region=eu-west-1
```

[Generate yourself an AWS access key](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_access-keys.html#Using_CreateAccessKey) and replace `YOUR_AWS_ACCESS_KEY_ID` and `YOUR_AWS_SECRET_ACCESS_KEY` in the above.

Working with multiple profiles
------------------------------

Our tooling, including dp cli, Terraform and Ansible, will select the correct profile automatically. If you are running other commands however you will need to specify which profile to use.

The following is the recommended approach that is safe to use with our tooling:

* Set `export AWS_DEFAULT_PROFILE=development` in your environment variables (e.g. `~/.bash_profile` or `~/.zshenv`)
* When using the aws CLI to work with production pass the `--profile=production` command line option.
* When using other tools you can set the profile using:

   ```
   AWS_PROFILE=production some_command ...
   ```
