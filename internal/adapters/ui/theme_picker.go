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
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func (t *tui) showThemePicker() {
	previousTheme := ActiveTheme

	list := tview.NewList()
	list.SetBorder(true).
		SetTitle(" Theme Picker (Enter: confirm, Esc: cancel) ").
		SetTitleAlign(tview.AlignCenter)
	list.ShowSecondaryText(false)
	list.SetHighlightFullLine(true)

	// Find the index of the current active theme
	currentIdx := 0
	for i, theme := range Themes {
		name := theme.Name
		if theme.Name == ActiveTheme.Name {
			name += " (current)"
			currentIdx = i
		}
		list.AddItem(name, "", 0, nil)
	}
	list.SetCurrentItem(currentIdx)

	// Apply theme colors to the list itself
	applyThemeToList := func(th Theme) {
		list.SetBackgroundColor(th.Base)
		list.SetBorderColor(th.Surface1)
		list.SetTitleColor(th.Subtext1)
		list.SetMainTextColor(th.Text)
		list.SetSelectedBackgroundColor(th.Blue)
		list.SetSelectedTextColor(th.Crust)
	}
	applyThemeToList(ActiveTheme)

	// Live preview on navigate
	list.SetChangedFunc(func(index int, mainText, secondaryText string, shortcut rune) {
		if index >= 0 && index < len(Themes) {
			ActiveTheme = Themes[index]
			t.initializeTheme()
			applyThemeToList(ActiveTheme)
		}
	})

	// Enter to confirm, Esc to cancel
	list.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() { //nolint:exhaustive // only Escape/Enter handled; other keys pass through
		case tcell.KeyEscape:
			// Revert to previous theme
			ActiveTheme = previousTheme
			t.initializeTheme()
			t.rebuildUI()
			return nil
		case tcell.KeyEnter:
			idx := list.GetCurrentItem()
			if idx >= 0 && idx < len(Themes) {
				ActiveTheme = Themes[idx]
				t.initializeTheme()
				t.rebuildUI()
				if t.onThemeSave != nil {
					t.onThemeSave(ActiveTheme.Name)
				}
				t.showStatusTemp("Theme: " + ActiveTheme.Name)
			}
			return nil
		}
		return event
	})

	t.app.SetRoot(list, true)
	t.app.SetFocus(list)
}
