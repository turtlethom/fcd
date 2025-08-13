#!/usr/bin/env bash

__fcd_clear() {
  truncate -s 0 $1
  printf "Cleared ALL Bookmarked Directories.\n"
}
