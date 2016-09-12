#!/bin/bash

#Checkout to release branch of project
git checkout master
git pull origin master
go install github.com/torch2424/goSmartHome

#Restart the server, kill the old screen, and replace it
screen -X -S karenSmartHome quit
screen -S karenSmartHome -d -m bash -c "goSmartHome --server=0.0.0.0:4000"
