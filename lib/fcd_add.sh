#!/usr/bin/env bash

fcd_add() {
  local fcd_file="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)/../.fcd.txt"
  local lp_input="$1"
  local lp_label="${lp_input%%=*}"
  local lp_path="${lp_input#*=}"

  # Check if path is a part of $HOME, is a directory
  if [[ ! "$lp_path" == $HOME/* || ! -d "$lp_path" ]]; then
    echo "Error: Argument must match 'LABEL=$HOME/...'"
    return 1
  fi

  # Check if label has more than 10 characters
  if [[ "${#lp_label}" -gt 10 ]]; then
    echo "Error: Labels must be 10 characters or less"
    return 1
  fi

  # Read file to check if path exists
  local found
  found=$(read_and_check_file "$fcd_file" "$lp_label" "$lp_path")
  write_bookmark "$fcd_file" "$lp_input" "$found"
}

read_and_check_file() {
  local fcd_file="$1"
  local lp_path="$2"
  local lp_label="$3"
  if grep -qE "^$lp_label=" "$fcd_file" || grep -qE "=$lp_label\$" "$fcd_file"; then
    echo true
  else
    echo false
  fi
}

write_bookmark() {
  local fcd_file="$1"
  local lp_input="$2"
  local found="$3"
  if [[ "$found" == false ]]; then
    echo "$lp_input" >> "$fcd_file"
    echo "Bookmarked '$lp_input'."
  else
    echo "'$lp_input' already bookmarked."
  fi
}
