#!/usr/bin/env bash

set -eu

brew install openjdk@8 /
             maven /
             docker-compose /
             cyhper-shell /
             node@14 /
             go /
             ghostscript /
             hasicorp/tap/vault /
             jq /
             yq

brew install --cask docker

go install github.com/smartystreets/goconvey@latest

