#!/bin/bash

if [[ $DEFAULT_BRANCH == "" ]]; then
    export DEFAULT_BRANCH=cmd-develop
fi

usage()  {
    printf "To change the default branch from cmd-develop, run: export DEFAULT_BRANCH=develop (for example)\n\n"
    echo "-i [set up dp env from scratch] || -co [checkout branches on default branch] || -c [just clone the dp repositories] || -p [pull all repositories on the current branch] || -h [prints this message]"; exit 1;
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

    cloneGoRepo "dp-code-list-api" $DEFAULT_BRANCH
    cloneGoRepo "dp-dataset-api" $DEFAULT_BRANCH
    cloneGoRepo "dp-dataset-exporter" $DEFAULT_BRANCH
    cloneGoRepo "dp-dimension-extractor" $DEFAULT_BRANCH
    cloneGoRepo "dp-dimension-importer" $DEFAULT_BRANCH
    cloneGoRepo "dp-filter-api" $DEFAULT_BRANCH
    cloneGoRepo "dp-frontend-dataset-controller" $DEFAULT_BRANCH
    cloneGoRepo "dp-frontend-filter-dataset-controller" $DEFAULT_BRANCH
    cloneGoRepo "dp-frontend-renderer" $DEFAULT_BRANCH
    cloneGoRepo "dp-frontend-router" $DEFAULT_BRANCH
    cloneGoRepo "dp-hierarchy-api" $DEFAULT_BRANCH
    cloneGoRepo "dp-import-api" $DEFAULT_BRANCH
    cloneGoRepo "dp-import-tracker" $DEFAULT_BRANCH
    cloneGoRepo "dp-observation-extractor" $DEFAULT_BRANCH
    cloneGoRepo "dp-observation-importer" $DEFAULT_BRANCH
    cloneGoRepo "dp-recipe-api" $DEFAULT_BRANCH
    cloneGoRepo "florence" $DEFAULT_BRANCH
    cloneGoRepo "dp-hierarchy-builder" $DEFAULT_BRANCH
    cloneGoRepo "dp-search-builder" $DEFAULT_BRANCH
    cloneGoRepo "dp-search-api" $DEFAULT_BRANCH

    cloneRepo "babbage" $DEFAULT_BRANCH
    cloneRepo "zebedee" $DEFAULT_BRANCH
    cloneRepo "dp-compose" "master"
    cloneRepo "sixteens" $DEFAULT_BRANCH
    cloneRepo "dp-dataset-exporter-xlsx" $DEFAULT_BRANCH
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

checkout() {
    checkoutGoRepo "dp-code-list-api" $DEFAULT_BRANCH
    checkoutGoRepo "dp-dataset-api" $DEFAULT_BRANCH
    checkoutGoRepo "dp-dataset-exporter" $DEFAULT_BRANCH
    checkoutGoRepo "dp-dimension-extractor" $DEFAULT_BRANCH
    checkoutGoRepo "dp-dimension-importer" $DEFAULT_BRANCH
    checkoutGoRepo "dp-filter-api" $DEFAULT_BRANCH
    checkoutGoRepo "dp-frontend-dataset-controller" $DEFAULT_BRANCH
    checkoutGoRepo "dp-frontend-filter-dataset-controller" $DEFAULT_BRANCH
    checkoutGoRepo "dp-frontend-renderer" $DEFAULT_BRANCH
    checkoutGoRepo "dp-frontend-router" $DEFAULT_BRANCH
    checkoutGoRepo "dp-hierarchy-api" $DEFAULT_BRANCH
    checkoutGoRepo "dp-import-api" $DEFAULT_BRANCH
    checkoutGoRepo "dp-import-tracker" $DEFAULT_BRANCH
    checkoutGoRepo "dp-observation-extractor" $DEFAULT_BRANCH
    checkoutGoRepo "dp-observation-importer" $DEFAULT_BRANCH
    checkoutGoRepo "dp-recipe-api" $DEFAULT_BRANCH
    checkoutGoRepo "florence" $DEFAULT_BRANCH
    checkoutGoRepo "dp-hierarchy-builder" $DEFAULT_BRANCH

    checkoutRepo "babbage" $DEFAULT_BRANCH
    checkoutRepo "zebedee" $DEFAULT_BRANCH
    checkoutRepo "dp-compose" "master"
    checkoutRepo "sixteens" $DEFAULT_BRANCH
    checkoutRepo "dp-dataset-exporter-xlsx" $DEFAULT_BRANCH
}

checkoutRepo() {
    printf "\033[0;32m Checking out service: $1 on branch: $2 -----------------------------\n \033[0m"

    cd $HOME/$1
    git checkout $2

    printf "\n"
}

checkoutGoRepo() {
    printf "\033[0;32m Checking out service: $1 on branch: $2 -----------------------------\n \033[0m"

    cd $GOPATH/src/github.com/ONSdigital/$1
    git checkout $2

    printf "\n"
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
    pullGoRepo "dp-hierarchy-builder"

    pullRepo "babbage" 
    pullRepo "zebedee" 
    pullRepo "dp-compose" 
    pullRepo "sixteens" 
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

printf "__        __   _                       _ \n"
printf "\ \      / /__| |__  ___ _   _ ___  __| |\n"
printf " \ \ /\ / / _ \ '_ \/ __| | | / __|/ _| |\n"
printf "  \ V  V /  __/ |_) \__ \ |_| \__ \ (_| |\n"
printf "   \_/\_/ \___|_.__/|___/\__, |___/\__,_|\n"
printf "                         |___/           \n\n"

case "$1" in
    -i)
        install
        clone
        ;;
    -co)
        setGOPATH
        checkout
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
