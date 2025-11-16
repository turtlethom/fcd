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

    # Remove wrapper block (if present)
    sed -i.bak "/# fcd wrapper function for 'fcd' command (generated)/,+1d" "$rc" || true

    # Remove any completion sourcing lines
    sed -i.bak '/source .*fcd\.bash/d' "$rc" || true
    sed -i.bak '/source .*fcd\.sh/d' "$rc" || true
    sed -i.bak '/fcd\.fish/d' "$rc" || true

    echo "[*] Cleaned FCD lines from $rc (backup in $rc.bak)"
  fi
done

# -----------------------------
# Remove user-level completion scripts
# -----------------------------
echo "[*] Removing completion scripts..."

# Bash
BASH_COMPLETION="$HOME/.local/share/fcd/fcd.bash"
if [ -f "$BASH_COMPLETION" ]; then
  rm -f "$BASH_COMPLETION"
  echo "[*] Removed Bash completion: $BASH_COMPLETION"
fi

# Zsh
ZSH_COMPLETION="$HOME/.zsh/completions/_fcd"
if [ -f "$ZSH_COMPLETION" ]; then
  rm -f "$ZSH_COMPLETION"
  echo "[*] Removed Zsh completion: $ZSH_COMPLETION"
fi

# Fish
FISH_COMPLETION="$HOME/.config/fish/completions/fcd.fish"
if [ -f "$FISH_COMPLETION" ]; then
  rm -f "$FISH_COMPLETION"
  echo "[*] Removed Fish completion: $FISH_COMPLETION"
fi

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
