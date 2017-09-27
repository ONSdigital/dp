#!/bin/bash

usage()  {
    echo "-i [set up and run dp env from scratch] || -c [just clone the dp repositories] || -h [prints this message]"; exit 1;
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

    brew cask install java ### install java
    brew install maven ### install maven
    brew install node ### install node
    npm install -g grunt-cli 

    brew install mongo ### install mongodb
    brew services restart mongo

    brew install zookeeper #install zookeeper
    brew services restart zookeeper

    brew install kafka #install kafka
    brew services restart kafka

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

    cloneGoRepo "dp-code-list-api" "cmd-develop"
    cloneGoRepo "dp-dataset-api" "cmd-develop"
    cloneGoRepo "dp-dataset-exporter" "cmd-develop"
    cloneGoRepo "dp-dimension-extractor" "cmd-develop"
    cloneGoRepo "dp-dimension-importer" "cmd-develop"
    cloneGoRepo "dp-filter-api" "cmd-develop"
    cloneGoRepo "dp-frontend-dataset-controller" "cmd-develop" 
    cloneGoRepo "dp-frontend-filter-dataset-controller" "cmd-develop"
    cloneGoRepo "dp-frontend-renderer" "cmd-develop"
    cloneGoRepo "dp-frontend-router" "cmd-develop"
    cloneGoRepo "dp-hierarchy-api" "cmd-develop"
    cloneGoRepo "dp-import-api" "cmd-develop"
    cloneGoRepo "dp-import-tracker" "cmd-develop"
    cloneGoRepo "dp-observation-extractor" "cmd-develop"
    cloneGoRepo "dp-observation-importer" "cmd-develop"
    cloneGoRepo "dp-recipe-api" "cmd-develop"
    cloneGoRepo "florence" "cmd-develop"

    cloneRepo "babbage" "develop"
    cloneRepo "zebedee" "cmd-develop"
    cloneRepo "dp-compose" "master"
    cloneRepo "sixteens" "cmd-develop"
}

cloneGoRepo() {
    if [ -d "$GOPATH/src/github.com/ONSdigital/$1" ]; then
        echo "$1 already cloned... skipping"
    else
        git clone -b $2 git@github.com:ONSdigital/$1.git $GOPATH/src/github.com/ONSdigital/$1
    fi
}

cloneRepo() {
    if [ -d "$HOME/$1" ]; then
        echo "$1 already cloned... skipping"
    else
        git clone -b $2 git@github.com:ONSdigital/$1.git $HOME/$1
    fi
}

case "$1" in
    -i)
        install
        clone
        run
        ;;
    -c)
        setGOPATH
        clone
        ;;
    -h)
        usage
        ;;
    *) 
        run
        ;;
esac

