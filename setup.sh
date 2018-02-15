#!/bin/bash -e
export LC_CTYPE=C
export LANG=C

# TODO: can add OS check so scripts work on linux as well
echo "presently the script only supports mac"

parent_dir_full=$(dirname "$PWD")
parent_dir=${parent_dir_full##*/}
curr_dir=${PWD##*/} 

echo "current directory name: $curr_dir"
echo "parent directory name: $parent_dir"
if [ -z "$curr_dir" ]; then
    echo "should set current directory as first parameter"
    echo "eg: $./setup.sh <current_folder_name>"
    exit 1
fi

## GOLANG
echo "checking go version..."
if command -v "go" &>/dev/null; then
    go version
else
    echo "go is not installed. please install go."
fi

## NPM
if command -v "npm" &>/dev/null; then
    echo "npm already installed. skipping..."
else
    echo "installing npm"
    brew install npm || true
fi

echo "installing dependencies from package.json"
npm install

# replace import "github.com/niranjan92/go-hackathon-starter" with import "github.com/<your_username>/<project_dir_name>"
find . -type f -name "*.go" | xargs sed -i .backup "s/niranjan92\/go_hackathon_starter/$parent_dir\/$curr_dir/g"
find . -type f -name "*.backup" -delete

echo "checking dep"
if command -v "dep" &>/dev/null; then
    echo "dep already installed. skipping..."
else
    echo "installing dep"
    go get -u github.com/golang/dep
    export PATH=$PATH:~/go/bin
fi
dep ensure

echo "creating db. Please check that database.yml is configured correctly. Press y to continue"
#TODO: add while loop based on user's input
buffalo db create -a  || true # ignore errors here

# run migrations
buffalo db migrate

echo "finished setting up project, you can run 'buffalo dev' to start the project"

