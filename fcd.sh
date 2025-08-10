#!/usr/bin/env bash

ROOT_DIR="$(dirname "${BASH_SOURCE[0]}")"
FCD_FILE="$ROOT_DIR/.fcd.txt"

source "$ROOT_DIR/lib/fcd_help.sh"
source "$ROOT_DIR/lib/fcd_print.sh"
source "$ROOT_DIR/lib/fcd_add.sh"
source "$ROOT_DIR/lib/fcd_clear.sh"
source "$ROOT_DIR/lib/utils.sh"

fcd() {
  # Check
  if [[ $# -lt 2 && "$1" != "-p" && "$1" != "-h" && "$1" != "-c" ]]; then
    echo "Usage: fcd [ -a 'LABEL=\$HOME/PATH' | -r 'LABEL' | -p | -c | -h ]" >&2
    return 1
  fi
  local current_flag=""
  local current_flag_val=""
  local flag_error_message="Error: Only one flag can be used at a time."

  while [[ $# -gt 0 ]]; do
    case "$1" in
      # <--- PRINT BOOKMARKS (.fcd.txt) TO CONSOLE ---> #
      -p)
        check_flag_used "$current_flag" "$flag_error_message" || return 1
        current_flag="-p"
        fcd_print "$FCD_FILE"
        shift
        ;;
      # <--- PRINT FCD DOCUMENTATION TO CONSOLE ---> #
      -h)
        check_flag_used "$current_flag" "$flag_error_message" || return 1
        current_flag="-h"
        fcd_help
        shift
        ;;
      # <--- CLEAR ALL BOOKMARKS FROM '.fcd.txt' ---> #
      -c)
        check_flag_used "$current_flag" "$flag_error_message" || return 1
        current_flag="-c"
        fcd_clear "$FCD_FILE"
        shift
        ;;
      # <--- ADD BOOKMARK TO '.fcd.txt' ---> #
      -a)
        check_flag_used "$current_flag" "$flag_error_message" || return 1
        check_flag_arg "$2" "-a" || return 1
        current_flag="-a"
        current_flag_val="$2"
        fcd_add "$current_flag_val"
        shift 2
        ;;
      # <--- REMOVE BOOKMARK FROM '.fcd.txt' ---> #
      -r)
        check_flag_used "$current_flag" "$flag_error_message" || return 1
        check_flag_arg "$2" "-r" || return 1
        current_flag="-r"
        current_flag_val="$2"
        # fcd_remove "$current_flag_val"
        shift 2
        ;;
      # <--- USE FZF TO SEARCH BOOKMARKS ---> #
      *)
        echo "Unknown argument: '$1'" >&2
        return 1
        ;;
      esac
    done
}

