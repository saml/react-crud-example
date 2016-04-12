#!/bin/bash

setup_vendor_symlinks() {
  vendordir="${1:-vendor}"
  cwd="${PWD}"

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

  cd "$cwd"
}

typings install &&
wgo restore &&
setup_vendor_symlinks &&
make

