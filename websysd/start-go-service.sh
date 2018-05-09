if [[ $1 != "" ]]; then
    TASK=$1
fi

source ~/.bash_profile
cd $GOPATH/src/github.com/ONSdigital/$TASK
make debug
