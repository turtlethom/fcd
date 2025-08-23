#!/usr/bin/env bash
#<-- UNIVERSAL UNINSTALL SCRIPT FOR FCD -->#
set -e

echo "[*] Starting FCD uninstallation..."

# -----------------------------
# Determine shell and RC file
# -----------------------------
SHELL_NAME=$(basename "$SHELL")
RC_FILE=""

case "$SHELL_NAME" in
  bash)
    if [[ "$OSTYPE" == "darwin"* ]]; then
      RC_FILE="$HOME/.bash_profile"
    else
      RC_FILE="$HOME/.bashrc"
    fi
    ;;
  zsh)
    RC_FILE="$HOME/.zshrc"
    ;;
  fish)
    RC_FILE="$HOME/.config/fish/config.fish"
    ;;
  *)
    RC_FILE="$HOME/.profile"
    ;;
esac

echo "[*] Detected shell: $SHELL_NAME"
echo "[*] Using RC file: $RC_FILE"

# -----------------------------
# Remove binary
# -----------------------------
BINARY="$HOME/.local/bin/fcd"
if [ -f "$BINARY" ]; then
    rm -f "$BINARY"
    echo "[*] Removed binary: $BINARY"
else
    echo "[*] Binary not found, skipping."
fi

# -----------------------------
# Remove shell wrapper
# -----------------------------
SHARE_DIR="$HOME/.local/share/fcd"

if [ "$SHELL_NAME" = "fish" ]; then
    WRAPPER="$SHARE_DIR/fcd.fish"
else
    WRAPPER="$SHARE_DIR/fcd.sh"
fi

if [ -f "$WRAPPER" ]; then
    rm -f "$WRAPPER"
    echo "[*] Removed wrapper: $WRAPPER"
else
    echo "[*] Wrapper not found, skipping."
fi

# -----------------------------
# Remove source line from RC file
# -----------------------------
if [ -f "$RC_FILE" ]; then
    if [ "$SHELL_NAME" = "fish" ]; then
        INSTALL_LINE="source $SHARE_DIR/fcd.fish"
    else
        INSTALL_LINE="source $SHARE_DIR/fcd.sh"
    fi

    # Remove exact matching line
    sed -i.bak "/$(echo "$INSTALL_LINE" | sed 's/[\/&]/\\&/g')/d" "$RC_FILE"
    echo "[*] Removed source line from $RC_FILE (backup saved as $RC_FILE.bak)"
else
    echo "[*] RC file not found, skipping."
fi

# -----------------------------
# Optional: remove PATH line
# -----------------------------
if [ "$SHELL_NAME" != "fish" ] && [ -f "$RC_FILE" ]; then
    PATH_LINE='export PATH=$HOME/.local/bin:$PATH'
    sed -i.bak "/$(echo "$PATH_LINE" | sed 's/[\/&]/\\&/g')/d" "$RC_FILE"
    echo "[*] Removed PATH line from $RC_FILE (backup saved as $RC_FILE.bak)"
fi

# -----------------------------
# Cleanup share directory if empty
# -----------------------------
if [ -d "$SHARE_DIR" ] && [ -z "$(ls -A "$SHARE_DIR")" ]; then
    rmdir "$SHARE_DIR"
    echo "[*] Removed empty share directory: $SHARE_DIR"
fi

echo "[*] FCD uninstallation complete!"
echo "Restart your shell to apply changes."
