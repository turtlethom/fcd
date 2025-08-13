source "$(dirname "${BASH_SOURCE[0]}")/utils.sh"

__fcd_remove() {
  local file="$1"
  local label="$2"
  label="${label^^}"
  sed -i "/$label/d" $file
  echo "fcd: '$label' successfully removed"
}
