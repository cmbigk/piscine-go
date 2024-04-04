#!/bin/bash

find . -name "*.sh" | sed 's/.sh//g' | sed 's|^./||' | sort -r | sed 's|*/||g' | rev | cut -d'/' -f1 | rev 

# taking off the ./ from file name https://serverfault.com/questions/354403/remove-path-from-find-command-output
# sorting in reverse https://www.geeksforgeeks.org/sort-command-linuxunix-examples/
# using cut to delete pathname = although should use exec or execdir or printf, but could not figure them out
# cut method from here https://askubuntu.com/questions/831546/is-there-any-way-to-make-the-cut-command-read-the-last-field-only