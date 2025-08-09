#!/usr/bin/env bash

fcd_clear() {
  truncate -s 0 $1
  printf "Cleared ALL Bookmarked Directories.\n"
}
