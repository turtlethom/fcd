#!/usr/bin/env bash

# <--- ENFORCES SINGLE FLAG USAGE ---> #
check_flag_used() {
  local flag="$1"
  local error_msg="$2"
  if [[ -n "$flag" ]]; then
    echo "$error_msg" &>2
    return 1
  fi
  return 0
}

# <--- Check user args contain only one flag --> #
check_flag_arg() {
  local arg="$1"
  local flag="$2"
  if [[ -z "$arg" || "$arg" == -* ]]; then
    echo "Error: '$flag' requires an argument." >&2
    return 1
  fi
  return 0
}

# <--- Check if bookmark label exists ---> #
find_bookmark_label() {
  local fcd_file="$1"
  local lp_label="$2"
  grep -qE "^$lp_label:" "$fcd_file";
}

# <-- Check if bookmark path exists ---> #
find_bookmark_path() {
  local fcd_file="$1"
  local path="$2"
  grep -qE ":$path\$" "$fcd_file";
}

# <--- Find bookmark's label or path in file 'LABEL:$HOME/...' --> #
find_label_or_path() {
  local fcd_file="$1"
  local lp_label="$2"
  local lp_path="$3"
  find_bookmark_label "$fcd_file" "$lp_label" || find_bookmark_path "$fcd_file" "$lp_path"
}

find_label_and_path() {
  local fcd_file="$1"
  local lp_label="$2"
  local lp_path="$3"
  find_bookmark_label "$fcd_file" "$lp_label" && find_bookmark_path "$fcd_file" "$lp_path"
}

write_bookmark() {
  local fcd_file="$1"
  local lp_input="$2"
  local found="$3"

  if (( found != 0 )); then
    echo "$lp_input" >> "$fcd_file"
    echo "Bookmarked '$lp_input'."
  else
    echo "'$lp_input' already bookmarked."
  fi
}
