source ~/.bash_profile
cd $GOPATH/src/github.com/ONSdigital/$TASK

# Only in the first run (i.e. 'dist' folder does not exist) -> we run make dev
if [ ! -d "dist" ]; then
  echo "First run -> make dev"
  make dev
fi

make debug
