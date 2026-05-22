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
	"strings"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

const (
	splashBoxWidth  = 44                    // bordered box width, in cells
	splashBoxHeight = 16                    // bordered box height, in cells
	splashFrame     = 55 * time.Millisecond // animation tick interval
)

// buildSplash returns a centered, themed loading screen: a bordered box with
// the app name and icon set into the top border, a tagline, and a small
// mosh-pit animation that doubles as the loading indicator.
func (t *tui) buildSplash() tview.Primitive {
	t.splashPit = newMoshSim()
	t.splashDone = make(chan struct{})

	t.splashBody = tview.NewTextView().
		SetDynamicColors(true).
		SetTextAlign(tview.AlignCenter)
	t.splashBody.SetBackgroundColor(ActiveTheme.Base)
	t.renderSplash(t.splashPit.render())

	// The box border is unbroken except by the title — the icon and app name
	// sit in the top edge, breaking the line there.
	box := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(t.splashBody, 0, 1, false)
	box.SetBorder(true).
		SetTitle(" 🤘 moshpit ").
		SetTitleAlign(tview.AlignCenter).
		SetBorderColor(ActiveTheme.Teal).
		SetBackgroundColor(ActiveTheme.Base)

	// Center the box on both axes with flexible spacers.
	midRow := tview.NewFlex().SetDirection(tview.FlexColumn).
		AddItem(nil, 0, 1, false).
		AddItem(box, splashBoxWidth, 0, false).
		AddItem(nil, 0, 1, false)
	root := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(nil, 0, 1, false).
		AddItem(midRow, splashBoxHeight, 0, false).
		AddItem(nil, 0, 1, false)

	// Any key dismisses the splash and reveals the main UI.
	root.SetInputCapture(func(_ *tcell.EventKey) *tcell.EventKey {
		t.dismissSplash()
		return nil
	})

	return root
}

// renderSplash draws the splash body: tagline, the mosh-pit animation frame
// (which doubles as the loading indicator), a continue hint, and the version.
func (t *tui) renderSplash(pit []string) {
	content := "\n" +
		"[" + Hex(ActiveTheme.Subtext0) + "]SSH / Mosh server manager[-]\n\n" +
		strings.Join(pit, "\n") + "\n\n" +
		"[" + Hex(ActiveTheme.Overlay0) + "]mosh any key to continue[-]"

	// Show the version only on real release builds; "develop" is the
	// unstamped placeholder used by `go run` / local builds.
	if t.version != "" && t.version != "develop" {
		content += "\n\n[" + Hex(ActiveTheme.Overlay0) + "]" + t.version + "[-]"
	}

	t.splashBody.SetText(content)
}

// animateSplash steps the mosh-pit animation every frame for as long as the
// splash is shown.
func (t *tui) animateSplash() {
	go func() {
		ticker := time.NewTicker(splashFrame)
		defer ticker.Stop()
		for {
			select {
			case <-t.splashDone:
				return
			case <-ticker.C:
				t.splashPit.step()
				frame := t.splashPit.render()
				t.app.QueueUpdateDraw(func() { t.renderSplash(frame) })
			}
		}
	}()
}

// dismissSplash stops the animation and swaps in the main UI. Safe to call
// more than once.
func (t *tui) dismissSplash() {
	t.splashOnce.Do(func() { close(t.splashDone) })
	t.app.SetRoot(t.root, true)
}
