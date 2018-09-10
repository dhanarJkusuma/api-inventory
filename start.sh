#!/bin/bash

# Change Directory Bin
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null && pwd )"
echo "[Inventory App] Changing directory $DIR--"

echo "[Inventory App] Checking Bin"
appBin="$DIR/bin"
if [ -f "$appBin/app" ]
then
    eval "$appBin/app"
else
    echo "[Inventory App] Trying to Install Inventory Application"
    # Change Directory App/Bin
    cdCmd="cd $appBin"
    eval $cdCmd
    eval "./install.sh"
fi
