#!/usr/bin/env bash
# Test whether FCD has been completely uninstalled
set -e

echo "[*] Running FCD uninstall test..."

CLEAN=true

# Check binary
if [ -f "$HOME/.local/bin/fcd" ]; then
    echo "[*] Binary still exists: ~/.local/bin/fcd"
    CLEAN=false
else
    echo "[ ] Binary removed"
fi

# Check share folder wrappers
if [ -f "$HOME/.local/share/fcd/fcd.sh" ]; then
    echo "[*] Bash/Zsh wrapper still exists: ~/.local/share/fcd/fcd.sh"
    CLEAN=false
else
    echo "[ ] Bash/Zsh wrapper removed"
fi

if [ -f "$HOME/.local/share/fcd/fcd.fish" ]; then
    echo "[*] Fish wrapper still exists: ~/.local/share/fcd/fcd.fish"
    CLEAN=false
else
    echo "[ ] Fish wrapper removed"
fi

# Check RC files for source line
RC_FILES=("$HOME/.bashrc" "$HOME/.bash_profile" "$HOME/.zshrc" "$HOME/.profile" "$HOME/.config/fish/config.fish")

for rc in "${RC_FILES[@]}"; do
    if [ -f "$rc" ] && grep -qF "source $HOME/.local/share/fcd/fcd" "$rc"; then
        echo "[*] Source line still present in $rc"
        CLEAN=false
    fi
done

# Final report
if [ "$CLEAN" = true ]; then
    echo "[ ] FCD fully uninstalled: all files removed"
else
    echo "[*] FCD not fully removed"
fi
