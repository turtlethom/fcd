source "$(dirname "${BASH_SOURCE[0]}")/utils.sh"

__fcd_remove() {
  local file="$1"
  local label="$2"
  label="${label^^}"

  if ! __find_bookmark_label "$file" "$label"; then
    echo "Error: $label cannot be found" >&2
    return 1
  fi
  sed -i "/$label/d" $file
  echo "fcd: '$label' successfully removed"
}
