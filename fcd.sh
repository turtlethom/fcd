#!/usr/bin/env bash

ROOT_DIR="$(dirname "${BASH_SOURCE[0]}")"
FCD_FILE="$ROOT_DIR/.fcd.txt"
source "$ROOT_DIR/lib/fcd_help.sh"
source "$ROOT_DIR/lib/fcd_print.sh"
source "$ROOT_DIR/lib/fcd_add.sh"
source "$ROOT_DIR/lib/fcd_clear.sh"
#############################################################################
#<< ***** --- ( FCD ) --- ***** >># 
# ->* Stands for 'fast/fuzzy change directory'
# ->* Command that utilizes fzf for quickly changing into bookmarked directories
# ->* Security-hardened fzf directory navigator with shortcuts & autocomplete.
# ***** --- REQUIREMENTS --- ***** #
# ->* fzf (for interactive selection)
#############################################################################

fcd() {
  # Confirm user is using command flag
  if [[ $# -lt 2 && "$1" != "-p" && "$1" != "-h" && "$1" != "-c" ]]; then
    echo "Usage: fcd [-a string | -c path | -r word | -p]" >&2
    return 1
  fi
  # Check command flags
  current_flag=""
  current_flag_val=""
  flag_error_message="Error: Only one flag can be used at a time."
  while [[ $# -gt 0 ]]; do
    case "$1" in
      # PRINT TO CONSOLE
      -p)
        if [[ -n "$current_flag" ]]; then
          echo "$flag_error_message" >&2
          return 1
        fi
        current_flag="-p"
        current_flag_val=""
        # Logic For Printing Bookmarks To Console
        fcd_print "$FCD_FILE"
        shift
        ;;
      # Help With Command Syntax
      -h)
        if [[ -n "$current_flag" ]]; then
          echo "$flag_error_message" >&2
          return 1
        fi
        current_flag="-h"
        current_flag_val=""
        # Print fcd documentation
        fcd_help
        shift
        ;;
      -c)
        if [[ -n "$current_flag" ]]; then
          echo "$flag_error_message" >&2
          return 1
        fi
        current_flag="-c"
        current_flag_val=""
        fcd_clear "$FCD_FILE"
        shift
        ;;
      # ADD BK, REMOVE BK
      -a|-r)
        if [[ -n "$current_flag" ]]; then
          echo "$flag_error_message" >&2
          return 1
        fi
        if [[ -z $2 || $2 == -* ]]; then
          echo "Error: Flag '$1' requires a argument." >&2
          return 1
        fi
        current_flag="$1"
        current_flag_val="$2"
        # Handle each
        if [[ "$current_flag" == "-a" ]]; then
          # Handle adding a label:directory
          fcd_add "$current_flag_val"
        elif [[ "$current_flag" == "-r" ]]; then
          # Handle removing a label:directory based on label
          echo "Using -r"
        fi
        shift 2
        ;;
      # Default Case
      *)
        echo "Unknown argument: '$1'" >&2
        return 1
        ;;
      esac
    done
}
