Cloud Services and Infrastructure in Digital Publishing
===========================

We use Amazon Web Services (AWS) in Digital Publishing to host all of our services, as well as our Continuous Integration (CI).

### Pre reading

#### Why cloud computing?
>  Cloud computing is the on-demand delivery of IT resources over the Internet with pay-as-you-go pricing. Instead of buying, owning, and maintaining physical data centers and servers, you can access technology services, such as computing power, storage, and databases, on an as-needed basis from a cloud provider like Amazon Web Services (AWS). 
> -- <cite> [What is cloud computing? - AWS](https://aws.amazon.com/what-is-cloud-computing/) </cite>

This means we can:
- acquire the resources we need as we need them, and release them when we no longer need them. 
- experiment with new ideas quickly
- easily scale our infrastructure (up or down) to match the demands

Other cloud platforms include Google Cloud Platform and Azure, and are used by teams in the ONS. In theory, we should be able to move our infrastructure to any of the cloud platforms; in practice, this is non-trivial. Migrating can be harder when using Managed Services (services that the Cloud Provider manages for us), e.g. we use Amazon Neptune (a graph database) which is fully managed by AWS. If we were to change cloud providers, we might need to consider a completely different solution.

#### Fundamentals
- Cloud computing models: IaaS, PaaS, SaaS. They're [explained well on stackoverflow](https://stackoverflow.com/a/16824454) (do scroll through the answers if the first one makes no sense to you - I particularly enjoyed the [Pizza as a Service](https://stackoverflow.com/a/50355536) analogy). 

- Make sure you're familiar with basic AWS terms: 
    -  [Availability Zones and Regions](https://aws.amazon.com/about-aws/global-infrastructure/regions_az/) 
    - [Virtual Private Clouds (VPCs) and Subnets](https://docs.aws.amazon.com/vpc/latest/userguide/how-it-works.html#how-it-works-subnet)  are also good to understand at a basic level

### Pre requisite

You should be given relevant AWS accounts by your manager, or a Tech Lead - speak to your manager if this does not happen soon upon starting with the team.

### Materials
- [AWS Overview](https://docs.aws.amazon.com/whitepapers/latest/aws-overview/introduction.html)
- The workshop will walk through the [Digital Publishing 101: Basic Infrastructure slides (WIP)](https://docs.google.com/presentation/d/1VQjYd6R6xDRluA_fTF4teu97umJam2y--XlewZjeBrU/edit?usp=sharing) over roughly 1 hour, with time for questions.

### Next steps

####Exploring the AWS console
Once your account has been set up, [sign in to the AWS console](https://eu-west-1.console.aws.amazon.com) for `ons-web-development` and have a look around.

Click on VPC (in the very long list of services!) --> [Your VPCs](https://eu-west-1.console.aws.amazon.com/vpc/home?region=eu-west-1#vpcs:) and explore the `develop` VPC. 
- Can you see the range of IPs that has been assigned to this network?

Go to [Subnets](https://eu-west-1.console.aws.amazon.com/vpc/home?region=eu-west-1#subnets:) and use the filter to look at the subnets in the `develop` environment. 
- Can you see the web and publishing subnets?
- Are the IP addresses what you'd expect?
- Can you see what is part of the different availability zones?





Further resources
----------------------------
[Pluralsight](https://app.pluralsight.com/library/) provides access to course and hands-on labs:
- [AWS Developer: Getting Started](https://app.pluralsight.com/library/courses/aws-developer-getting-started/table-of-contents)
- [Create and manage users with AWS IAM](https://app.pluralsight.com/labs/detail/08ba0ff7-3064-467b-87a8-df2838f3208e)
- To understand AWS concepts such as VPC, Subnets and Security Groups in more detail, some members of our team have found [this pluralsight course](https://www.pluralsight.com/courses/aws-networking-deep-dive-vpc) very useful.



Learn about Infrastructure as Code (IaC) to understand how we provision and manage our infrastructure.
- [IaC with Terraform](https://learn.hashicorp.com/collections/terraform/aws-get-started)