## Adjusting the number of ASG boxes

For example, in the publishing group:

```shell
cd dp-setup/terraform/publishing
gpg -d develop.tfvars.asc >develop.tfvars
```

Create a feature branch.

Look in file **`develop.tfvars`** at: **`asg_max_size`** and **`asg_min_size`** and increase those values by **`1`** (NOTE: they should both be the same value)

Encrypt the changed file: **`develop.tfvars`**, with:

```shell
gpg -ear develop ./develop.tfvars
```

Push the changed file (**`develop.tfvars.asc`**), do a PR, get it approved.

Apply the following commands:

(NOTE: The first init bellow may throw out some problems, so cd into appropriate directories and attend to any needed terraform init's / apply's ...)

```shell
terraform init
terraform workspace list
terraform workspace select develop
terraform plan -var-file=develop.tfvars
terraform apply -var-file=develop.tfvars
```

This may involve cd'ing into dp-ci/terraform/s3 where the the **`plan`** and **`apply`** commands would be:

```shell
terraform plan -var-file=terraform.tfvars
terraform apply -var-file=terraform.tfvars
```

Then we run the **`ansible`** scripts for publishing on develop to setup the new ASG box, using the IP address of the new box. The first ansible line below for the new BOX with the --check flag may throw up an error for some Docker thing, which is not relevant for a new BOX, so the second ansible line can be run straight away:

```shell
cd dp-setup/ansibe
ansible-playbook --vault-id=develop@.develop.pass -i inventories/develop --check -l 1.2.3.4 publishing.yml
ansible-playbook --vault-id=develop@.develop.pass -i inventories/develop -l 1.2.3.4 publishing.yml
```