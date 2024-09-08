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
    sudo podman build -t localhost/cs4389-$1 -f $1/Dockerfile .
}

# Build containers
buildContainer api
buildContainer webapp
