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
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/taylorbanks/moshpit/internal/core/ports"
)

// showSSHKeyPicker shows a modal list of available SSH public keys. The user
// picks one, ssh-copy-id runs with `-i` pointing at that key. Esc cancels.
func (t *tui) showSSHKeyPicker(alias string, keys []ports.SSHKeyInfo) {
	list := tview.NewList().ShowSecondaryText(true)
	list.SetBorder(true).
		SetTitle(fmt.Sprintf(" 🔑 Install which key onto %q? ", alias)).
		SetTitleAlign(tview.AlignCenter).
		SetBorderColor(ActiveTheme.Teal).
		SetBackgroundColor(ActiveTheme.Base)
	list.SetMainTextColor(ActiveTheme.Text).
		SetSecondaryTextColor(ActiveTheme.Subtext0).
		SetSelectedTextColor(ActiveTheme.Base).
		SetSelectedBackgroundColor(ActiveTheme.Teal).
		SetBackgroundColor(ActiveTheme.Base)

	for i := range keys {
		key := keys[i] // capture
		fingerprint := key.Fingerprint
		if fingerprint == "" {
			fingerprint = "(no fingerprint — ssh-keygen unavailable or not a key)"
		}
		list.AddItem(key.DisplayName, fingerprint, 0, func() {
			t.returnToMain()
			t.runSSHCopyID(alias, key)
		})
	}

	list.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey { //nolint:exhaustive // only Escape handled
		if event.Key() == tcell.KeyEscape {
			t.returnToMain()
			t.showStatusTemp("ssh-copy-id cancelled")
			return nil
		}
		return event
	})

	// Center the picker on screen with flex spacers.
	pickerH := len(keys)*2 + 2 // 2 rows per item (main + secondary) + border
	if pickerH < 6 {
		pickerH = 6
	}
	row := tview.NewFlex().SetDirection(tview.FlexColumn).
		AddItem(nil, 0, 1, false).
		AddItem(list, 72, 0, true).
		AddItem(nil, 0, 1, false)
	page := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(nil, 0, 1, false).
		AddItem(row, pickerH, 0, true).
		AddItem(nil, 0, 1, false)

	t.app.SetRoot(page, true)
	t.app.SetFocus(list)
}
