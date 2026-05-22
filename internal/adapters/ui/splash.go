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
	"time"

	"github.com/rivo/tview"
)

// splashDuration is how long the loading screen stays visible before the main
// UI takes over. moshpit loads near-instantly, so without a deliberate minimum
// the splash would never be seen. Kept short on purpose — just enough to show
// the branding on each launch.
const splashDuration = 700 * time.Millisecond

// splashHeight is the fixed line count of the splash text block, used to
// vertically center it within the terminal.
const splashHeight = 9

// buildSplash returns a centered, themed loading screen shown at startup.
func (t *tui) buildSplash() tview.Primitive {
	name := "[" + Hex(ActiveTheme.Text) + "::b]m o s h[-]" +
		"[" + Hex(ActiveTheme.Teal) + "::b]p i t[-]"

	content := "🤘\n\n" +
		name + "\n\n" +
		"[" + Hex(ActiveTheme.Subtext0) + "]SSH / Mosh server manager[-]\n\n" +
		"[" + Hex(ActiveTheme.Green) + "]Loading SSH config…[-]\n\n" +
		"[" + Hex(ActiveTheme.Overlay0) + "]" + t.version + "[-]"

	text := tview.NewTextView().
		SetDynamicColors(true).
		SetTextAlign(tview.AlignCenter).
		SetText(content)
	text.SetBackgroundColor(ActiveTheme.Base)

	// Vertical centering: equal flexible spacers above and below a fixed block.
	return tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(nil, 0, 1, false).
		AddItem(text, splashHeight, 0, false).
		AddItem(nil, 0, 1, false)
}

// scheduleSplashDismiss swaps the splash for the main UI after splashDuration.
// The swap is queued onto the tview event loop so it is draw-safe.
func (t *tui) scheduleSplashDismiss() {
	go func() {
		time.Sleep(splashDuration)
		t.app.QueueUpdateDraw(func() {
			t.app.SetRoot(t.root, true)
		})
	}()
}
