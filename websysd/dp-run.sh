#!/bin/bash

usage()  {
    echo "-i [set up dp env from scratch] || -c [just clone the dp repositories] || -p pull all repositories on the current branch || -h [prints this message]"; exit 1;
}

run() {
    websysd
}

setGOPATH() {
    if [ "$GOPATH" == "" ]; then
        echo "setting GOPATH"
        tee -a ~/.bash_profile <<< "export GOPATH=~/go"
        tee -a ~/.bash_profile <<< "export PATH=$PATH:$GOPATH/bin"
        exec bash -l
    fi
}

install() {
    echo "starting dp environment installation (this may take several minutes)... ";

    /usr/bin/ruby -e "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)" ### install homebrew

    brew install go ### install golang
    setGOPATH

    brew cask install java8 ### install java8
    brew install maven ### install maven
    brew install node ### install node
    npm install -g grunt-cli 

    brew install mongo ### install mongodb
    brew services restart mongo

    brew install kafka #install kafka
    brew services restart zookeeper
    /usr/local/Cellar/kafka/*/libexec/bin/kafka-server-start.sh -daemon /usr/local/Cellar/kafka/*/libexec/config/server.properties

    brew install postgres #install postgres
    brew services restart postgres

    brew install docker # install docker
    brew services restart docker

    brew install neo4j #install neo
    tee -a /usr/local/Cellar/neo4j/*/libexec/conf/neo4j.conf <<< "dbms.security.auth_enabled=false"
    brew services restart neo4j

    go get github.com/websysd/websysd #get websysd
}

clone() {
    echo "cloning dp repositories"

    mkdir -p $GOPATH/src/github.com/ONSdigital

    cloneGoRepo "dp-api-router" "develop"
    cloneGoRepo "dp-code-list-api" "develop"
    cloneGoRepo "dp-dataset-api" "develop"
    cloneGoRepo "dp-dataset-exporter" "develop"
    cloneGoRepo "dp-dimension-extractor" "develop"
    cloneGoRepo "dp-dimension-importer" "develop"
    cloneGoRepo "dp-download-service" "develop"
    cloneGoRepo "dp-filter-api" "develop"
    cloneGoRepo "dp-frontend-dataset-controller" "develop"
    cloneGoRepo "dp-frontend-filter-dataset-controller" "develop"
    cloneGoRepo "dp-frontend-geography-controller" "develop"
    cloneGoRepo "dp-frontend-renderer" "develop"
    cloneGoRepo "dp-frontend-router" "develop"
    cloneGoRepo "dp-hierarchy-builder" "develop"
    cloneGoRepo "dp-hierarchy-api" "develop"
    cloneGoRepo "dp-import-api" "develop"
    cloneGoRepo "dp-import-tracker" "develop"
    cloneGoRepo "dp-import-reporter" "develop"
    cloneGoRepo "dp-observation-extractor" "develop"
    cloneGoRepo "dp-observation-importer" "develop"
    cloneGoRepo "dp-recipe-api" "develop"
    cloneGoRepo "dp-search-builder" "develop"
    cloneGoRepo "dp-search-api" "develop"
    cloneGoRepo "florence" "develop"

    cloneRepo "babbage" "develop"
    cloneRepo "zebedee" "develop"
    cloneRepo "sixteens" "develop"
    cloneRepo "the-train" "develop"
    cloneRepo "dp-dataset-exporter-xlsx"
}

cloneGoRepo() {
    if [ -d "$GOPATH/src/github.com/ONSdigital/$1" ]; then
        echo "$1 already cloned... skipping"
    else
        go get github.com/ONSdigital/$1
    fi
}

cloneRepo() {
    if [ -d "$HOME/$1" ]; then
        echo "$1 already cloned... skipping"
    else
        git clone -b $2 git@github.com:ONSdigital/$1.git $HOME/$1
    fi
}

pull() {
    pullGoRepo "dp-api-router"
    pullGoRepo "dp-code-list-api" 
    pullGoRepo "dp-dataset-api" 
    pullGoRepo "dp-dataset-exporter" 
    pullGoRepo "dp-dimension-extractor" 
    pullGoRepo "dp-dimension-importer" 
    pullGoRepo "dp-download-service"
    pullGoRepo "dp-filter-api"
    pullGoRepo "dp-frontend-dataset-controller"
    pullGoRepo "dp-frontend-filter-dataset-controller" 
    pullGoRepo "dp-frontend-geography-controller"
    pullGoRepo "dp-frontend-renderer" 
    pullGoRepo "dp-frontend-router" 
    pullGoRepo "dp-hierarchy-api"
    pullGoRepo "dp-hierarchy-builder"
    pullGoRepo "dp-import-api" 
    pullGoRepo "dp-import-tracker" 
    pullGoRepo "dp-observation-extractor"
    pullGoRepo "dp-observation-importer" 
    pullGoRepo "dp-recipe-api"
    pullGoRepo "dp-search-api"
    pullGoRepo "dp-search-builder"
    pullGoRepo "florence"

    pullRepo "babbage" 
    pullRepo "zebedee" 
    pullRepo "dp-compose" 
    pullRepo "sixteens" 
    pullRepo "the-train"
    pullRepo "dp-dataset-exporter-xlsx"
}

pullRepo() {
    if [ -d "$HOME/$1" ]; then
        echo "pulling $1..."
        cd $HOME/$1
        git pull
    else
        echo "repo $1 missing... please clone"
    fi
}

pullGoRepo() {
    if [ -d "$GOPATH/src/github.com/ONSdigital/$1" ]; then
        echo "pulling $1..."
        cd $GOPATH/src/github.com/ONSdigital/$1
        git pull
    else
        echo "repo $1 missing... please clone"
    fi
}

case "$1" in
    -i)
        install
        clone
        ;;
    -c)
        setGOPATH
        clone
        ;;
    -h)
        usage
        ;;
    -p)
        pull
        ;;
    *) 
        run
        ;;
esac

