#!/usr/bin/env bash
# ---------------------------------------------------------
# FCD UNINSTALL SCRIPT
# Safely removes binary, wrappers, completions, PATH lines,
# Zsh fpath entries, config directory, and generated RC lines.
# Idempotent (safe to run multiple times).
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
rm -f "$SHARE_DIR/fcd.sh" "$SHARE_DIR/fcd.fish" && echo "[✓] Removed shell wrappers"
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

# Lines added by install.sh
WRAPPER_COMMENT="# fcd wrapper (generated)"
PATH_COMMENT="# Add ~/.local/bin to PATH for FCD (generated)"

WRAPPER_BASH="source $HOME/.local/share/fcd/fcd.sh"
WRAPPER_FISH="source $HOME/.local/share/fcd/fcd.fish"

COMPLETION_BASH="source $HOME/.local/share/fcd/fcd.bash"
ZSH_FPATH_LINE="fpath+=$HOME/.zsh/completions"

# Zsh-only lines install.sh might add
ZSH_AUTLOAD_LINE="autoload -Uz compinit"
ZSH_COMPINIT_LINE="compinit"

# ---------------------------------------------------------
# Remove wrapper + completion + PATH lines from each RC file
# ---------------------------------------------------------
for rc in "${RC_FILES[@]}"; do
  if [ -f "$rc" ]; then

    sed -i.bak "\|$WRAPPER_COMMENT|d" "$rc"
    sed -i.bak "\|$WRAPPER_BASH|d" "$rc"
    sed -i.bak "\|$WRAPPER_FISH|d" "$rc"

    sed -i.bak "\|$COMPLETION_BASH|d" "$rc"
    sed -i.bak "\|$ZSH_FPATH_LINE|d" "$rc"

    # Remove PATH line added by install.sh
    sed -i.bak "\|$PATH_COMMENT|d" "$rc"
    sed -i.bak "\|export PATH=\$HOME/.local/bin:\$PATH|d" "$rc"

    # Remove zsh compinit helper lines if we added them
    sed -i.bak "\|$ZSH_AUTLOAD_LINE|d" "$rc"
    sed -i.bak "\|$ZSH_COMPINIT_LINE|d" "$rc"

    echo "[*] Cleaned RC file: $rc"
  fi
done
echo "---"

# ---------------------------------------------------------
# Remove completion scripts
# ---------------------------------------------------------
rm -f "$SHARE_DIR/fcd.bash" \
  "$HOME/.zsh/completions/_fcd" \
  "$HOME/.config/fish/completions/fcd.fish" &&
  echo "[✓] Removed completion scripts"
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

echo "--------------------------------------------------"
echo "[✓] FCD completely uninstalled!"
echo "--------------------------------------------------"
