#!/bin/sh
DIR="$(dirname "$(readlink -f $0)")"
ON_CYAN="\033[48;2;16;192;192m"
RED="\033[38;2;192;0;16m"
YELLOW="\033[38;2;160;160;16m"
GREEN="\033[38;2;16;128;16m"
. "$DIR/mario-base.sh"
