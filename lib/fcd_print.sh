#!/usr/bin/env bash

__fcd_print() {
  if [[ ! -s "$1" ]]; then
    echo "Bookmarks empty."
  else
    cat "$1"
  fi
}
