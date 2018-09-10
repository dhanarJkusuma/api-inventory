#!/bin/bash

# Change Directory Bin
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null && pwd )"
echo "[Inventory App] Changing directory $DIR"

appBin="$DIR/app"
if [ -f "$appBin" ]
then
	rm $appBin
fi

# Change Directory App
cdCmd="cd $DIR/../"
eval $cdCmd

# Change to Directory DB
eval "cd db"
if [ -f "app.db" ]
then
	eval "rm app.db"
fi

eval "touch app.db"


# Change Directory App
cdCmd="cd $DIR/../"
eval $cdCmd

# Downloading Dependency
echo "[Inventory App] Resolving Dependency"
eval "go get -v"

echo "[Inventory App] Build App"
# Build/Compile
buildCmd="go build -o $DIR/app ."
eval $buildCmd > $DIR/build.log

eval "./start.sh"