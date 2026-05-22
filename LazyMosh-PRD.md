# LazyMosh: Mosh Support Feature Addition

### TL;DR

LazyMosh introduces robust mosh protocol support to the existing LazySSH open-source app, allowing users to seamlessly toggle between SSH and mosh on a per-host basis via an intuitive UI element. This update makes LazyMosh a lightweight, cost-effective alternative to expensive and bloated remote terminal apps like Termius, catering to users who require stable roaming SSH connectivity without disturbing or disrupting their existing SSH workflows. All current LazySSH capabilities—including safe, non-destructive SSH config editing—are fully preserved.

---

## Goals

### Business Goals

* Position LazyMosh as a superior, open-source alternative to commercial SSH clients like Termius.

* Attract users who need roaming connection support and reliability across unstable networks.

* Increase active user base and visibility through improved user experience and protocol flexibility.

* Drive open-source contributions and GitHub engagement (stars, forks, issues, pull requests).

### User Goals

* Enable per-host mosh protocol support without disrupting SSH configuration files or key management.

* Offer a persistent UI toggle that remembers each user’s mosh/SSH protocol preference per host.

* Provide automatic fallback to SSH when mosh is unavailable, ensuring uninterrupted access.

* Maintain seamless SSH config compatibility, safe editing, and a non-intrusive, familiar workflow.

### Non-Goals

* Not changing the existing SSH config modification behavior—all safety mechanisms (non-destructive writes, atomic operations, and automated backups) remain identical to LazySSH.

* Not implementing advanced session management, full-featured file transfer, or GUI-based terminal emulation.

* Not supporting additional connection protocols beyond SSH and mosh.

---

## User Stories

**DevOps Engineer**

* As a DevOps Engineer, I want to enable mosh for critical infrastructure servers, so my connections persist when switching Wi-Fi or dealing with network interruptions.

* As a DevOps Engineer, I want to be notified if mosh isn't available and have the tool automatically revert to SSH, so my workflow is never blocked.

* As a DevOps Engineer, I want to enable mosh for all servers with a specific tag (e.g., "production"), so I can configure multiple hosts at once instead of toggling individually.

**Remote Developer**

* As a Remote Developer, I want to toggle mosh on for my remote development machines, so I can continue working even as I change locations or lose connectivity.

* As a Remote Developer, I want my mosh/SSH preference saved for each host, so I don’t have to set it every time.

**System Administrator**

* As a System Administrator, I want to see at a glance which hosts have mosh available or enabled, so I can quickly troubleshoot or adjust configurations.

* As a System Administrator, I rely on LazyMosh to safely and non-destructively edit my SSH config, preserving comments, formatting, order, and unmanaged fields, just as LazySSH did.

---

## Functional Requirements

* **Per-Host Connection Protocol Toggle (Priority: High)**

  * **UI Toggle:** A three-state toggle in the TUI for each host: SSH (default), Mosh Enabled, and Mosh Unavailable.

  * **Persistent Setting:** Per-host protocol preference persists across sessions, stored outside the SSH config in a dedicated LazyMosh settings file.

  * **No SSH Config Corruption:** All toggles/settings are managed separately; SSH config edits (e.g., for hosts, tags, keys) remain available with all original safety mechanisms.

  * **Bulk Operations (Priority: Medium):** Support enabling/disabling mosh for all hosts matching a tag (e.g., "enable mosh for all hosts tagged production").

* **Mosh Availability Detection (Priority: High)**

  * **Local Binary Check:** Detect mosh binary on the user's local system, searching OS-specific standard paths and using `which`/`command -v`.

  * **Remote Support Check (future):** Optionally, check mosh support on the remote host and reflect it in the UI in later phases.

* **Automatic Fallback & Error Handling (Priority: High)**

  * **Warning & Fallback:** When mosh is selected but unavailable, present a clear warning, set the toggle to "Mosh Unavailable," and seamlessly fallback to SSH. No session is ever blocked.

* **Settings Storage (Priority: High)**

  * **LazyMosh Directory:** Store protocol preferences and other LazyMosh app settings in `~/.lazymosh/settings.json` (or similar). Only mosh-related and app-specific settings go here—SSH configs remain in `~/.ssh/config`.

  * **Preserved Modernization:** All core capabilities of LazySSH to add/edit/delete hosts, pin/tag, and edit your SSH config with non-destructive, atomic operations and automatic backups, remain present and unchanged.

* **App Renaming (Priority: Medium)**

  * Update branding, documentation, and greet screen to reflect the transition from LazySSH to LazyMosh.

---

## User Experience

**Entry Point & First-Time User Experience**

* User launches LazyMosh via CLI as before; greet screen or status bar introduces the new per-host mosh toggle.

* If upgrading from LazySSH, all workflows remain familiar and all prior features (config editing, backup, tag/pin, fast navigation) are identical.

**Core Experience**

1. **Host List:** User is shown the list of hosts, parsed as before from their existing `~/.ssh/config`. All legacy config modification and management features remain available.

2. **Protocol Toggle:** Each host entry now includes a visual toggle for SSH/mosh selection, defaulting to SSH.

  * If mosh binary not detected, the toggle sets to "Mosh Unavailable," displaying a warning (e.g., “Mosh not installed. Using SSH instead.”).

  * Toggle states use clear color-coding and/or icons for accessibility.

3. **Connection:** User selects a host to connect:

  * If mosh is enabled and available, LazyMosh constructs and invokes the appropriate `mosh user@host` command, reusing parsed SSH configuration as needed.

  * If mosh was selected but is unavailable, the user receives a non-blocking warning and the app falls back to a standard SSH session.

  * On disconnect, a session log indicates the protocol used.

  * Per-host preferences persist and reload on future launches.

4. **Bulk Toggle:** Users may toggle mosh for all hosts matching a tag (e.g., 'production') via a bulk option.

**Advanced Features & Edge Cases**

* Users can manually edit `~/.lazymosh/settings.json` for scripting/automation.

* If settings file is deleted or corrupted, toggles reset to SSH default without affecting host access.

* Mosh binary detection is robust and cross-platform.

* If mosh unexpectedly fails, user receives actionable error messages and immediate fallback.

**UI/UX Highlights**

* Three-state high-contrast toggle; text alternatives for colorblind users.

* Responsive TUI supporting terminal resizing and large configs.

* Edge case warnings never block sessions.

---

## Narrative

Remote developer Alex works from various locations, relying on SSH for server access. Prior tools dropped sessions during Wi-Fi changes, forcing Alex to switch to heavy, paid apps with little customizability. Using LazyMosh, Alex is instantly able to toggle mosh for any server—without changing her SSH config or worrying about costly mistakes. If mosh isn’t installed, LazyMosh lets her know and simply uses SSH. Her hand-tuned config, with all custom comments and structuring, stays totally intact through every update. Over time, Alex enjoys rock-solid remote sessions and the confidence that every change is safe, backed up, and never destructive—true open-source empowerment.

---

## Success Metrics

### User-Centric Metrics

* % of active users leveraging mosh toggle.

* Rate of successful mosh connections vs. SSH fallback.

* User satisfaction (GitHub stars, issues, feedback).

* Growth in persistent per-host preference usage.

### Business Metrics

* User base growth post-release.

* Increase in GitHub stars/forks in 60 days.

* Drop in connection interruption reports.

* Momentum in protocol-support PRs.

### Technical Metrics

* Fallback latency when mosh is unavailable (<200ms).

* Cross-platform mosh-detection accuracy (>98%).

* Low session invocation error rate (<1% protocol misdetection).

* Fast per-host preference persistence/recall (<50ms).

### Tracking Plan

* Toggle state changes (host and bulk).

* Connection outcomes (mosh/SSH/fallback).

* Mosh unavailability warnings.

* First-use and bulk-toggle event tracking.

* Settings file access and error monitoring.

---

## Technical Considerations

### Technical Needs

* Cross-platform mosh binary detection via standard paths + PATH lookup.

* Per-host protocol preferences via JSON/YAML at `~/.lazymosh/settings.json`.

* TUI extension for three-state protocol toggle per host.

* Connection abstraction for SSH/mosh invocation, using config parse.

* Separation: mosh settings and app settings in `~/.lazymosh/`; SSH configs continue at `~/.ssh/config`.

### Integration Points

* Continues to safely read and write SSH config files using the existing non-destructive parser that preserves comments, spacing, and unmanaged fields. All existing backup mechanisms (`config.original.backup` and rolling `~/.ssh/config--lazymosh.backup` files) remain unchanged.

* Relies on standard OS path lookups for mosh; may use shell utilities for best coverage.

* TUI enhancements built atop existing structure—no migration to GUI.

### Data Storage & Privacy

* All LazyMosh settings (mosh toggle state, pinned servers, tags, app preferences) stored *only* in `~/.lazymosh/`. This is the sole location for new per-host protocol data.

* SSH config modification continues exactly as in LazySSH: safe, atomic edits, with non-destructive diffing, automated backups, comment/format preservation, and never storing credentials.

* No telemetry or credential copying; all connections flow through SSH/mosh as invoked.

* No external communications except explicit user-initiated sessions.

### Scalability & Performance

* Capable of handling hundreds of hosts without UI lag.

* Efficient settings I/O; isolated mosh data guarantees no config performance penalty.

### Potential Challenges

* Accurate mosh detection across varying platforms/install locations.

* Handling incomplete or erroneous mosh installs robustly.

* Ensuring settings file failures never impact SSH config access.

* Clean, clear UX for fallback errors and warnings.

---

## Milestones & Sequencing

### Project Estimate

* **Extra-Small:** 3 working days maximum, targeting a working MVP for testing.

### Team Size & Composition

* **Extra-Small Team:** 1 developer using Claude Code to modify the existing codebase.

### Suggested Phases

**Phase 1: MVP Core (Day 1–2)**

* Mosh binary detection

* Per-host three-state toggle UI

* Preference storage in `~/.lazymosh/`

* Mosh connection/fallback logic

**Phase 2: Polish & Bulk Operations (Day 3)**

* Bulk/tag-based mosh toggling

* Refined warnings

* Settings migration if upgrading from LazySSH

* Branding update to LazyMosh

**Phase 3 (Future/Post-MVP)**

* Documentation refresh

* README and release notes

---

**Key Architectural Principle:**  

All SSH config modification capabilities of LazySSH are preserved. LazyMosh only adds protocol selection with per-host mosh preference tracked in `~/.lazymosh/`. SSH config edits—including host add, edit, tag, or removal—remain safely handled with non-destructive parsing, atomic writes, and robust rolling backups.