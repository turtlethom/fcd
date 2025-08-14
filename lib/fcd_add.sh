#!/usr/bin/env bash

source "$(dirname "${BASH_SOURCE[0]}")/utils.sh"

__fcd_add() {
  local fcd_file="$1"
  local lp_input="$2"
  local lp_label
  local lp_path
  local lp_adj_path

  # <--- CHECK IF INPUT "~/some/path" OR "LABEL:~/some/path" ---> #
  if [[ "$lp_input" == *:* ]]; then
    lp_label="${lp_input%%:*}"
    lp_path="${lp_input#*:}"
  else
    lp_path="$lp_input"
    lp_label="$(basename "$lp_input")"
  fi

  lp_label="${lp_label^^}"
  lp_adj_path="${lp_path/#\~/$HOME}"
  # Strip '/' from directory if user includes it
  lp_adj_path="${lp_adj_path#%/}"

  # Check if label has more than 25 characters
  if [[ "${#lp_label}" -gt 25 ]]; then
    echo "Error: Max character length of label exceeded"
    return 1
  fi
  # Check if path is within $HOME
  if [[ "$lp_adj_path" != "$HOME"/* ]]; then
    echo "Error: Path must be inside $HOME"
    return 1
  fi
  # Check if path is a directory
  if [[ ! -d "$lp_adj_path" ]]; then
    echo "Error: '$lp_adj_path' is not a directory"
    return 1
  fi
  # Handle finding label or path in file, write bookmark conditionally
  if __find_label_or_path "$fcd_file" "$lp_label" "$lp_adj_path"; then
    __write_bookmark "$fcd_file" "$lp_label:$lp_adj_path" 0
  else
    __write_bookmark "$fcd_file" "$lp_label:$lp_adj_path" 1
  fi
}
