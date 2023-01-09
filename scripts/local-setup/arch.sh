#!/usr/bin/env bash

set -eu

yay -S jre8-openjdk docker docker-compose maven cypher-shell go nodejs-lts-fermium go-yq ghostscript vault aws-session-manager-plugin
go install github.com/smartystreets/goconvey@latest

arr=(
    git@github.com:ONSdigital/dp-cli.git

    git@github.com:ONSdigital/dp-compose
    git@github.com:ONSdigital/florence
    git@github.com:ONSdigital/The-Train
    git@github.com:ONSdigital/dp-api-router
    git@github.com:ONSdigital/dp-image-api
    git@github.com:ONSdigital/dp-image-importer
    git@github.com:ONSdigital/dp-upload-service
    git@github.com:ONSdigital/dp-download-service

    git@github.com:ONSdigital/dp-dataset-api
    git@github.com:ONSdigital/dp-frontend-filter-dataset-controller
    git@github.com:ONSdigital/dp-frontend-geography-controller

    git@github.com:ONSdigital/dp-recipe-api
    git@github.com:ONSdigital/dp-import-api
    git@github.com:ONSdigital/dp-upload-service
    git@github.com:ONSdigital/dp-import-tracker
    git@github.com:ONSdigital/dp-dimension-extractor
    git@github.com:ONSdigital/dp-dimension-importer
    git@github.com:ONSdigital/dp-observation-extractor
    git@github.com:ONSdigital/dp-observation-importer
    git@github.com:ONSdigital/dp-hierarchy-builder
    git@github.com:ONSdigital/dp-dimension-search-builder
    git@github.com:ONSdigital/dp-publishing-dataset-controller

    git@github.com:ONSdigital/dp-cantabular-server
    git@github.com:ONSdigital/dp-cantabular-api-ext
    git@github.com:ONSdigital/zebedee
    git@github.com:ONSdigital/dp-recipe-api
    git@github.com:ONSdigital/dp-import-api
    git@github.com:ONSdigital/dp-dataset-api
    git@github.com:ONSdigital/dp-import-cantabular-dataset
    git@github.com:ONSdigital/dp-import-cantabular-dimension-options

    git@github.com:ONSdigital/dp-dimension-search-api
    git@github.com:ONSdigital/dp-code-list-api
    git@github.com:ONSdigital/dp-hierarchy-api
    git@github.com:ONSdigital/dp-filter-api
    git@github.com:ONSdigital/dp-dataset-exporter
    git@github.com:ONSdigital/dp-dataset-exporter-xlsx

    git@github.com:ONSdigital/dp-zebedee-content.git
)

for repo in ${arr[@]}
do
  (
    cd ../../..
    path=$(echo $repo | cut -d "/" -f 2)
    if [[ -d "$path" ]]; then
        echo "path: $path already exists... skipping"
    else
      set +e
      git clone "$repo"
      set -e
    fi
  )
done 

(cd ../../../dp-cli/cmd/dp; go install  )
(cd ../../../dp-compose; docker compose up --wait)
