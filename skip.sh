#!/bin/bash

ls -l | awk 'NR%2==0'

# https://stackoverflow.com/questions/21309020/remove-odd-or-even-lines-from-a-text-file