#!/usr/bin/env bash
set -e

echo "[*] Starting FCD uninstallation..."

SHARE_DIR="$HOME/.local/share/fcd"
BINARY="$HOME/.local/bin/fcd"

# -----------------------------
# Remove binary
# -----------------------------
if [ -f "$BINARY" ]; then
  rm -f "$BINARY"
  echo "[*] Removed binary: $BINARY"
else
  echo "[*] Binary not found, skipping."
fi

# -----------------------------
# Remove all shell wrappers
# -----------------------------
for wrapper in "$SHARE_DIR/fcd.sh" "$SHARE_DIR/fcd.fish"; do
  if [ -f "$wrapper" ]; then
    rm -f "$wrapper"
    echo "[*] Removed wrapper: $wrapper"
  fi
done

# -----------------------------
# Remove source lines from all RC files
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
    # Remove lines containing 'fcd.sh' or 'fcd.fish'
    sed -i.bak '/fcd\.sh/d' "$rc" || true
    sed -i.bak '/fcd\.fish/d' "$rc" || true
    echo "[*] Removed FCD source lines from $rc (backup saved as $rc.bak)"
  fi
done

# -----------------------------
# Cleanup share directory if empty
# -----------------------------
if [ -d "$SHARE_DIR" ] && [ -z "$(ls -A "$SHARE_DIR")" ]; then
  rmdir "$SHARE_DIR"
  echo "[*] Removed empty share directory: $SHARE_DIR"
fi

# -----------------------------
# Remove config directory (~/.config/fcd)
# -----------------------------
CONFIG_DIR="$HOME/.config/fcd"
if [ -d "$CONFIG_DIR" ]; then
  rm -rf "$CONFIG_DIR"
  echo "[*] Removed config directory: $CONFIG_DIR"
else
  echo "[ ] Config directory not found, skipping: $CONFIG_DIR"
fi

echo "[*] FCD uninstallation complete!"
echo "Restart your shell to apply changes."
