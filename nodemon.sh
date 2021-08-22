#!/usr/bin/env bash

# Colors
RED="\e[0;91m"
NC="\e[0m" # No Color

if ! go build -o pingbot.out; then
  echo -e "${reset}[${RED}COMPILE ERROR${NC}]"
  exit 1
else
    if ! ./pingbot.out; then
      echo -e "${reset}[${RED}PROGRAM PANIC${NC}]"
      exit 1
    else
      echo -e "${reset}[${RED}PROGRAM END${NC}]"
    fi
fi
