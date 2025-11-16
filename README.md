# 🚀 fcd — Fast Change Directory

**fcd** is a **shortcut dropdown menu** for jumping into your saved directories **blazingly fast**, powered by:  
- [`bubbletea`](https://github.com/charmbracelet/bubbletea)  
- [`cobra`](https://github.com/spf13/cobra)  
- Written in **Golang**, works seamlessly in **bash**, **zsh**, and **fish**.  

---

## 📝 Note
fcd currently supports **only directories located inside your `$HOME`**. Shortcuts outside `$HOME` are **not allowed**.

---

## 📖 Usage

You can use `fcd` in **interactive mode**, or via specific **commands**:

### 🔹 Interactive
- `fcd` — Opens the interactive menu to explore and jump to saved shortcuts.

### 📌 Bookmark/Shortcut Management
- `fcd add <shortcut_path>` — Add a directory to saved shortcuts.  
- `fcd add <shortcut_label:shortcut_path>` — Add a directory with a custom label.  
- `fcd remove <label>` — Remove a bookmark by label (case insensitive).  
- `fcd print` — List all saved bookmarks in `LABEL:PATH` format.  
- `fcd clear` — Remove **all** saved shortcuts.  

### 🏃 Navigation
- `fcd branch <label>` — Jump to a saved directory by its label.  

### 🎨 Colors & Themes
- `fcd colors set -p <primary> -s <secondary> -t <tertiary>` — Set theme colors for fcd menu.  
- `fcd colors list` — List all available colors.  

### ⚙️ System & Help
- `fcd completion` — Generate shell auto-completion scripts.  
- `fcd help` — Show help for any command.  

---

## 💡 Examples

```bash
# Open interactive menu
fcd

# Add bookmarks
fcd add ~/projects/myapp
fcd add "JOB:~/workspace"

# Remove bookmark
fcd remove nvim

# Jump to saved directory
fcd branch personalconfig

# Clear all shortcuts
fcd clear

# Print all bookmarks
fcd print

# Manage colors
fcd colors set -p pink -s white -t black
fcd colors list
### 🛠 Compatibility Status
fcd is fully functional for the following shells:
- **bash**
- **zsh**
- **fish**

Support for **Windows PowerShell** is currently in progress.
