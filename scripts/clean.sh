#! /usr/bin/env sh


mkdir -pv _trash

rm -vf _trash/*

# find . \( -path ./scripts -o -path ./scratch -o -path ./libs \) -prune -false -o -type f ! -name '*.go' ! -name '*.mod' ! -name '*.sum' ! -name '*.md' -exec rm -v {} \;

find . \
\( -path ./scripts -o -path ./scratch -o -path ./libs -o -path ./.idea \) -prune -false \
-o -type f \
! -name '*.go' ! -name '*.mod' ! -name '*.sum' ! -name '*.md' \
-exec mv -v {} _trash/ \;
