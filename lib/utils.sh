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

# <--- CHECK IF ARGUMENT FOR FLAG EXISTS (NOT ANOTHER FLAG)
check_flag_arg() {
  local arg="$1"
  local flag="$2"
  if [[ -z "$arg" || "$arg" == -* ]]; then
    echo "Error: '$flag' requires an argument." >&2
    return 1
  fi
  return 0
}
