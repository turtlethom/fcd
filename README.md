## 🚀 fcd — Fast Change Directory  
A **shortcut dropdown menu** (powered by [`bubbletea`](https://github.com/charmbracelet/bubbletea)) for jumping into your bookmarked directories **blazingly fast**.  
### 📝 Note
fcd currently supports **only directories located inside your `$HOME`**. Bookmarks outside `$HOME` are not allowed.
---

### 📖 Usage

| Command | Description |
|---------|-------------|
| `fcd` | Launch the `bubbletea` menu to explore and jump to bookmarked directories. |
| `fcd branch <LABEL>` | Jump directly to a bookmarked directory by its label. |
| `fcd add <PATH>` or `fcd add <LABEL:PATH>` | Add a new bookmark. If only a path is given, the label is automatically set to the folder name. All paths must be prefixed with `~/`|
| `fcd remove <LABEL>` | Remove a bookmark by its label. |
| `fcd clear` | Clear **all** bookmarks. |
| `fcd print` | Print all stored bookmarks in `LABEL:PATH` format. |
| `fcd setcolor -p <color> -s <color> -t <color>` | Set colors for fcd menu (Primary, Secondary, Tertiary) |
| `fcd listcolors` | List all colors available for fcd theme |

---

### 💡 Examples

```bash
fcd
# → Opens interactive menu to saved shortcuts 

fcd add ~/projects/myapp
# → Adds a bookmark to '~/projects/myapp' labeled "MYAPP"

fcd add "JOB:~/work"
# → Adds a bookmark to '~/work' labeled "JOB"

fcd remove nvim
# → Removes a bookmark labeled "NVIM" (case insensitive) from user configuration

fcd branch BASHCONFIG
# → Changes current working directory to the path matching "BASHCONFIG" (case insensitive)

fcd clear
# → Clears all shortcuts from user configuration

fcd print
# → Lists all saved bookmarks within user configuration

fcd setcolor -p pink -s white -t black
# → Set theme color for primary, secondary, and tertiary colors

fcd listcolors
# → Lists all colors available to fcd menu
```
---
### 🛠 Compatibility Status
fcd is fully functional for the following shells:
- **bash**
- **zsh**
- **fish**

Support for **Windows PowerShell** is currently in progress.
