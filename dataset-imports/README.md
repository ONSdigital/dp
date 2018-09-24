
CMD Dataset Imports
===================

This repo is for loading datasets via database imports onto a developers local machine.

 # IMPORTANT

 To use this mechanism, you must have the stack running and be willing to have the scripts blank your
 neo4j graph database as well as your datasets and imports mongo databases.

 # Usage

Note - all the files use the same naming conventions (graph.dump etc) so before downloading a new one, make
sure to move or delete any previous files downloaded for importing.

* Clone the dp repo and navigate to this directory.
* `chmod +x getData.sh`
* `./getData.sh <url for s3 bucket>`   # links provided below
* `chmod +x loadData.sh`
* `./loadData.sh`

If you have already associated the target dataset with a collection (i.e if an `API Dataset` page already exists)
then the imported dataset will not show - just delete then re-create the `API Dataset` page to fix.



 # S3 Links For Importing Datasets

 `https://s3-eu-west-1.amazonaws.com/ons-dp-cmd-database-imports/data_midYearPopEst/` - Mid Year Population Estimates
