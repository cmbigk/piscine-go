#!/bin/bash

#HERO_ID=70

curl -s https://01.gritlab.ax/assets/superhero/all.json | jq '.[] | select(.id == '$HERO_ID') .connections .relatives' | sed 's/\n/\\n/g' | tr -d '"'

# -s is for silent mode
# -r is the raw output without "" around it
# tr -d '"' from https://stackoverflow.com/questions/9733338/shell-script-remove-first-and-last-quote-from-a-variable