# File name: build.sh
# Description: Used to build all required containers.
# Author: MagnusChase03

#!/bin/bash

export CR='sudo podman'
if [ $# -gt 0 ]; then
    export CR=$1
fi 

# Build a specified container.
# 
# Arguments:
#     - $1: The name of the container
#     - $2: The name of the folder which contains a Dockerfile to be built.
#
# Returns:
#     - N/A
function buildContainer() {
    $CR build -t localhost/cs4389-$1 -f $2/Dockerfile .
}


# Build containers
buildContainer api ./api
buildContainer webapp ./webapp
buildContainer redis ./databases/redis
buildContainer mariadb ./databases/mariadb
