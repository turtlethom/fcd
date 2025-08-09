#!/usr/bin/env bash

fcd_help() {
  printf "fcd - fast 'cd' with fuzzy finder\n\n"
  printf "Usage:\n"
  printf "  fcd -a ['LABEL/\$HOME/PATH']\n" 
  printf "    *) Add a bookmarked directory.\n"
  printf "    *) Bookmarks a LABEL with a PATH under your home directory.\n"
  printf "  fcd -r [LABEL]\n"
  printf "    *) Remove a bookmarked directory.\n"
  printf "    *) Deletes the bookmark identified by LABEL.\n"
  printf "  fcd -c\n"                  
  printf "    *) Clear all bookmarked directories.\n"
  printf "  fcd -p\n" 
  printf "    *) Lists all saved bookmarks.\n"
  printf "  fcd -h\n"                  
  printf "    *) Show this help message.\n"
}
