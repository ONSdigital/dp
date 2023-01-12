#!/usr/bin/env bash

set -eu

if [[ "$OSTYPE" == "linux-gnu"* ]]; then
    yay -S jre8-openjdk docker docker-compose maven cypher-shell go nodejs-lts-fermium go-yq ghostscript vault aws-session-manager-plugin aws-cli-v2 ansible python-boto3 jdk8-openjdk


elif [[ "$OSTYPE" == "darwin"* ]]; then
    brew install openjdk@8 maven docker-compose cyhper-shell node@14 go ghostscript jq yq
                
    # hasicorp/tap/vault /
    brew install --cask docker session-manager-plugin
fi

printf "Please write your SSH key name.\nThis is usually FirstnameLastname\n\n"
read SSHUSER


go install github.com/smartystreets/goconvey@latest
go install github.com/DarthSim/overmind/v2@latest

arr=(
    git@github.com:ONSdigital/dp-cli
    git@github.com:ONSdigital/dp-setup
    git@github.com:ONSdigital/dp-ci
    git@github.com:ONSdigital/dp-code-list-scripts
    git@github.com:ONSdigital/dp-hierarchy-builder
    git@github.com:ONSdigital/dp-configs

    git@github.com:ONSdigital/babbage
    git@github.com:ONSdigital/zebedee
    git@github.com:ONSdigital/sixteens

    git@github.com:ONSdigital/dp-frontend-router
    git@github.com:ONSdigital/dp-frontend-renderer

    git@github.com:ONSdigital/dp-frontend-homepage-controller
    git@github.com:ONSdigital/dp-frontend-cookie-controller
    git@github.com:ONSdigital/dp-frontend-dataset-controller

    git@github.com:ONSdigital/dp-frontend-feedback-controller

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

(cd ../../../dp-cli/cmd/dp; go install)
(cd ../../../dp-compose; docker compose up --wait)
(cd ../../../dp-zebedee-content; make install)
(cd ../../../dp-cli; cp -a config/example_config.yml ~/.dp-cli-config.yml)
(cd ../../../dp-cli; cp -a config/example_config.yml ~/.dp-cli-config.yml)
(cd ../../../dp-ci; git checkout main; git pull)
(cd ../../../dp-setup; git checkout main; git pull)

working_dir=$(cd ../../..; pwd)

VAR=$(cat <<EOF
dp-setup-path: "$working_dir/dp-setup" # The path to the dp-setup repo on your machine.
dp-ci-path: "$working_dir/dp-ci" # The path to the dp-ci repo on your machine.
dp-hierarchy-builder-path: "$working_dir/dp-hierarchy-builder" # The path to the dp-hierarchy-builder repo on your machine.
dp-code-list-scripts-path: "$working_dir/dp-code-list-scripts" # The path to the dp-code-list-scripts repo on your machine.
EOF
)

echo "$VAR 
$(tail +8 ~/.dp-cli-config.yml)" > ~/.dp-cli-config.yml
sed -i "s/ssh-user: JamesHetfield/ssh-user: ${SSHUSER}/" ~/.dp-cli-config.yml

STARTUP_FILE=$(cat <<'EOF'
# Digital Publishing services
export zebedee_root=~/Documents/website/zebedee-content/generated
export ENABLE_PRIVATE_ENDPOINTS=true
export ENABLE_PERMISSIONS_AUTH=true
export ENCRYPTION_DISABLED=true
export DATASET_ROUTES_ENABLED=true
export FORMAT_LOGGING=true
export SERVICE_AUTH_TOKEN="fc4089e2e12937861377629b0cd96cf79298a4c5d329a2ebb96664c88df77b67"

export TRANSACTION_STORE=$zebedee_root/zebedee/transactions
export WEBSITE=$zebedee_root/zebedee/master
export PUBLISHING_THREAD_POOL_SIZE=10
EOF
)

(
ons_
if [[ -f ~/.ons_startup_file.lock ]]; then
   echo ""
else
    echo "$STARTUP_FILE" > ~/.ons_startup_file
    if [[ -f ~/.bashrc ]]; then
        echo "source ~/.ons_startup_file" >> ~/.bashrc
        touch .ons_startup_file.lock
    fi
    if [[ -f ~/.zshrc ]]; then
        echo "source ~/.ons_startup_file" >> ~/.zshrc
        touch .ons_startup_file.lock
    fi
fi
source ~/.ons_startup_file
)

mkdir -p "$zebedee_root"
cp -a ~/Downloads/cms-content.zip $zebedee_root

dp remote allow sandbox
dp-zebedee-content generate

overmind start -f Procfile.web
