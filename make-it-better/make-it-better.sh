#!/bin/bash

#create file "0" with permission dr-------x, creation date 1986-01-05 and creation time 00:00

# dr-------x 1986-01-05 00:00 0
# -r------w- 1986-11-13 00:01 1
# -rw----r-- 1988-03-05 00:10 2
# lrwxrwxrwx 1990-02-16 00:11 3 -> 0
# -r-x--x--- 1990-10-07 01:00 4
# -r--rw---- 1990-11-07 01:01 5
# -r--rw---- 1991-02-08 01:10 6
# -r-x--x--- 1991-03-08 01:11 7
# -rw----r-- 1994-05-20 10:00 8
# -r------w- 1994-06-10 10:01 9
# dr-------x 1995-04-10 10:10 A

chmod 777 *

# mkdir {test1,test2,test3}
mkdir 0
mkdir A

TZ=utc touch -t 8601050000 0
TZ=utc touch -t 9504101010 A

# create a file using a specified creation time
# touch -t YYMMDDHHMM fileName

TZ=utc touch -t 8611130001 1
TZ=utc touch -t 8803050010 2
TZ=utc touch -t 9010070100 4
TZ=utc touch -t 9011070101 5
TZ=utc touch -t 9102080110 6
TZ=utc touch -t 9103080111 7
TZ=utc touch -t 9405201000 8
TZ=utc touch -t 9406101001 9



# ln -s source_file symbolic_link
# for a directory it will be: ln -s /home/james james
ln -s 0 3
TZ=utc touch -h -t 9002160011 3
TZ=utc touch -h -t 8601050000 0



# set permission now
chmod 401 0 # 4 = r--, 0 = ---, 1 = --x
chmod 402 1 # 2 = -w-
chmod 604 2 # 6 = rw-
chmod -h 777 3 # 7 = rwx # -h is for changing permissions on symbolic link: ref man chmod
chmod 510 4 # 5 = r-x
chmod 460 5
chmod 460 6
chmod 510 7
chmod 604 8
chmod 402 9
chmod 401 A
chmod 401 0

tar --exclude="*.sh" -cf done.tar * #this command packs .sh file itself, trying to improve it

ls -l
