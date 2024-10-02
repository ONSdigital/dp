Provisioning
===========================

## Terraform
> Terraform is an open-source infrastructure as code software tool that provides a consistent CLI workflow to manage hundreds of cloud services. Terraform codifies cloud APIs into declarative configuration files. It allows infrastructure to be expressed as code in a simple, human readable language called HCL (HashiCorp Configuration Language). It reads configuration files and provides an execution plan of changes, which can be reviewed for safety and then applied and provisioned.

Terraform provisioning builds the infrastructure for the test and live environments (`develop` and `production`) which involves setting up stacks (`app-users`, `web`, `publishing`, `mongodb`, etc). Therefore, whenever an infrastructure change is required, terraform needs to be run. The use of terraform is located in [dp-setup/terraform][dp-setup-terraform]
 
More information about Terraform can be found on their website - https://www.terraform.io/

## Ansible
> Ansible is an open-source software provisioning and configuration management tool enabling infrastructure as code. It uses no agents and no additional custom security infrastructure, so it's easy to deploy and uses YAML, in the form of Ansible Playbooks, that allows to describe the automation jobs. 

Ansible provisioning can be considered as adding and configuring applications *to* the test and production service environments in the cloud (`develop` and `production`). Therefore, whenever the change to underlying software is required, ansible needs to be run. The use of ansible is located in [dp-setup/ansible][dp-setup-ansible]

More information about Ansible can be found on their website - https://www.ansible.com/overview/how-ansible-works

[//]: # (Reference Links and Images)
[dp-setup-ansible]: <https://github.com/ONSdigital/dp-setup/tree/awsb/ansible>
[dp-setup-terraform]: <https://github.com/ONSdigital/dp-setup/tree/awsb/terraform>