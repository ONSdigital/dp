Common Data Terms
===========================

This training document breaks down common terms we use to describe our data, and the relationships between different data types.

## Pre-reading
None.

## Materials

### What do we publish? Datasets, editions and versions

While the ONS collect and analyse a lot of data across a variety of topics, communicating what all of that data means to our users requires organization. Statistical information can be termed as data that has been organised to serve a useful purpose - when it comes to publishing, we refer to these organized units as Datasets.

**Datasets** are collections of information on a topic, usually published on a repeating basis. As an example, a dataset may be the Annual Count of Welsh Pets.

Datasets are broken down into **editions** - predictable breakpoints in the release structure of a dataset. If this is an _annual_ dataset, there may be a different _edition_ for each year, e.g. 2018, 2019, 2020.

Within editions, we must ensure we can adapt to errors or provide more up to date information. **Versions** are the specific releases of an edition of a dataset. If a user downloads a file, it is a specific version for example, version 2 of the 2019 edition of the Annual Count of Welsh Pets dataset.

Datasets and editions use names which provide some context, whereas versions are always incrementing counts (1,2,3...)

### What kinds of datasets exist?
[This sales by product](https://www.ons.gov.uk/businessindustryandtrade/manufacturingandproductionindustry/datasets/ukmanufacturerssalesbyproductprodcom) dataset page has a 'Current' edition and the latest version is accessible through the XLSX link. Other versions can be seen in the 'Previous versions' link.

[The CPI time series](https://www.ons.gov.uk/economy/inflationandpriceindices/datasets/consumerpriceindices) dataset page doesn't feature the same edition container as time series datasets are an exception where there is only 1 fixed edition and every release is only noted as a separate version.

In a time series dataset, each version of the dataset will contain all previous data _and_ the most recent update. In an editioned dataset, each edition may only contain the information relevant to its timeframe or view on the dataset.

Some datasets are available through the Customize my Data system. These datasets are known as Filterable datasets, and follow all the same edition and versioning rules. [These datasets](https://www.ons.gov.uk/datasets/mid-year-pop-est/editions/mid-2019-april-2020-geography/versions/2) feature the edition and version information in the URL, and provide additional functionality for users to request subsets of data that they're interested in.

### What is the data made up of?

Take this CSV snippet as an example of some statistics:
``` csv
Count, Area, House Type, Pets
1, Cardiff, Detached, Dogs
28, Cardiff, Terraced, Dogs
5, Newport, Detached, Cats
98, Newport, Terraced, Cats
```
This CSV might be the version 2 file of the Annual Count of Welsh Pets we mentioned earlier. This file contains 4 **observations**, across 3 **dimensions** each featuring 2 **options**.

An **observation** is a specific statistical value. These may be counts, percentages, indexes or any other value.

A **dimension** is one of the categories of information the dataset is broken down into - the columns in this CSV: area, house type and pets.

**Dimension options** are the specific values available within a dimension. For the area dimension, Cardiff and Newport are the options.

An observation can be defined as a _"unique combination of dimension options"_. There can't be 2 possible values for Newport, Terraced, Cats; one of the options would have to change to result in a different observation figure.

### When is a version not a version? When it's an instance!

In order to decouple our import pipelines from the publishing preparation stages, we introduce an additional internal-only term: instance.

An **instance** is the term for a newly imported copy of data, before it's ready to be published. An instance will eventually become a version, but we only want to do that when we're sure that we're ready for the next version number assigned to be attached to this instance.

As an example, imagine we're publishing a weekly dataset every Monday, but we don't publish on Bank Holidays. After several bank holidays in December/January, we're loading up the 2 files we need to publish for the 2 weeks we've missed. Because they are only _instances_ at the time of upload, they can be uploaded in any order; but we need to make sure we attach them to a dataset in their chronological order so week 51 correctly comes before week 52 in the list of versions.

In this way, 'versions' aren't a completely new resource, but a version number is an additional piece of metadata attached to an 'instance' record in the database. If we decide we don't want to publish that instance after all, we can remove the version number field from the instance record and the number can be assigned to another instance as needed.

## Next steps

None.

Further resources
----------------------------
- The CSV file format used in the Customise my Data systems is known as a **V4** file, and is [further defined here](https://docs.google.com/document/d/1szyhfZZYO3c5-jeMVrXOJha5JCltCDYZC2ySNcruNOc/edit)
- Look at the 'Terminology' section on the [Developer Site](https://developer.ons.gov.uk/) for another take on some of these terms.
