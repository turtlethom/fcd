#!/usr/bin/env bash
# ---------------------------------------------------------
# FCD UNINSTALL SCRIPT
# Removes binary, wrappers, completions, config, and
# cleans shell RC files — safe and idempotent.
# ---------------------------------------------------------
set -e

echo "--------------------------------------------------"
echo "[*] Uninstalling FCD..."
echo "--------------------------------------------------"

BIN="$HOME/.local/bin/fcd"
SHARE_DIR="$HOME/.local/share/fcd"
CONFIG_DIR="$HOME/.config/fcd"

# ---------------------------------------------------------
# Remove binary
# ---------------------------------------------------------
rm -f "$BIN" && echo "[✓] Removed binary"
echo "---"

# ---------------------------------------------------------
# Remove wrapper files
# ---------------------------------------------------------
rm -f "$SHARE_DIR/fcd.sh" "$SHARE_DIR/fcd.fish" && echo "[✓] Removed wrapper files"
echo "---"

# ---------------------------------------------------------
# RC files to clean
# ---------------------------------------------------------
RC_FILES=(
    "$HOME/.bashrc"
    "$HOME/.bash_profile"
    "$HOME/.zshrc"
    "$HOME/.profile"
    "$HOME/.config/fish/config.fish"
)

WRAPPER_COMMENT="# fcd wrapper (generated)"
WRAPPER_BASH="source $HOME/.local/share/fcd/fcd.sh"
WRAPPER_FISH="source $HOME/.local/share/fcd/fcd.fish"

COMPLETION_BASH="source $HOME/.local/share/fcd/fcd.bash"
COMPLETION_ZSH="source $HOME/.zsh/completions/_fcd"

# ---------------------------------------------------------
# Remove wrapper + completion lines
# ---------------------------------------------------------
for rc in "${RC_FILES[@]}"; do
    if [ -f "$rc" ]; then
        sed -i.bak "\|$WRAPPER_COMMENT|d" "$rc"
        sed -i.bak "\|$WRAPPER_BASH|d" "$rc"
        sed -i.bak "\|$WRAPPER_FISH|d" "$rc"
        sed -i.bak "\|$COMPLETION_BASH|d" "$rc"
        sed -i.bak "\|$COMPLETION_ZSH|d" "$rc"
        sed -i.bak "\|fpath+=$HOME/.zsh/completions|d" "$rc"
        echo "[*] Cleaned RC entries from $rc"
    fi
done
echo "---"

# ---------------------------------------------------------
# Remove completion files
# ---------------------------------------------------------
rm -f "$SHARE_DIR/fcd.bash" \
      "$HOME/.zsh/completions/_fcd" \
      "$HOME/.config/fish/completions/fcd.fish" \
      && echo "[✓] Removed completion scripts"
echo "---"

# ---------------------------------------------------------
# Remove config directory
# ---------------------------------------------------------
rm -rf "$CONFIG_DIR" && echo "[✓] Removed config directory"
echo "---"

# ---------------------------------------------------------
# Remove empty share directory
# ---------------------------------------------------------
if [ -d "$SHARE_DIR" ] && [ -z "$(ls -A "$SHARE_DIR")" ]; then
    rmdir "$SHARE_DIR"
    echo "[✓] Removed empty share directory"
fi
echo "---"

echo "[✓] FCD uninstalled successfully!"
