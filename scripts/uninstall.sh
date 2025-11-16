#!/usr/bin/env bash
set -e

echo "[*] Uninstalling FCD..."

BIN="$HOME/.local/bin/fcd"
SHARE_DIR="$HOME/.local/share/fcd"
CONFIG_DIR="$HOME/.config/fcd"

# -----------------------------
# Remove binary
# -----------------------------
rm -f "$BIN" && echo "[✓] Removed binary"

# -----------------------------
# Remove wrappers
# -----------------------------
rm -f "$SHARE_DIR/fcd.sh" "$SHARE_DIR/fcd.fish" && echo "[✓] Removed wrapper files"

# -----------------------------
# Clean RC files
# -----------------------------
RC_FILES=(
    "$HOME/.bashrc"
    "$HOME/.bash_profile"
    "$HOME/.zshrc"
    "$HOME/.profile"
    "$HOME/.config/fish/config.fish"
)

COMMENT="# fcd wrapper (generated)"
WRAPPER_BASH="$HOME/.local/share/fcd/fcd.sh"
WRAPPER_FISH="$HOME/.local/share/fcd/fcd.fish"

for rc in "${RC_FILES[@]}"; do
    if [ -f "$rc" ]; then
        # Use | delimiter to avoid slashes issues
        sed -i.bak "\|$COMMENT|d" "$rc" || true
        sed -i.bak "\|$WRAPPER_BASH|d" "$rc" || true
        sed -i.bak "\|$WRAPPER_FISH|d" "$rc" || true
        echo "[*] Cleaned FCD entries from $rc"
    fi
done

# -----------------------------
# Remove per-user completion scripts
# -----------------------------
rm -f "$HOME/.local/share/fcd/fcd.bash" \
      "$HOME/.zsh/completions/_fcd" \
      "$HOME/.config/fish/completions/fcd.fish" \
      2>/dev/null && echo "[✓] Removed completion scripts"

# -----------------------------
# Remove config directory
# -----------------------------
rm -rf "$CONFIG_DIR" && echo "[✓] Removed config directory"

# -----------------------------
# Remove empty share directory
# -----------------------------
if [ -d "$SHARE_DIR" ] && [ -z "$(ls -A "$SHARE_DIR")" ]; then
    rmdir "$SHARE_DIR"
    echo "[✓] Removed empty share directory"
fi

echo "[✓] Uninstalled successfully!"
