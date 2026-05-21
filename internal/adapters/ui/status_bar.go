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
	"github.com/rivo/tview"
)

func DefaultStatusText() string {
	k := Hex(ActiveTheme.Text)
	return "[" + k + "]↑↓[-] Navigate  • [" + k + "]Enter[-] SSH  • [" + k + "]f[-] Forward  • [" + k + "]x[-] Stop Forward  • [" + k + "]c[-] Copy SSH  • [" + k + "]a[-] Add  • [" + k + "]e[-] Edit  • [" + k + "]g[-] Ping  • [" + k + "]v[-] Group  • [" + k + "]d[-] Delete  • [" + k + "]p[-] Pin/Unpin  • [" + k + "]/[-] Search  • [" + k + "]q[-] Quit"
}

func NewStatusBar() *tview.TextView {
	status := tview.NewTextView().SetDynamicColors(true)
	status.SetBackgroundColor(ActiveTheme.Mantle)
	status.SetTextAlign(tview.AlignCenter)
	status.SetText(DefaultStatusText())
	return status
}
