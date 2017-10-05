#!/bin/bash

usage()  {
    echo "-i [set up dp env from scratch] || -c [just clone the dp repositories] || -p pull all repositories on the current branch || -s sets up data stores corrrectly || -h [prints this message]"; exit 1;
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

pull() {
    pullGoRepo "dp-code-list-api" 
    pullGoRepo "dp-dataset-api" 
    pullGoRepo "dp-dataset-exporter" 
    pullGoRepo "dp-dimension-extractor" 
    pullGoRepo "dp-dimension-importer" 
    pullGoRepo "dp-filter-api"
    pullGoRepo "dp-frontend-dataset-controller"
    pullGoRepo "dp-frontend-filter-dataset-controller" 
    pullGoRepo "dp-frontend-renderer" 
    pullGoRepo "dp-frontend-router" 
    pullGoRepo "dp-hierarchy-api"
    pullGoRepo "dp-import-api" 
    pullGoRepo "dp-import-tracker" 
    pullGoRepo "dp-observation-extractor"
    pullGoRepo "dp-observation-importer" 
    pullGoRepo "dp-recipe-api" 
    pullGoRepo "florence" 

    pullRepo "babbage" 
    pullRepo "zebedee" 
    pullRepo "dp-compose" 
    pullRepo "sixteens" 
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

prepare() {
    createuser dp -d -w
    createdb --owner dp FilterJobs
    psql -U dp FilterJobs -f $GOPATH/src/github.com/ONSdigital/dp-filter-api/scripts/InitDatabase.sql
    $GOPATH/src/github.com/ONSdigital/dp-dataset-api/scripts/InitDatabase.sh
    $GOPATH/src/github.com/ONSdigital/dp-code-list-api/scripts/setup.sh
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
    -s)
	prepare
	;;
    *) 
        run
        ;;
esac

