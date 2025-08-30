## 🚀 fcd — Fast Change Directory  
A **shortcut dropdown menu** (powered by [`bubbletea`](https://github.com/charmbracelet/bubbletea)) for jumping into your bookmarked directories **blazingly fast**.  
### 📝 Note
fcd currently supports **only directories located inside your `$HOME`**. Bookmarks outside `$HOME` are not allowed.
---

### 📖 Usage

| Command | Description |
|---------|-------------|
| `fcd` | Launch the `fzf` menu to explore and jump to bookmarked directories. |
| `fcd branch <LABEL>` | **(IN PROGRESS)** Jump directly to a bookmarked directory by its label. |
| `fcd add <PATH>` or `fcd add <LABEL:PATH>` | Add a new bookmark. If only a path is given, the label is automatically set to the folder name. All paths must be prefixed with `~/`|
| `fcd remove <LABEL>` | Remove a bookmark by its label. |
| `fcd clear` | Clear **all** bookmarks. |
| `fcd print` | Print all stored bookmarks in `LABEL:PATH` format. |

---

### 💡 Examples

```bash
fcd add ~/projects/myapp
# → Adds a bookmark to '~/projects/myapp' labeled "MYAPP"

fcd add "JOB:~/work"
# → Adds a bookmark to '~/work' labeled "JOB"

fcd remove nvim
# → Removes a bookmark labeled "NVIM" (case insensitive)

fcd clear
# → Clears all shortcuts from user configuration

fcd print
# → Lists all bookmarks
```
---
### 🛠 Compatibility Status
fcd is fully functional for the following shells:
- **bash**
- **zsh**
- **fish**
Support for **Windows PowerShell** is currently in progress.
