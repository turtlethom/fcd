#!/bin/bash

#############################################################################
#<< ***** --- ( FCD ) --- ***** >># 
# ->* Stands for 'fast/fuzzy change directory'
# ->* Command that utilizes fzf for quickly changing into bookmarked directories
# ->* Security-hardened fzf directory navigator with shortcuts & autocomplete.
# ***** --- REQUIREMENTS --- ***** #
# ->* fzf (for interactive selection)
#############################################################################

fcd() {
  local BOOKMARKED="./dev/.pcd.txt"
  declare -a ORDERED_LABELS
  declare -A PATHS
  local SELECTED_KEY

  # Confirm user is using command flag
  if [[ $# -lt 2 && "$1" != "-p" ]]; then
    echo "Usage: fcd [-a string | -c path | -r word | -p]" >&2
    return 1
  fi
  # Check command flags
  current_flag=""
  current_flag_val=""
  flag_error_message="Error: Only one flag can be used at a time.">&2
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
        echo "Using -p"
        shift
        ;;
      # ADD BK, REMOVE BK, CUSTOM DIR
      -a|-r|-c)
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
          echo "Using -a"
        elif [[ "$current_flag" == "-r" ]]; then
          # Handle removing a label:directory based on label
          echo "Using -r"
        elif [[ "$current_flag" == "-c" ]]; then
          # Handle using fcd with a custom dir
          echo "Using -c"
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
