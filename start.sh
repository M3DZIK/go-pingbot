#!/usr/bin/env bash

clear

# Colors
green="\e[0;92m"
red="\e[0;91m"
nc="\e[0m" # No Color

printf "${nc}[${green}Start${nc}]\n"

# Add permission and start
chmod +x pingbot.out
./pingbot.out

printf "${nc}[${red}END${nc}]\n"

# Wait 5 seconds
sleep 5

# Loop start e.g. on update => restart
bash start.sh
