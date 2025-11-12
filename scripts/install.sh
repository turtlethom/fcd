#!/usr/bin/env bash
#<-- UNIVERSAL INSTALLATION SCRIPT FOR FCD -->#
set -e

echo "[*] Starting FCD installation..."

# -----------------------------
# Resolve directories
# -----------------------------
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_ROOT="$SCRIPT_DIR/.."
BUILD_DIR="$PROJECT_ROOT/build"

mkdir -p "$BUILD_DIR"

# -----------------------------
# Build fcd binary
# -----------------------------
echo "[*] Building fcd binary..."
go build -o "$BUILD_DIR/fcd" "$PROJECT_ROOT/main.go"

# -----------------------------
# Install binary to ~/.local/bin
# -----------------------------
echo "[*] Installing binary to ~/.local/bin..."
mkdir -p "$HOME/.local/bin"
cp "$BUILD_DIR/fcd" "$HOME/.local/bin/"

# -----------------------------
# Install shell wrappers
# -----------------------------
SHARE_DIR="$HOME/.local/share/fcd"
mkdir -p "$SHARE_DIR"

# Bash/Zsh wrapper
cat >"$SHARE_DIR/fcd.sh" <<'EOF'
fcd() {
  local output status
  output=$("$HOME/.local/bin/fcd" "$@")
  status=$?

  # if the binary fails, just print its output
  if [ $status -ne 0 ]; then
    echo "$output"
    return $status
  fi

  # only cd if output is a valid directory
  if [ -d "$output" ]; then
    cd "$output" || return
  elif [ -n "$output" ]; then
    echo "$output"
  fi
}
EOF

# Fish wrapper
cat >"$SHARE_DIR/fcd.fish" <<'EOF'
function fcd
    set output (~/.local/bin/fcd $argv)
    set status $status

    if test $status -ne 0
        echo $output
        return $status
    end

    if test -d "$output"
        cd $output
    else if test -n "$output"
        echo $output
    end
end
EOF

# -----------------------------
# Detect shell and RC file
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
# Add source line to RC file
# -----------------------------
if [ "$SHELL_NAME" = "fish" ]; then
	INSTALL_LINE="source $SHARE_DIR/fcd.fish"
else
	INSTALL_LINE="source $SHARE_DIR/fcd.sh"
fi

if ! grep -Fxq "$INSTALL_LINE" "$RC_FILE" 2>/dev/null; then
	echo "$INSTALL_LINE" >>"$RC_FILE"
	echo "[*] Added FCD wrapper to $RC_FILE"
else
	echo "[*] FCD wrapper already present in $RC_FILE"
fi

# -----------------------------
# Ensure ~/.local/bin is in PATH
# -----------------------------
PATH_LINE='export PATH=$HOME/.local/bin:$PATH'
if [ "$SHELL_NAME" != "fish" ]; then
	if ! grep -Fxq "$PATH_LINE" "$RC_FILE" 2>/dev/null; then
		echo "$PATH_LINE" >>"$RC_FILE"
		echo "[*] Added ~/.local/bin to PATH in $RC_FILE"
	fi
fi

# -----------------------------
# Done
# -----------------------------
echo "[*] Installation complete!"
echo "Restart your shell or run: source $RC_FILE"
echo "Then you can use: fcd"
