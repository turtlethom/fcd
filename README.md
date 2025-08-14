## 🚀 fcd — Fast Change Directory  
A **visual dropdown menu** (powered by [`fzf`](https://github.com/junegunn/fzf)) for jumping into your bookmarked directories **blazingly fast**.  
### 📝 Note
fcd currently supports **only directories located inside your `$HOME`**. Bookmarks outside `$HOME` are not allowed.
---

### 📖 Usage

| Command | Description |
|---------|-------------|
| `fcd` | Launch the `fzf` menu to explore and jump to bookmarked directories. |
| `fcd -b <LABEL>` | **(IN PROGRESS)** Jump directly to a bookmarked directory by its label. |
| `fcd -a <PATH>` or `fcd -a <LABEL:PATH>` | Add a new bookmark. If only a path is given, the label is automatically set to the folder name. All paths must be prefixed with `~/`|
| `fcd -r <LABEL>` | Remove a bookmark by its label. |
| `fcd -c` | Clear **all** bookmarks. |
| `fcd -p` | Print all stored bookmarks in `LABEL:PATH` format. |

---

### 💡 Examples

```bash
fcd -a ~/projects/myapp
# → Adds a bookmark to '~/projects/myapp' labeled "MYAPP"

fcd -a "WORK:~/work"
# → Adds a bookmark to '~/work' labeled "WORK"

fcd -p
# → Lists all bookmarks
```
---
### 🛠 Compatibility Status
fcd is fully functional on modern Linux distributions with Bash ≥ 4.x.  
Support for **macOS (default Bash 3.2)**, **older Bash versions**, and **zsh** is currently in progress.
