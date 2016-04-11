#!/bin/bash
echo "Setting up symlink for go vendor. Example: vendor/github.com -> vendor/src/github.com"

vendordir="${1:-vendor}"

cd "$vendordir"
for x in *
do
  if [[ -h "$x" ]]
  then
    echo rm "$x"
    rm "$x"
  fi
done

for x in src/*
do
  if [[ -d "$x" ]]
  then
    echo ln -s "$x" "${x##*/}"
    ln -s "$x" "${x##*/}"
  fi
done
