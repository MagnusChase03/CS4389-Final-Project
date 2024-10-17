# File name: test.sh
# Description: Small utility function to test route functionallity.
# Author: MagnusChase03

#!/bin/bash

if [ $# -lt 3 ]; then
    echo "Usage: ./test.sh <data> <cookie> <route>"
    exit 1
fi

curl -X POST -d "$1" --cookie "$2" -k -v "https://localhost:8081$3"
