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

Other cloud platforms include Google Cloud Platform and Azure, and are used by teams in the ONS. In theory, we should be able to move our infrastructure to any of the cloud platforms; in practice, this is non-trivial!

#### Fundamentals
- Cloud computing models: IaaS, PaaS, SaaS. They're [explained well on stackoverflow](https://stackoverflow.com/a/16824454) (do scroll through the answers if the first one makes no sense to you - I particularly enjoyed the [Pizza as a Service](https://stackoverflow.com/a/50355536) analogy). 

- Make sure you're familiar with basic AWS terms: 
    -  [Availability Zones and Regions](https://aws.amazon.com/about-aws/global-infrastructure/regions_az/) 
    - [Virtual Private Clouds (VPCs) and Subnets](https://docs.aws.amazon.com/vpc/latest/userguide/how-it-works.html#how-it-works-subnet)  are also good to understand at a basic level

### Pre requisite

You should be given relevant AWS accounts by your manager or a Tech Lead - speak to your manager if this does not happen soon upon starting with the team. 

### Materials

The workshop will walk through the [Digital Publishing 101: Basic Infrastructure slides (WIP)](https://docs.google.com/presentation/d/1bXZw2rJTBa7spXVQNuA0lVQSN-2jZiz9khn6hJsvWwc/edit#slide=id.g950e206bef_0_474) over roughly 1 hour, with time for questions.

### Next steps

None


Further resources
----------------------------

- To understand AWS concepts such as VPC, Subnets and Security Groups in more detail, some members of our team have found [this pluralsight course](https://www.pluralsight.com/courses/aws-networking-deep-dive-vpc) very useful.

- Learn about Infrastructure as Code (IaC) to understand how we provision and manage our infrastructure.