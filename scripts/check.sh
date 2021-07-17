#! /usr/bin/env bash

echo Executing check.sh
echo
echo

echo Working directory: `pwd` \(`dirs`\)
echo File path: ./grow-player
echo Real path: `realpath ./grow-player`
echo File size: `du -h ./grow-player | awk '{print $1}'`
echo File signature: `file ./grow-player`
echo
echo

echo Executing binary:
./grow-player
bin_exit_code=$?
echo
echo

echo Finished execute. Exit status: \"$bin_exit_code\"

exit $bin_exit_code