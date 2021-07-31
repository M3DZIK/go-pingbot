#!/usr/bin/env bash

# Colors
red="\e[0;91m"
reset="\e[0m"

if ! go build -o pingbot.out; then
  echo -e "${reset}[${red}COMPILE ERROR${reset}]"
  exit 1
else
    if ! ./pingbot.out; then
      echo -e "${reset}[${red}PROGRAM PANIC${reset}]"
      exit 1
    else
      echo -e "${reset}[${red}PROGRAM END${reset}]"
    fi
fi
