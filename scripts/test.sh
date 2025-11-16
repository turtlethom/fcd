#!/usr/bin/env bash
# -------------------------------------------------------------
# FCD Test Suite — Uninstall Verification + Completion Testing
# -------------------------------------------------------------
# This script performs:
#
# 1. Completion generation tests (bash, zsh, fish)
# 2. Original uninstall verification tests
#
# IMPORTANT:
# - Binary test checks for the installed binary in ~/.local/bin/fcd.
# - No modifications are made to the user's system.
# -------------------------------------------------------------

set -euo pipefail

echo "[*] Running FCD test suite..."

CLEAN=true

# =============================================================
# SECTION 1 — Completion Generation Tests
# =============================================================

echo
echo "--------------------------------------------------"
echo "[1] Completion Generation Tests"
echo "--------------------------------------------------"

INSTALL_BIN="$HOME/.local/bin/fcd"

if [[ -f "$INSTALL_BIN" ]]; then
  echo "[✓] Found installed binary → $INSTALL_BIN"

  # Test completion generation
  echo "[*] Testing Bash completion..."
  if "$INSTALL_BIN" completion bash >/dev/null 2>&1; then
    echo "[✓] Bash completion OK"
  else
    echo "[✗] Bash completion FAILED"
    CLEAN=false
  fi

  echo "[*] Testing Zsh completion..."
  if "$INSTALL_BIN" completion zsh >/dev/null 2>&1; then
    echo "[✓] Zsh completion OK"
  else
    echo "[✗] Zsh completion FAILED"
    CLEAN=false
  fi

  echo "[*] Testing Fish completion..."
  if "$INSTALL_BIN" completion fish >/dev/null 2>&1; then
    echo "[✓] Fish completion OK"
  else
    echo "[✗] Fish completion FAILED"
    CLEAN=false
  fi
else
  echo "[!] Installed binary not found — skipping completion tests"
fi

# =============================================================
# SECTION 2 — Uninstallation Verification Tests
# (unchanged from original test + adapted for new structure)
# =============================================================

echo
echo "--------------------------------------------------"
echo "[2] Uninstallation Verification Tests"
echo "--------------------------------------------------"

# -----------------------------
# Paths
# -----------------------------
BIN="$HOME/.local/bin/fcd"
SHARE_DIR="$HOME/.local/share/fcd"
CONFIG_DIR="$HOME/.config/fcd"

RC_FILES=(
  "$HOME/.bashrc"
  "$HOME/.bash_profile"
  "$HOME/.zshrc"
  "$HOME/.profile"
  "$HOME/.config/fish/config.fish"
)

COMPLETIONS=(
  "$HOME/.local/share/fcd/fcd.bash"
  "$HOME/.zsh/completions/_fcd"
  "$HOME/.config/fish/completions/fcd.fish"
)

WRAPPERS=(
  "$SHARE_DIR/fcd.sh"
  "$SHARE_DIR/fcd.fish"
)

# -----------------------------
# Check binary
# -----------------------------
echo "[*] Checking binary..."
if [[ -f "$BIN" ]]; then
  echo "[✗] Binary still exists: $BIN"
  CLEAN=false
else
  echo "[✓] Binary removed"
fi
echo "---"

# -----------------------------
# Check wrapper scripts
# -----------------------------
echo "[*] Checking wrapper scripts..."
for wrapper in "${WRAPPERS[@]}"; do
  if [[ -f "$wrapper" ]]; then
    echo "[✗] Wrapper still exists: $wrapper"
    CLEAN=false
  else
    echo "[✓] Wrapper removed: $wrapper"
  fi
done
echo "---"

# -----------------------------
# Check RC files for wrapper or completion lines
# -----------------------------
echo "[*] Checking rc files for leftover lines..."
for rc in "${RC_FILES[@]}"; do
  if [[ -f "$rc" ]]; then

    if grep -qF "fcd.sh" "$rc"; then
      echo "[✗] Bash/Zsh wrapper line still present in $rc"
      CLEAN=false
    fi
    if grep -qF "fcd.fish" "$rc"; then
      echo "[✗] Fish wrapper line still present in $rc"
      CLEAN=false
    fi
    if grep -qF "fcd.bash" "$rc"; then
      echo "[✗] Bash completion line still present in $rc"
      CLEAN=false
    fi
    if grep -qF "_fcd" "$rc"; then
      echo "[✗] Zsh completion line still present in $rc"
      CLEAN=false
    fi

  fi
done
echo "---"

# -----------------------------
# Check per-user completion scripts
# -----------------------------
echo "[*] Checking completion script files..."
for comp in "${COMPLETIONS[@]}"; do
  if [[ -f "$comp" ]]; then
    echo "[✗] Completion script exists: $comp"
    CLEAN=false
  else
    echo "[✓] Completion removed: $comp"
  fi
done
echo "---"

# -----------------------------
# Check share directory
# -----------------------------
echo "[*] Checking shared directory cleanup..."
if [[ -d "$SHARE_DIR" && -n "$(ls -A "$SHARE_DIR")" ]]; then
  echo "[✗] Share directory not empty: $SHARE_DIR"
  CLEAN=false
else
  echo "[✓] Share directory fully removed or empty"
fi
echo "---"

# -----------------------------
# Check config directory
# -----------------------------
echo "[*] Checking configuration directory..."
if [[ -d "$CONFIG_DIR" ]]; then
  echo "[✗] Config directory still exists: $CONFIG_DIR"
  CLEAN=false
else
  echo "[✓] Config directory removed"
fi
echo "---"

# =============================================================
# Final report
# =============================================================
echo
echo "--------------------------------------------------"
if [[ "$CLEAN" == true ]]; then
  echo "[✓] FCD fully uninstalled — all tests passed"
else
  echo "[✗] FCD not fully removed — see details above"
fi
echo "--------------------------------------------------"
