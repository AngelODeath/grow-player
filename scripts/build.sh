#! /usr/bin/env bash

echo Executing build.sh
echo


if test -f "grow-player"; then
    echo Binary file "grow-player" already exists.
    echo - Moving to "grow-player.old"

    dest_filename=
    if test -f "grow-player.old"; then
      echo Old binary file \"grow-player.old\" already exists.
      echo - Removing \"grow-player.old\"
      rm -f grow-player.old
  fi
  mv -v grow-player grow-player.old
  echo Finished moving binary file.
fi

echo Building...
echo

# Full build
go build -v -x -a .

# Quick build
#go build -v -x .

echo Finished build. Exit status: \"$?\"
echo

./scripts/check.sh