#!/usr/bin/env bash

# Colors
green="\e[0;92m"
red="\e[0;91m"
nc="\e[0m" # No Color

printf "${nc}[${green}Start${nc}]\n"
if ! go build -o pingbot.out; then
  printf "${nc}[${red}COMPILE ERROR${nc}]\n"
  exit 1
else
  if ! ./pingbot.out; then
    printf "${nc}[${red}PROGRAM PANIC${nc}]\n"
    exit 1
  else
    printf "${nc}[${red}PROGRAM END${nc}]\n"
  fi
fi
