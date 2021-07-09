#! /usr/bin/env bash

echo Executing check.sh
echo
echo

echo Current path: `pwd`
echo GOPATH: $GOPATH

target_bin=`realpath --relative-to=.. .`

echo Filename: $target_bin
echo Full path: `realpath ./$target_bin`
echo File size: `du -h ./$target_bin | awk '{print $1}'`

echo File type: `file ./$target_bin`
echo
echo

echo Running binary:
./$target_bin
echo
echo

echo Exit check.sh
echo
echo