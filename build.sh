#!/bin/bash

function buildContainer() {
    cd $1
    sudo podman build -t localhost/cs4389-$1 .
    cd ../
}

buildContainer api
