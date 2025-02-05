#!/bin/sh
DIR="$(dirname "$(readlink -f $0)")"
ON_CYAN="\033[46m"
RED="\033[31m"
YELLOW="\033[33m"
GREEN="\033[32m"
. "$DIR/mario-base.sh"
