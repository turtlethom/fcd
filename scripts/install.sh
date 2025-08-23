#!/usr/bin/env bash
#<-- INSTALLATION SCRIPT -->#
set -e

# Build fcd program from source
echo "[*] Building fcd..."
go build -o build/fcd main.go

# Install binary locally
echo "[*] Installing binary to ~/.local/bin"
mkdir -p "$HOME/.local/bin"
cp build/fcd "$HOME/.local/bin/"

# Install shell wrapper in share folder
echo "[*] Installing shell wrapper..."
mkdir -p "$HOME/.local/share/fcd"
cat > "$HOME/.local/share/fcd/fcd.sh" <<'EOF'
fcd() {
  target=$("$HOME/.local/bin/fcd" "$@")
  if [ -n "$target" ]; then
    cd "$target" || return
  fi
}
EOF

# Detect shell
SHELL_NAME=$(basename "$SHELL")
RC_FILE=""

# Check which shell config to write to
case "$SHELL_NAME" in
  bash)
    # MAC OS
    if [[ "$OSTYPE" == "darwin"* ]]; then
      RC_FILE="$HOME/.bash_profile"
    # LINUX
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
echo "[*] Using rc file: $RC_FILE"

INSTALL_LINE="source $HOME/.local/share/fcd/fcd.sh"

# Fish uses a different syntax
if [ "$SHELL_NAME" = "fish" ]; then
  INSTALL_LINE="source $HOME/.local/share/fcd/fcd.sh; functions -c fcd fcd"
fi

# Add install line if missing from config
if ! grep -Fq "$INSTALL_LINE" "$RC_FILE" 2>/dev/null; then
  echo "$INSTALL_LINE" >> "$RC_FILE"
  echo "[*] Added fcd to $RC_FILE"
else
  echo "[*] fcd already in $RC_FILE"
fi

# Ensure ~/.local/bin is in PATH
if ! echo "$PATH" | grep -q "$HOME/.local/bin"; then
  echo "export PATH=\$HOME/.local/bin:\$PATH" >> "$RC_FILE"
  echo "[*] Added ~/.local/bin to PATH in $RC_FILE"
fi

echo "[*] Done! Restart your shell or run 'source $RC_FILE'"
