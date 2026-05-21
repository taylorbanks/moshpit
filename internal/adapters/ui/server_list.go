// Copyright 2025.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ui

import (
	"github.com/taylorbanks/moshpit/internal/core/domain"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type ServerList struct {
	*tview.List
	servers           []domain.Server
	entries           []listEntry
	onSelection       func(domain.Server)
	onSelectionChange func(domain.Server)
	onReturnToSearch  func()
}

func NewServerList() *ServerList {
	list := &ServerList{
		List: tview.NewList(),
	}
	list.build()
	return list
}

func (sl *ServerList) build() {
	sl.List.ShowSecondaryText(false)
	sl.List.SetBorder(true).
		SetTitle(" Servers ").
		SetTitleAlign(tview.AlignCenter).
		SetBorderColor(ActiveTheme.Surface1).
		SetTitleColor(ActiveTheme.Subtext1)
	sl.List.
		SetSelectedBackgroundColor(ActiveTheme.Blue).
		SetSelectedTextColor(ActiveTheme.Crust).
		SetHighlightFullLine(true)

	sl.List.SetChangedFunc(func(index int, mainText string, secondaryText string, shortcut rune) {
		if index >= 0 && index < len(sl.entries) {
			entry := sl.entries[index]
			if entry.isHeader {
				return
			}
			if entry.server != nil && sl.onSelectionChange != nil {
				sl.onSelectionChange(*entry.server)
			}
		}
	})

	sl.List.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		//nolint:exhaustive // We only handle specific keys and pass through others
		switch event.Key() {
		case tcell.KeyLeft, tcell.KeyRight, tcell.KeyBackspace, tcell.KeyBackspace2, tcell.KeyESC:
			if sl.onReturnToSearch != nil {
				sl.onReturnToSearch()
			}
			return nil
		}
		return event
	})
}

func (sl *ServerList) UpdateServers(servers []domain.Server) {
	sl.servers = servers
	sl.entries = make([]listEntry, len(servers))
	for i := range servers {
		sl.entries[i] = listEntry{server: &sl.servers[i]}
	}
	sl.List.Clear()

	for i := range servers {
		primary, secondary := formatServerLine(servers[i])
		idx := i
		sl.List.AddItem(primary, secondary, 0, func() {
			if sl.onSelection != nil {
				sl.onSelection(sl.servers[idx])
			}
		})
	}

	if sl.List.GetItemCount() > 0 {
		sl.List.SetCurrentItem(0)
		if sl.onSelectionChange != nil {
			sl.onSelectionChange(sl.servers[0])
		}
	}
}

// UpdateServersGrouped populates the list with grouped entries (headers + servers).
func (sl *ServerList) UpdateServersGrouped(entries []listEntry) {
	sl.entries = entries
	// Rebuild flat servers slice for compatibility
	sl.servers = nil
	for _, e := range entries {
		if e.server != nil {
			sl.servers = append(sl.servers, *e.server)
		}
	}
	sl.List.Clear()

	// Count servers per group for header display
	counts := make(map[int]int) // header index -> count
	for i, e := range entries {
		if e.isHeader {
			count := 0
			for j := i + 1; j < len(entries); j++ {
				if entries[j].isHeader {
					break
				}
				count++
			}
			counts[i] = count
		}
	}

	for i, e := range entries {
		idx := i
		if e.isHeader {
			header := formatGroupHeader(e.tag, counts[i])
			sl.List.AddItem(header, "", 0, nil)
		} else if e.server != nil {
			primary, secondary := formatServerLine(*e.server, true)
			sl.List.AddItem("  "+primary, secondary, 0, func() {
				if sl.onSelection != nil && idx < len(sl.entries) && sl.entries[idx].server != nil {
					sl.onSelection(*sl.entries[idx].server)
				}
			})
		}
	}

	// Select first non-header entry
	if sl.List.GetItemCount() > 0 {
		firstServer := sl.findNextServer(0, 1)
		if firstServer >= 0 {
			sl.List.SetCurrentItem(firstServer)
			if sl.entries[firstServer].server != nil && sl.onSelectionChange != nil {
				sl.onSelectionChange(*sl.entries[firstServer].server)
			}
		}
	}
}

// findNextServer finds the next non-header entry starting from idx in the given direction.
// Returns -1 if none found.
func (sl *ServerList) findNextServer(idx int, direction int) int {
	for i := idx; i >= 0 && i < len(sl.entries); i += direction {
		if !sl.entries[i].isHeader {
			return i
		}
	}
	return -1
}

// SkipToNextServer adjusts the current selection to skip headers in the given direction.
// Returns true if a valid server was found and selected.
func (sl *ServerList) SkipToNextServer(direction int) bool {
	current := sl.List.GetCurrentItem()
	if current < 0 || current >= len(sl.entries) {
		return false
	}
	if !sl.entries[current].isHeader {
		return true
	}
	next := sl.findNextServer(current+direction, direction)
	if next >= 0 {
		sl.List.SetCurrentItem(next)
		return true
	}
	// Try wrapping
	if direction > 0 {
		next = sl.findNextServer(0, 1)
	} else {
		next = sl.findNextServer(len(sl.entries)-1, -1)
	}
	if next >= 0 {
		sl.List.SetCurrentItem(next)
		return true
	}
	return false
}

func (sl *ServerList) GetSelectedServer() (domain.Server, bool) {
	idx := sl.List.GetCurrentItem()
	if idx >= 0 && idx < len(sl.entries) {
		entry := sl.entries[idx]
		if !entry.isHeader && entry.server != nil {
			return *entry.server, true
		}
	}
	return domain.Server{}, false
}

func (sl *ServerList) OnSelection(fn func(server domain.Server)) *ServerList {
	sl.onSelection = fn
	return sl
}

func (sl *ServerList) OnSelectionChange(fn func(server domain.Server)) *ServerList {
	sl.onSelectionChange = fn
	return sl
}

func (sl *ServerList) OnReturnToSearch(fn func()) *ServerList {
	sl.onReturnToSearch = fn
	return sl
}
