# File name: build.sh
# Description: Used to build all required containers.
# Author: MagnusChase03

#!/bin/bash

# Build a specified container.
# 
# Arguments:
#     - $1: The name of the folder which contains a Dockerfile to be built.
#
# Returns:
#     - N/A
function buildContainer() {
    cd $1
    sudo podman build -t localhost/cs4389-$1 .
    cd ../
}

# Build containers
buildContainer api

# Clean up old images
sudo podman system prune -f
