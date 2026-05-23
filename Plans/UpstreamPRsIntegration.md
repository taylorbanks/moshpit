# Upstream PR Integration Plan — `Adembc/lazyssh` open PRs → `moshpit`

**Status:** Plan only. No code changes yet. Authored 2026-05-23.

---

## 1. Goal

Bring valuable open pull requests from `Adembc/lazyssh` into `moshpit` **as
contributions from their original authors**, not as anonymous code drops.
Every applicable PR's commits should be cherry-picked from the **author's
own fork** so that `git log` and the GitHub contributors graph credit the
original author. Where the diff has drifted too far for a clean
cherry-pick, we manually port the code and attribute via a
`Co-authored-by:` trailer.

## 2. Non-goals

- **Not** merging anything in this pass — this document is the plan only.
- Not replacing moshpit features already present (themes, grouped view,
  mosh protocol, splash, `i` toggle) with upstream equivalents that
  duplicate or regress them.
- Not landing massive feature PRs (i18n, multi-thousand-line additions)
  in the same wave as small fixes — they get their own evaluation pass.

## 3. Hard constraints

- **Preserve authorship.** Use `git cherry-pick -x <SHA>` whenever the
  cherry-pick is clean; rely on the original committer metadata. When a
  manual port is required, the moshpit commit MUST end with
  `Co-authored-by: Original Author <email>` so GitHub still attributes.
- **Never break** the moshpit features already shipped: mosh toggle (m/M),
  theme picker (T), grouped view (v), splash + `i` toggle, last-SSH
  column (l), 60+ SSH config fields, automatic ssh_config backups,
  protocol prefs in `~/.moshpit/metadata.json`.
- **Build + vet + gofmt + test must be green** after every single PR
  integrated. Any PR that breaks them gets reverted on its branch, not
  papered over.
- **One PR per integration branch**, one squash/merge commit on `main`
  per PR. No mega-merges.

## 4. moshpit's diverged surface (what conflicts will hit)

Every UI file under `internal/adapters/ui/` has moshpit-specific edits.
The data adapter (`internal/adapters/data/ssh_config_file/`) and core
ports/services were touched for the mosh protocol detector, metadata
manager, and `AppConfig`. `cmd/main.go` was heavily restructured for the
mosh-metadata migration path, theme/group/splash plumbing, and the cobra
`Version` flag. New moshpit-only files: `theme.go`, `theme_picker.go`,
`group.go`, `splash.go`, `splashpit.go`, `protocol_detector.go`,
`config_manager.go`.

Any incoming PR that touches one of those areas is a likely conflict.

## 5. PR survey

PRs are open against `Adembc/lazyssh` as of 2026-05-23. Source format:
`#NN  fork:branch  Δ  files  date  title`.

| #   | Author / fork                        | Δ              | Files | Subject                                                          | Tier |
| --- | ------------------------------------ | -------------- | ----- | ---------------------------------------------------------------- | ---- |
| 115 | manato-tajiri / main                 | +25 / -3       | 2     | fix: 't' tag editing on fresh entries, allow last-tag removal    | 1    |
| 114 | maxadc / feat/i18n-chinese-support   | +2751 / -406   | 24    | feat(ui): Chinese localization + i18n framework                  | Skip |
| 113 | gonsalvesc / XDG-base-directory      | +20 / -4       | 2     | improve(config): use XDG base dirs for config + logs             | 2    |
| 112 | DelphicOkami / config.dSupport       | +1565 / -252   | 14    | feat(config): top-level file includes                            | 3    |
| 110 | komapro / docs/fedora-copr           | +11 / -2       | 1     | docs: Fedora Copr install                                        | Skip |
| 109 | ntheanh201 / feat/cli-server-alias   | +60 / -0       | 5     | feat(cli): direct SSH via alias arg (`lazyssh myhost`)           | 2    |
| 107 | OlalalalaO / fix/after-usemouse      | +10 / -2       | 1     | fix: j/k stop working after mouse click                          | 1    |
| 106 | breakersun / fix/backspace-in-input  | +30 / -0       | 1     | fix: backspace in input fields                                   | 1    |
| 102 | k161196 / feat/number-navigation     | +25 / -4       | 5     | feat(ui): numeric focus shortcuts (search/servers/details)       | 2    |
| 101 | k161196 / feat/active-sessions       | +637 / -74     | 6     | feat: active sessions panel + per-session controls               | 3    |
| 100 | davidszp / feat/theme-support        | +1276 / -77    | 24    | feat(ui): dark/light/system theme support with auto-follow       | Skip\* |
| 99  | Q0 / group                           | +501 / -33     | 11    | feat: server groups with nested groups + tmux                    | 3    |
| 98  | OleksandrKucherenko / claude/issue-91| +2856 / -44    | 14    | feat: Git SSH key configuration                                  | 3    |
| 97  | leoncamel / align-host-alias         | +20 / -3       | 2     | fix: align server alias                                          | Skip\* |
| 88  | gaoyifan / pr/persistent-sort-settings| +157 / -1     | 4     | feat: persist sort mode                                          | 2    |
| 87  | levinion / main                      | +17 / -2       | 4     | feat: shortcut to copy host                                      | 1    |
| 86  | arniom / fix/85-user-validation      | +8 / -2        | 2     | fix(ui): allow `@` and `:` in username                           | 1    |
| 84  | omani / feature_83                   | +81 / -28      | 4     | feat: `--sshconfig` flag for custom ssh config path              | 2    |
| 76  | aabichou / fix/Include-globs         | +236 / -4      | 6     | fix: preserve Include directives + global ssh config             | 3    |
| 60  | malaiwah / feat/parser-include       | +739 / -51     | 16    | feat: recursively follow `Include`'d files                       | 3    |
| 59  | malaiwah / feat/ui-fuzzy-search      | +197 / -12     | 2     | feat(ui): fuzzy subsequence scoring for `/` search               | 2    |
| 56  | malaiwah / feat/ui-install-ssh-key   | +41 / -3       | 6     | feat: install SSH key via `ssh-copy-id` (`K` shortcut)           | 2    |
| 52  | tan9 / feat/ping-all                 | +422 / -29     | 10    | feat: Ping All with right-aligned status indicators              | 2    |
| 33  | tan9 / feat/add-from-ssh             | +2606 / -116   | 16    | feat: Paste SSH command parser → add as new server               | 3    |

\* `#100` (themes) and `#97` (alias align) are skipped because moshpit
already covers the same territory differently (digunix's 12-theme
library + aligned column layout). One narrow piece of `#100` — OS-system
theme auto-follow — is worth re-evaluating later as a standalone
enhancement (see §11.4).

## 6. Per-PR analysis

### Tier 1 — small, isolated, low conflict-risk

These should land first to validate the integration mechanics and warm
up the CI pipeline.

- **#86** (`arniom`) — username validation regex fix. `validation.go` +
  test only. moshpit's `validation.go` has lint-only diffs since fork →
  clean cherry-pick expected.
- **#106** (`breakersun`) — backspace handling in input fields. Single
  file (`server_form.go`). moshpit's `server_form.go` has only gofmt
  diffs → clean expected.
- **#107** (`OlalalalaO`) — mouse-click stealing j/k focus. Single
  file (`handlers.go`). The change is localized to mouse handling, not
  the global key switch where moshpit added entries. Expect 1 hunk to
  hand-resolve at worst.
- **#115** (`manato-tajiri`) — `t` tag-editing bugfix. `handlers.go` +
  `server_list.go`. Localized to tag-editing flow which moshpit hasn't
  touched. Likely clean.
- **#87** (`levinion`) — copy-host shortcut. Adds a key (probably `H`
  or `y`). Touches `handlers.go`, `server_details.go`, `status_bar.go`,
  README. moshpit changed all three — small text-region conflicts
  possible. Need to check the chosen key doesn't collide with moshpit's
  `m`/`M`/`T`/`v`/`l`/`i` (likely a free letter).

### Tier 2 — needs minor adaptation, individually valuable

- **#113** (`gonsalvesc`) — XDG Base Directory. Must integrate with
  moshpit's `.lazyssh → .lazymosh → .moshpit` migration in
  `cmd/main.go:getMetadataPath`. New path: respect
  `$XDG_CONFIG_HOME/moshpit` and `$XDG_STATE_HOME/moshpit` when set, with
  the existing `~/.moshpit` fallback. Migration logic should also check
  legacy `~/.lazyssh`/`~/.lazymosh` regardless of XDG.
- **#109** (`ntheanh201`) — `moshpit <alias>` direct connect. Adapt to
  moshpit's cobra rootCmd. Add as a positional arg the rootCmd accepts
  before launching the TUI. Touches `cmd/main.go`, `server_service`,
  ports. Verify it doesn't interfere with the splash on TTY-attached
  vs non-interactive launches.
- **#84** (`omani`) — `--sshconfig` flag. Same cobra surface as #109;
  do after #109 so they share one main.go rework.
- **#88** (`gaoyifan`) — persist sort mode. PR introduces a new
  `settings.go`. moshpit already has `AppConfig` in
  `internal/adapters/data/ssh_config_file/config_manager.go` (used for
  `Theme`, `GroupedView`, `ShowSplash`). Adapt: add `SortMode string` to
  `AppConfig`, wire through `NewTUI` constructor and a save callback —
  same pattern as theme / grouped / splash. The PR's separate settings
  file gets dropped in favor of the unified config.
- **#102** (`k161196`) — numeric focus shortcuts (1/2/3 jump to
  search / list / details). Adapt the new key cases into moshpit's
  global key switch in `handlers.go`. Make sure they don't shadow
  anything moshpit added.
- **#56** (`malaiwah`) — `K` shortcut runs `ssh-copy-id`. This is on
  moshpit's "Upcoming" list. Conflict surface: introduces a `hint_bar.go`
  in lazyssh that moshpit doesn't have (moshpit uses `status_bar.go`).
  Adapt by porting the `K` action and binding into moshpit's existing
  status-bar / handlers, and adding a `K` row to the details-pane
  Commands list.
- **#59** (`malaiwah`) — fuzzy search scoring. Two files. Conflicts on
  `handlers.go` (heavily modified) and `server_service.go` (mosh-related
  changes). Should be a careful hand-port of the scoring function, not
  a raw cherry-pick.
- **#52** (`tan9`) — Ping All. New feature, three commits. Touches
  `server_form`, `status_bar`, `tui`, `domain.Server`. Adapt around
  moshpit's domain/server changes. Verify ping output coexists with
  moshpit's protocol indicator column.

### Tier 3 — complex, overlapping, or feature-heavy

These need their own focused integration windows after Tier 1+2 land.

- **Include directive cluster — #76, #60, #112.** All three approach
  `Include` handling differently. #76 is the smallest (preserve Include
  directives on write). #60 recursively follows them (adds parser
  rework). #112 adds top-level includes. Action: do #76 first
  (minimal, low risk). Evaluate #60 and #112 only after #76 is in.
  Likely pick one of #60 or #112, not both.
- **#101** (`k161196`) — active sessions panel. Substantial new UI
  surface (new pane, kill/prefill actions, ssh-arg parser refactor).
  Touches `domain.Server` and `ports.ServerService` — both
  moshpit-modified. Plan its own multi-day integration window.
- **#99** (`Q0`) — nested server groups + tmux integration. Overlaps
  conceptually with moshpit's tag-based grouped view but is a different
  model (explicit `Group` field, nested groups, tmux session per
  group). Decide whether to add alongside tags or refactor moshpit's
  grouping to use both. Decision required before integration begins.
- **#98** (`OleksandrKucherenko`) — Git SSH key configuration. Adds a
  whole subsystem (sshkeys_list, sshkey_details, git_info, git_ssh_setup,
  edit_key_comment). Doesn't conflict with moshpit's mosh logic but is
  a large feature — defer.
- **#33** (`tan9`) — paste-SSH parser. New ssh_parser, alias_utils,
  validation. Touches server_form and validation heavily. Useful but
  large. Defer.

### Skipped (why, in one line each)

- **#114** (i18n Chinese) — 2700+ line architectural change introducing
  i18n framework, AuthMethod/LoginMode, and config import/export.
  Re-scope as its own project rather than a PR integration.
- **#100** (davidszp themes) — moshpit already has a 12-theme library
  + in-app picker. Extracting just the OS-theme-follow logic might be
  worth its own enhancement later — track in §11.4.
- **#110** (Fedora Copr install) — the Copr repo is for `Adembc/lazyssh`
  binaries; a moshpit Copr would need its own setup. Don't carry the
  doc snippet until that exists.
- **#97** (align server alias) — superseded by digunix's aligned
  column layout already in moshpit. Confirm with a side-by-side before
  finally dismissing.

## 7. Recommended order

The order optimizes for: (a) build confidence on small wins first,
(b) land config/CLI changes before feature changes that depend on them,
(c) finish each tier before starting the next.

**Wave 1 — Tier 1 (small bugfixes & shortcuts)**
1. `#86` username validation
2. `#106` backspace in inputs
3. `#107` j/k after mouse click
4. `#115` tag editing fix
5. `#87` copy-host shortcut

**Wave 2 — Config & CLI surface**
6. `#113` XDG base dirs
7. `#109` direct connect via alias
8. `#84` `--sshconfig` flag
9. `#88` persist sort mode

**Wave 3 — UI fixes & light features**
10. `#102` numeric focus shortcuts
11. `#59` fuzzy search scoring
12. `#56` `ssh-copy-id` install (`K`)
13. `#52` Ping All

**Wave 4 — Tier 3 evaluation pass (decisions, not yet integrations)**
14. Decide on the Include cluster: do `#76` first.
15. Decide on `#101` active sessions — own branch.
16. Decide on `#99` groups vs moshpit's tags — design call.
17. Decide on `#98` Git SSH key — own branch.
18. Decide on `#33` paste-SSH parser — own branch.

## 8. Authorship-preserving mechanics

### 8.1 Add the author's fork as a remote (per PR)

```bash
git remote add <login> https://github.com/<login>/lazyssh.git
git fetch <login> <branch>
```

Use the GitHub login as the remote name. Naming consistently keeps the
mental model clean across 20+ remotes.

### 8.2 Find the commits to pick

`gh pr view <num> --repo Adembc/lazyssh --json commits` lists every
commit on the PR's head branch. Cherry-pick only the commits authored
by the PR's author — skip any `Merge branch 'main'` commits and any
prior upstream commits that show up in the PR's history due to merge
direction.

### 8.3 Clean cherry-pick

```bash
git checkout -b pr-<num>-<short-description>
git cherry-pick -x <sha1> [<sha2> ...]
```

`-x` adds a `(cherry picked from commit <sha>)` line to the body —
keeps a paper trail back to lazyssh. Author + email are preserved on
the cherry-picked commit; that is what makes the PR author a moshpit
contributor.

### 8.4 Conflict resolution that preserves authorship

When cherry-pick hits a conflict:

1. Resolve only as much as needed to satisfy the PR's intent.
2. **Do not** restructure the surrounding code — moshpit's deviations
   stay as they are.
3. `git cherry-pick --continue` so the original commit's author and
   message survive. The merge resolution shows up as part of that
   commit, still authored by the upstream contributor.

### 8.5 When cherry-pick is not viable

If a file moved, was renamed, was effectively rewritten in moshpit, or
the PR's structure is too tangled to cherry-pick:

1. Hand-port the change.
2. Commit it with your own authorship but include the contributor in
   the trailer:

```
   feat(ui): port <PR description> from lazyssh#<num>

   Co-authored-by: <Author Name> <author@email>
   Originally proposed at: https://github.com/Adembc/lazyssh/pull/<num>
```

GitHub picks up `Co-authored-by:` and shows both contributors on the
commit. The link in the body preserves provenance.

### 8.6 Cleanup

After landing on moshpit `main`, `git remote remove <login>` so the
local repo doesn't drift into a haystack of stale remotes.

## 9. Pre-integration baseline

Before Wave 1 starts:

1. From clean `main`: capture a "known-good" snapshot.
   ```bash
   git checkout main && git pull
   go build ./... && go vet ./... && go test ./...
   gofmt -l internal/ cmd/
   ```
   All must pass / be empty. If not, fix on `main` before integrating
   anything else.

2. Functional smoke test — open the TUI and exercise:
   - Splash renders + dismisses on keypress.
   - `i` toggles the splash preference (status bar confirms, details
     pane chip updates).
   - `T` opens theme picker; `v` toggles grouped view; `m` toggles
     SSH/Mosh; `M` bulk-toggles by tag.
   - At least one SSH connection completes.

3. Record the smoke-test result in a one-line note in this plan's
   `## 13 Log` section before each wave starts.

## 10. Per-PR playbook (template)

For every PR `#NN`:

1. **Branch.** From up-to-date `main`:
   `git checkout -b pr-NN-<slug>`.
2. **Remote.** `git remote add <author> https://github.com/<author>/lazyssh.git && git fetch <author>`.
3. **Cherry-pick or port** per §8.
4. **Build gates.** Must all pass before commit / merge:
   - `go build ./...`
   - `go vet ./...`
   - `gofmt -l internal/ cmd/` → blank
   - `go test ./...`
5. **Regression smoke.** Re-run the §9 smoke test. Anything that worked
   before this PR must still work.
6. **PR-specific functional test.** Exercise the feature the PR adds.
   Record what you did.
7. **Open a PR on moshpit** (against `main`). Title: copy from upstream
   with `(via lazyssh#NN)` appended. Body: link upstream PR, list
   contributor by GitHub handle, summarize the change.
8. **Auto-review** via `/autoplan` or `/plan-eng-review` if useful.
9. **Merge** as a single commit (preserve cherry-picked commits if a
   merge keeps authorship cleanly; otherwise squash-merge with
   `Co-authored-by:` trailer).
10. **Update `Plans/UpstreamPRsIntegration.md`** §13 log with the
    resulting commit SHA.

## 11. PR-specific integration notes

### 11.1 #113 XDG base dirs — adaptation steps

Current `cmd/main.go:getMetadataPath` has the
`~/.lazyssh → ~/.lazymosh → ~/.moshpit` migration baked in. The XDG
patch only knows about `~/.lazyssh`. Adaptation:

- Compute `xdgConfigDir = $XDG_CONFIG_HOME ?? ~/.config`,
  `xdgStateDir = $XDG_STATE_HOME ?? ~/.local/state`.
- New canonical paths: `$xdgConfigDir/moshpit/{config.json,metadata.json}`
  and `$xdgStateDir/moshpit/log` (logger).
- Add `$xdgConfigDir/moshpit` to the migration probe order, after
  `~/.moshpit`. Don't break existing users whose data lives in
  `~/.moshpit/`.
- Document the new search order in the README's Security Notice
  section.

### 11.2 #109 + #84 CLI surface

Both edit the cobra rootCmd path. Land them together on the same
integration branch:

- `#109`: rootCmd takes optional positional `[alias]`; when present,
  bypass the TUI and exec `ssh <alias>` (or `mosh <alias>` if the
  per-host metadata says so). Make sure this code path **does not**
  render the splash — only the TUI path does.
- `#84`: add `--sshconfig <path>` flag; thread through repository
  constructor. Default still resolves to `~/.ssh/config`.

### 11.3 #88 persist sort mode — fold into AppConfig

Don't bring upstream's separate `settings.go` file. Instead:

- Add `SortMode string` to `AppConfig` (omitempty).
- In `cmd/main.go`, load `appConfig.SortMode`, pass to `NewTUI`, add a
  `onSortSave` callback — same shape as `onThemeSave` etc.
- In `tui.go`, surface a `SortMode` setter the existing `handleSortToggle`
  / `handleSortReverse` already call; have them also call the save
  callback.

### 11.4 #100 (skipped) future enhancement — OS theme auto-follow

If a future "follow system dark/light" toggle is wanted, the relevant
pieces from `#100` are `theme_detect_darwin.go`, `theme_detect_linux.go`,
`theme_detect_windows.go`, `theme_watch_*.go`. Those can be ported in
isolation as a new toggle in the theme picker without bringing in
davidszp's full theme architecture. Tracked here, not scheduled.

## 12. Verification gates summary

After **every** PR integration, before merging the integration branch:

- Build clean: `go build ./...`
- Vet clean: `go vet ./...`
- Format clean: `gofmt -l internal/ cmd/` is empty
- Tests pass: `go test ./...`
- Smoke (manual): launch the TUI; splash + dismiss; theme picker;
  group toggle; mosh toggle; one SSH connection works.
- The PR's own feature works as advertised.

Anything red → fix on the branch or abort the integration.

## 13. Log

| Wave | PR  | Branch              | Result | Merged commit | Notes                                                                                                |
| ---- | --- | ------------------- | ------ | ------------- | ---------------------------------------------------------------------------------------------------- |
| 1    | 107 | pr-107-mouse-focus  | merged | `4c2d030`     | Cherry-picked `54d33de` from `OlalalalaO/lazyssh:fix/after-usemouse`; auto-merge resolved; FF-merge. |

### Notes on PRs evaluated but skipped during Wave 1

- **#86** — niche (`@`/`:` in usernames is only used for Kerberos / Windows-AD / SSO bastions). Holding unless a concrete need surfaces.
- **#106** — the backspace bug doesn't exist in moshpit: `setupKeyboardShortcuts`'s input capture only acts on Ctrl combos / KeyEscape / KeyCtrlS and falls through for every other key, so tview's InputField receives backspace normally. Defensive code without a bug to fix.

## 14. Rollback

Each integration is a single squash-merge on `main`. If a regression is
found post-merge, `git revert <merge-sha>` is enough. The contributor's
upstream PR stays the authoritative source — moshpit's cherry-pick can
be re-attempted later if rolled back.

If a series of merges later turn out to interact badly, revert in
reverse order and reopen the affected wave. Don't try to surgically
patch across multiple landed PRs.

---

*This plan is intentionally written so any future maintainer (or the
DA in a new session) can pick up at any point. Update §13 as work
proceeds.*
