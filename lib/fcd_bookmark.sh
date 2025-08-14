__fcd_bookmark() {
  local file="$1"
  local label="$2"
  label="${label^^}"

  # Search for LABEL:PATH in file
  local entry
  entry=$(grep -E "^${label}:" "$file" 2>/dev/null)

  if [ -z "$entry" ]; then
    echo "fcd: Bookmark '$label' not found" >&2
    return 1
  fi

  # Extract the path (text after first colon)
  local path="${entry#*:}"

  # Expand ~ to $HOME
  if [ "${path:0:1}" = "~" ]; then
    path="${HOME}${path:1}"
  fi

  # Ensure path exists and is a directory
  if [ ! -d "$path" ]; then
    echo "fcd: Path for '$label' does not exist: $path" >&2
    return 1
  fi

  # Change directory
  cd "$path" || return 1
}
