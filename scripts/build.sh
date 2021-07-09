#! /usr/bin/env bash

echo Executing build.sh
echo

echo Current path: `pwd`
echo GOPATH: $GOPATH

target_bin=`realpath --relative-to=.. .`

if test -f "$target_bin"; then
    echo Binary file "$target_bin" already exists.
    echo - Moving to "$target_bin.old"

    dest_filename="$target_bin.old"
    if test -f "$dest_filename"; then
      echo Old binary file \"$dest_filename\" already exists.
      echo - Removing \"$dest_filename\"
      rm -f $dest_filename
  fi
  mv -v $target_bin $dest_filename
  echo Finished moving binary file.
fi

echo Building...
echo
# Full build
go build -o $target_bin -v -x -a .

# Quick build
#go build -o $target_bin -v -x .

echo Finished build. Exit status: \"$?\"
echo

echo Exit build.sh
echo

./scripts/check.sh