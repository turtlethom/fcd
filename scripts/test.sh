#!/usr/bin/env bash
# Test whether FCD has been completely uninstalled
set -e

echo "[*] Running FCD uninstall test..."

CLEAN=true

# -----------------------------
# Paths
# -----------------------------
BIN="$HOME/.local/bin/fcd"
SHARE_DIR="$HOME/.local/share/fcd"
CONFIG_DIR="$HOME/.config/fcd"

# RC files to check for wrapper and completion lines
RC_FILES=(
  "$HOME/.bashrc"
  "$HOME/.bash_profile"
  "$HOME/.zshrc"
  "$HOME/.profile"
  "$HOME/.config/fish/config.fish"
)

# Per-user completion scripts
COMPLETIONS=(
  "$HOME/.local/share/fcd/fcd.bash"  # Bash
  "$HOME/.zsh/completions/_fcd"      # Zsh
  "$HOME/.config/fish/completions/fcd.fish" # Fish
)

# -----------------------------
# Check binary
# -----------------------------
if [ -f "$BIN" ]; then
  echo "[*] Binary still exists: $BIN"
  CLEAN=false
else
  echo "[ ] Binary removed"
fi

# -----------------------------
# Check shell wrappers
# -----------------------------
for wrapper in "$SHARE_DIR/fcd.sh" "$SHARE_DIR/fcd.fish"; do
  if [ -f "$wrapper" ]; then
    echo "[*] Wrapper still exists: $wrapper"
    CLEAN=false
  else
    echo "[ ] Wrapper removed: $wrapper"
  fi
done

# -----------------------------
# Check RC files for wrapper or completion lines
# -----------------------------
for rc in "${RC_FILES[@]}"; do
  if [ -f "$rc" ]; then
    if grep -qF "fcd.sh" "$rc"; then
      echo "[*] Bash/Zsh wrapper line still present in $rc"
      CLEAN=false
    fi
    if grep -qF "fcd.fish" "$rc"; then
      echo "[*] Fish wrapper line still present in $rc"
      CLEAN=false
    fi
    if grep -qF "fcd.bash" "$rc"; then
      echo "[*] Bash completion line still present in $rc"
      CLEAN=false
    fi
    if grep -qF "_fcd" "$rc"; then
      echo "[*] Zsh completion line still present in $rc"
      CLEAN=false
    fi
  fi
done

# -----------------------------
# Check per-user completion scripts
# -----------------------------
for comp in "${COMPLETIONS[@]}"; do
  if [ -f "$comp" ]; then
    echo "[*] Completion script still exists: $comp"
    CLEAN=false
  else
    echo "[ ] Completion script removed: $comp"
  fi
done

# -----------------------------
# Check if share directory exists and is empty
# -----------------------------
if [ -d "$SHARE_DIR" ] && [ -n "$(ls -A "$SHARE_DIR")" ]; then
  echo "[*] Share directory not empty: $SHARE_DIR"
  CLEAN=false
else
  echo "[ ] Share directory cleaned up"
fi

# -----------------------------
# Check config directory
# -----------------------------
if [ -d "$CONFIG_DIR" ]; then
  echo "[*] Config directory still exists: $CONFIG_DIR"
  CLEAN=false
else
  echo "[ ] Config directory removed"
fi

# -----------------------------
# Final report
# -----------------------------
if [ "$CLEAN" = true ]; then
  echo "[ ] FCD fully uninstalled: all files removed"
else
  echo "[*] FCD not fully removed"
fi
