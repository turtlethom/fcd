#!/usr/bin/env bash
# Test whether FCD has been completely uninstalled
set -e

echo "[*] Running FCD uninstall test..."

CLEAN=true
SHARE_DIR="$HOME/.local/share/fcd"
BINARY="$HOME/.local/bin/fcd"

# -----------------------------
# Check binary
# -----------------------------
if [ -f "$BINARY" ]; then
  echo "[*] Binary still exists: $BINARY"
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
# Check RC files for source lines
# -----------------------------
RC_FILES=(
  "$HOME/.bashrc"
  "$HOME/.bash_profile"
  "$HOME/.zshrc"
  "$HOME/.profile"
  "$HOME/.config/fish/config.fish"
)

for rc in "${RC_FILES[@]}"; do
  if [ -f "$rc" ]; then
    if grep -qF "fcd.sh" "$rc"; then
      echo "[*] Bash/Zsh source line still present in $rc"
      CLEAN=false
    fi
    if grep -qF "fcd.fish" "$rc"; then
      echo "[*] Fish source line still present in $rc"
      CLEAN=false
    fi
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
CONFIG_DIR="$HOME/.config/fcd"
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
