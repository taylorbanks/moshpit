<div align="center">
  <img src="./docs/logo.png" alt="moshpit logo" width="600" height="600"/>
</div>

---

moshpit is a terminal-based, interactive SSH/Mosh manager inspired by tools like lazydocker and k9s — but built for managing your fleet of servers directly from your terminal.
<br/>
With moshpit, you can quickly navigate, connect, manage, and transfer files between your local machine and any server defined in your `~/.ssh/config`. Toggle seamlessly between SSH and Mosh protocols for optimal connectivity. No more remembering IP addresses or running long scp commands — just a clean, keyboard-driven UI with powerful protocol flexibility.

---

## ✨ Features

### Server Management
- 📜 Read & display servers from your `~/.ssh/config` in a scrollable list.
- ➕ Add a new server from the UI with comprehensive SSH configuration options.
- ✏ Edit existing server entries directly from the UI with a tabbed interface.
- 🗑 Delete server entries safely.
- 📌 Pin / unpin servers to keep favorites at the top.
- 🏓 Ping server to check status.

### Protocol Flexibility
- 🔄 **Toggle between SSH and Mosh** on a per-host basis (press `m`)
- 🤘 **Mosh support** for reliable roaming connections across unstable networks
- 📊 **Visual protocol indicators** - see at a glance which servers use mosh
- ⚡ **Automatic fallback** - seamlessly falls back to SSH if mosh is unavailable
- 🏷 **Bulk protocol toggle** - set protocol for all servers with a specific tag (press `Shift+M`)
- ✅ **Cross-platform mosh detection** - automatically detects mosh availability

### Quick Server Navigation
- 🔍 Fuzzy search by alias, IP, or tags.
- 🖥 One‑keypress connection to the selected server (Enter) - uses SSH or Mosh based on preference.
- 🏷 Tag servers (e.g., prod, dev, test) for quick filtering.
- ↕️ Sort by alias or last connection (toggle + reverse).

### Advanced SSH Configuration
- 🔗 Port forwarding (LocalForward, RemoteForward, DynamicForward).
- 🤘 Connection multiplexing for faster subsequent connections.
- 🔐 Advanced authentication options (public key, password, agent forwarding).
- 🔒 Security settings (ciphers, MACs, key exchange algorithms).
- 🌐 Proxy settings (ProxyJump, ProxyCommand).
- ⚙️ Extensive SSH config options organized in tabbed interface.

### Key Management
- 🔑 SSH key autocomplete with automatic detection of available keys.
- 📝 Smart key selection with support for multiple keys.


### Upcoming
- 📁 Copy files between local and servers with an easy picker UI.
- 🔑 SSH Key Deployment Features:
    - Use default local public key (`~/.ssh/id_ed25519.pub` or `~/.ssh/id_rsa.pub`)
    - Paste custom public keys manually
    - Generate new keypairs and deploy them
    - Automatically append keys to `~/.ssh/authorized_keys` with correct permissions
---

## 🔐 Security Notice

moshpit does not introduce any new security risks.
It is simply a UI/TUI wrapper around your existing `~/.ssh/config` file.

- All SSH connections are executed through your system's native ssh binary (OpenSSH).

- Mosh connections are executed through your system's mosh binary when available.

- Private keys, passwords, and credentials are never stored, transmitted, or modified by moshpit.

- Your existing IdentityFile paths and ssh-agent integrations work exactly as before.

- moshpit only reads and updates your `~/.ssh/config`. A backup of the file is created automatically before any changes.

- Protocol preferences (SSH vs Mosh) are stored separately in `~/.moshpit/metadata.json` and never affect your SSH config.

- File permissions on your SSH config are preserved to ensure security.


## 🛡️ Config Safety: Non‑destructive writes and backups

- **Non‑destructive edits**: moshpit only writes the minimal required changes to your ~/.ssh/config. It uses a parser that preserves existing comments, spacing, order, and any settings it didn't touch. Your handcrafted comments and formatting remain intact.
- **Atomic writes**: Updates are written to a temporary file and then atomically renamed over the original, minimizing the risk of partial writes.
- **Protocol preferences**: Mosh/SSH protocol preferences are stored separately in `~/.moshpit/metadata.json` and never modify your SSH config.
- **Backups**:
  - One‑time original backup: Before moshpit makes its first change, it creates a single snapshot named config.original.backup beside your SSH config. If this file is present, it will never be recreated or overwritten.
  - Rolling backups: On every subsequent save, moshpit also creates a timestamped backup named like: ~/.ssh/config-<timestamp>-moshpit.backup. The app keeps at most 10 of these backups, automatically removing the oldest ones.
- **Migration**: Automatically migrates existing data from `~/.lazyssh/` or `~/.lazymosh/` to `~/.moshpit/` on first run.

## 📷 Screenshots

> **Note**: Screenshots will be updated to reflect the latest moshpit branding and mosh protocol features.

<div align="center">

### 🤘 Startup
<img src="./docs/loader.png" alt="App starting splash/loader" width="800" />

Clean loading screen when launching the app

---

### 📋 Server Management Dashboard
<img src="./docs/list server.png" alt="Server list view" width="900" />

Main dashboard displaying all configured servers with status indicators, pinned favorites at the top, and easy navigation

---

### 🔎 Search
<img src="./docs/search.png" alt="Fuzzy search servers" width="900" />

Fuzzy search functionality to quickly find servers by name, IP address, or tags

---

### ➕ Add/Edit Server
<img src="./docs/add server.png" alt="Add a new server" width="900" />

Tabbed interface for managing SSH connections with extensive configuration options organized into:
- **Basic** - Host, user, port, keys, tags
- **Connection** - Proxy, timeouts, multiplexing, canonicalization
- **Forwarding** - Port forwarding, X11, agent
- **Authentication** - Keys, passwords, methods, algorithm settings
- **Advanced** - Security, cryptography, environment, debugging

---

### 🔐 Connect to server
<img src="./docs/ssh.png" alt="SSH connection details" width="900" />

SSH into the selected server

</div>

---

## 📦 Installation

### Option 1: Homebrew (macOS)

```bash
brew install taylorbanks/homebrew-tap/moshpit
```

### Option 2: Download Binary from Releases

Download from [GitHub Releases](https://github.com/taylorbanks/moshpit/releases). You can use the snippet below to automatically fetch the latest version for your OS/ARCH (Darwin/Linux and amd64/arm64 supported):

```bash
# Detect latest version
LATEST_TAG=$(curl -fsSL https://api.github.com/repos/taylorbanks/moshpit/releases/latest | jq -r .tag_name)
# Download the correct binary for your system
curl -LJO "https://github.com/taylorbanks/moshpit/releases/download/${LATEST_TAG}/moshpit_$(uname)_$(uname -m).tar.gz"
# Extract the binary
tar -xzf moshpit_$(uname)_$(uname -m).tar.gz
# Move to /usr/local/bin or another directory in your PATH
sudo mv moshpit /usr/local/bin/
# enjoy!
moshpit
```

### Option 3: Build from Source

```bash
# Clone the repository
git clone https://github.com/taylorbanks/moshpit.git
cd moshpit

# Build
go build -o moshpit cmd/main.go

# Run it
./moshpit
```

### Requirements

- **SSH**: OpenSSH (pre-installed on most systems)
- **Mosh** (optional): Install for enhanced roaming connection support
  - macOS: `brew install mosh`
  - Ubuntu/Debian: `sudo apt install mosh`
  - Fedora: `sudo dnf install mosh`
  - Arch: `sudo pacman -S mosh`

---

## ⌨️ Key Bindings

| Key   | Action                              |
| ----- | ----------------------------------- |
| /     | Toggle search bar                   |
| ↑↓/jk | Navigate servers                    |
| Enter | Connect to selected server          |
| m     | Toggle SSH/Mosh protocol            |
| M     | Bulk toggle protocol by tag         |
| f     | Port forward                        |
| x     | Stop forwarding                     |
| c     | Copy SSH command to clipboard       |
| g     | Ping selected server                |
| r     | Refresh background data             |
| a     | Add server                          |
| e     | Edit server                         |
| t     | Edit tags                           |
| d     | Delete server                       |
| p     | Pin/Unpin server                    |
| s     | Toggle sort field                   |
| S     | Reverse sort order                  |
| q     | Quit                                |

**In Server Form:**
| Key    | Action               |
| ------ | -------------------- |
| Ctrl+H | Previous tab         |
| Ctrl+L | Next tab             |
| Ctrl+S | Save                 |
| Esc    | Cancel               |

Tip: The hint bar at the top of the list shows the most useful shortcuts.

---

## 🤝 Contributing

Contributions are welcome!

- If you spot a bug or have a feature request, please [open an issue](https://github.com/taylorbanks/moshpit/issues).
- If you'd like to contribute, fork the repo and submit a pull request ❤️.

We love seeing the community make moshpit better 🤘

### Semantic Pull Requests

This repository enforces semantic PR titles via an automated GitHub Action. Please format your PR title as:

- type(scope): short descriptive subject
Notes:
- Scope is optional and should be one of: ui, cli, config, parser.

Allowed types in this repo:
- feat: a new feature
- fix: a bug fix
- improve: quality or UX improvements that are not a refactor or perf
- refactor: code change that neither fixes a bug nor adds a feature
- docs: documentation only changes
- test: adding or refactoring tests
- ci: CI/CD or automation changes
- chore: maintenance tasks, dependency bumps, non-code infra
- revert: reverts a previous commit

Examples:
- feat(ui): add server pinning and sorting options
- fix(parser): handle comments at end of Host blocks
- improve(cli): show friendly error when ssh binary missing
- refactor(config): simplify backup rotation logic
- docs: add installation instructions for Homebrew
- ci: cache Go toolchain and dependencies

Tip: If your PR touches multiple areas, pick the most relevant scope or omit the scope.

---

## ⭐ Support

If you find moshpit useful, please consider giving the repo a **star** ⭐️ and join [stargazers](https://github.com/taylorbanks/moshpit/stargazers).

<!-- Uncomment and update with your own support link if desired:
☕ You can also support me by [buying me a coffee](https://www.buymeacoffee.com/yourusername) ❤️
<br/>
<a href="https://buymeacoffee.com/yourusername" target="_blank"><img src="https://cdn.buymeacoffee.com/buttons/v2/default-yellow.png" width="200"></a>
-->


---

## 🙏 Acknowledgments

- Built with [tview](https://github.com/rivo/tview) and [tcell](https://github.com/gdamore/tcell).
- Inspired by [k9s](https://github.com/derailed/k9s) and [lazydocker](https://github.com/jesseduffield/lazydocker).
- Originally forked from [LazySSH by Adembc](https://github.com/Adembc/lazyssh) - enhanced with mosh protocol support and rebranded as moshpit.
- Mosh protocol by [mosh.org](https://mosh.org) - the mobile shell for reliable remote connections.

