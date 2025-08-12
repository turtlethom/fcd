#!/usr/bin/env bash

source "$(dirname "${BASH_SOURCE[0]}")/utils.sh"

fcd_add() {
  local fcd_file="$1"
  local lp_input="$2"
  local lp_label="${lp_input%%:*}"
  lp_label="${lp_label^^}"
  local lp_path="${lp_input#*:}"
  local lp_adj_path="${lp_path/#\~/$HOME}"

  # Check if label has more than 10 characters
  if [[ "${#lp_label}" -gt 10 ]]; then
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
  if find_label_or_path "$fcd_file" "$lp_label" "$lp_adj_path"; then
    write_bookmark "$fcd_file" "$lp_label:$lp_adj_path" 0
  else
    write_bookmark "$fcd_file" "$lp_label:$lp_adj_path" 1
  fi
}
