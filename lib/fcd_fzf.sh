source "$(dirname "${BASH_SOURCE[0]}")/utils.sh"

__fcd_fzf() {
  local FCD_FILE="$1"
  if [[ ! -f "$FCD_FILE" ]]; then
    echo "fcd: No bookmarks have been found"
    return 1
  fi

  declare -A paths=()
  declare -a labels_in_order=()

  while IFS=":" read -r label path; do
    if [[ -z "$label" || -z "$path" ]]; then
      continue
    fi

    if [[ ! -d "$path" ]]; then
      echo "Warning: Directory '$path' for '$label' does not exist" >&2
      continue
    fi

    paths["$label"]="$path"
    labels_in_order+=("$label")
  done <$FCD_FILE

  if [[ ${#paths[@]} -eq 0 ]]; then
    echo "No valid bookmarks found." >&2
    return 1
  fi

  local -a prefixed_labels=()
  local prefix_index=0
  for label in "${labels_in_order[@]}"; do
    local prefix
    # Calculates prefix "a-z" based on ASCII code
    prefix=$(printf "\\$(printf '%03o' $((97 + prefix_index % 26)))")
    if ((prefix_index >= 26)); then
    #   prefix=$(printf "\\$(printf '%03o' $((97 + (prefix_index / 26 - 1) % 26)))")$prefix
      echo "Error: Bookmark limit reached - 26" >&2
      return 1
    fi
    prefixed_labels+=("$prefix) $label")
    ((prefix_index++))
  done

  # <--- Controls Displaying The FZF SELECTION MENU ---> # 
  local selected_key
  selected_key=$(
    printf "%s\n" "${prefixed_labels[@]}" | fzf --reverse --height 60% \
      --bind 'ctrl-j:down,ctrl-k:up' \
      --header="Select bookmark (Ctrl-J/K: Navigate, Enter: Confirm)"
  )

  # <-- Checks Selected Key ---> #
  if [[ -n "$selected_key" ]]; then
    local original_label="${selected_key#*) }"
    cd "${paths[$original_label]}" || {
      echo "Error: Failed to cd to '${paths[$original_label]}'" >&2
      return 1
    }
    echo "Changed to: ${paths[$original_label]}"
  fi
}
