## 🚀 fcd — Fast Change Directory  
A **visual dropdown menu** (powered by [`fzf`](https://github.com/junegunn/fzf)) for jumping into your bookmarked directories **blazingly fast**.  

---

### 📖 Usage

| Command | Description |
|---------|-------------|
| `fcd` | Launch the `fzf` menu to explore and jump to bookmarked directories. |
| `fcd -b <LABEL>` | **(IN PROGRESS)** Jump directly to a bookmarked directory by its label. |
| `fcd -a <PATH>` or `fcd -a <LABEL:PATH>` | Add a new bookmark. If only a path is given, the label is automatically set to the folder name. |
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
