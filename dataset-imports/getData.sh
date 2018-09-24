#!/bin/bash

url=$1
graph="graph.dump"
datasets="mongodb_datasets"
imports="mongodb_imports"

curl $url$datasets --output $datasets
curl $url$imports --output $imports
curl $url$graph --output $graph
