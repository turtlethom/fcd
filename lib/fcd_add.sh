#!/usr/bin/env bash

fcd_add() {
  local fcd_file="$1"
  local lp_input="$2"
  local lp_label="${lp_input%%:*}"
  local lp_path="${lp_input#*:}"
  local lp_adj_path="${lp_path/#\~/$HOME}"

  # Check if label has more than 10 characters
  if [[ "${#lp_label}" -gt 10 ]]; then
    echo "Error: Labels must be 10 characters or less"
    return 1
  fi

  if [[ "$lp_adj_path" != "$HOME"/* ]]; then
    echo "Error: Path must be inside $HOME"
    return 1
  fi

  if [[ ! -d "$lp_adj_path" ]]; then
    echo "Error: '$lp_adj_path' is not a directory"
    return 1
  fi

  # Read file to check if path exists
  local found
  found=$(read_and_check_file "$fcd_file" "$lp_label" "$lp_adj_path")
  write_bookmark "$fcd_file" "$lp_label:$lp_adj_path" "$found"
}

read_and_check_file() {
  local fcd_file="$1"
  local lp_label="$2"
  local lp_adj_path="$3"

  if grep -qE "^$lp_label:" "$fcd_file" || grep -qE ":$lp_adj_path\$" "$fcd_file"; then
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
