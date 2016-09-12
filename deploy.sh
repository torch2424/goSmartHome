#!/bin/bash

#Inform of deployment
echo "Deploying the project..."

#Checkout to release branch of project
cd /home/pi/goPath/src/github.com/torch2424/goSmartHome
git checkout master
git pull origin master

# Export go path and install
export GOROOT="/usr/local/go"
export GOPATH="/home/pi/goPath"
export PATH="$PATH:$GOROOT/bin:$GOPATH/bin"
go install github.com/torch2424/goSmartHome

#Restart the server, kill the old screen, and replace it
screen -X -S karenSmartHome quit
screen -S karenSmartHome -d -m bash -c "goSmartHome --server=0.0.0.0:4000"
