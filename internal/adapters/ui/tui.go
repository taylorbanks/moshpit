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
	"go.uber.org/zap"

	"github.com/taylorbanks/moshpit/internal/core/domain"
	"github.com/taylorbanks/moshpit/internal/core/ports"
	"github.com/rivo/tview"
)

type App interface {
	Run() error
}

type tui struct {
	logger *zap.SugaredLogger

	version string
	commit  string

	app           *tview.Application
	serverService ports.ServerService

	header     *AppHeader
	searchBar  *SearchBar
	serverList *ServerList
	details    *ServerDetails
	statusBar  *tview.TextView

	root    *tview.Flex
	left    *tview.Flex
	content *tview.Flex

	sortMode           SortMode
	groupedView        bool
	showLastSSH        bool
	onThemeSave        func(string)
	onGroupedViewSave  func(bool)
}

func NewTUI(logger *zap.SugaredLogger, ss ports.ServerService, version, commit string, onThemeSave func(string), groupedView bool, onGroupedViewSave func(bool)) App {
	return &tui{
		logger:            logger,
		app:               tview.NewApplication(),
		serverService:     ss,
		version:           version,
		commit:            commit,
		showLastSSH:       true,
		onThemeSave:       onThemeSave,
		groupedView:       groupedView,
		onGroupedViewSave: onGroupedViewSave,
	}
}

func (t *tui) Run() error {
	defer func() {
		if r := recover(); r != nil {
			t.logger.Errorw("panic recovered", "error", r)
		}
	}()
	t.app.EnableMouse(true)
	t.initializeTheme().buildComponents().buildLayout().bindEvents().loadInitialData()
	t.app.SetRoot(t.root, true)
	t.logger.Infow("starting TUI application", "version", t.version, "commit", t.commit)
	if err := t.app.Run(); err != nil {
		t.logger.Errorw("application run error", "error", err)
		return err
	}
	return nil
}

func (t *tui) initializeTheme() *tui {
	th := ActiveTheme
	tview.Styles.PrimitiveBackgroundColor = th.Base
	tview.Styles.ContrastBackgroundColor = th.Surface0
	tview.Styles.BorderColor = th.Surface1
	tview.Styles.TitleColor = th.Subtext1
	tview.Styles.PrimaryTextColor = th.Text
	tview.Styles.TertiaryTextColor = th.Subtext0
	tview.Styles.SecondaryTextColor = th.Subtext0
	tview.Styles.GraphicsColor = th.Surface1
	return t
}

func (t *tui) buildComponents() *tui {
	t.header = NewAppHeader(t.version, t.commit, RepoURL)
	t.searchBar = NewSearchBar().
		OnSearch(t.handleSearchInput).
		OnEscape(t.blurSearchBar).
		OnNavigate(t.handleSearchNavigate)
	IsForwarding = t.serverService.IsForwarding
	IsMoshAvailable = t.serverService.IsMoshAvailable

	t.serverList = NewServerList().
		OnSelectionChange(t.handleServerSelectionChange).
		OnReturnToSearch(t.handleReturnToSearch)
	t.details = NewServerDetails()
	t.statusBar = NewStatusBar()

	// default sort mode
	t.sortMode = SortByAliasAsc

	return t
}

func (t *tui) buildLayout() *tui {
	t.left = tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(t.searchBar, 3, 0, false).
		AddItem(t.serverList, 0, 1, true)

	right := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(t.details, 0, 1, false)

	t.content = tview.NewFlex().SetDirection(tview.FlexColumn).
		AddItem(t.left, 0, 3, true).
		AddItem(right, 0, 2, false)

	t.root = tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(t.header, 2, 0, false).
		AddItem(t.content, 0, 1, true).
		AddItem(t.statusBar, 1, 0, false)
	return t
}

func (t *tui) bindEvents() *tui {
	t.root.SetInputCapture(t.handleGlobalKeys)
	return t
}

func (t *tui) loadInitialData() *tui {
	servers, _ := t.serverService.ListServers("")
	sortServersForUI(servers, t.sortMode)
	t.updateListTitle()
	t.updateServerList(servers)

	return t
}

// updateServerList populates the server list using grouped or flat mode.
func (t *tui) updateServerList(servers []domain.Server) {
	if t.groupedView {
		entries := groupServersByTag(servers, t.sortMode)
		t.serverList.UpdateServersGrouped(entries)
	} else {
		t.serverList.UpdateServers(servers)
	}
}

func (t *tui) updateListTitle() {
	if t.serverList != nil {
		title := " Servers — Sort: " + t.sortMode.String()
		if t.groupedView {
			title += " | Grouped"
		}
		title += " "
		t.serverList.SetTitle(title)
	}
}

// rebuildUI rebuilds all components and layout after a theme change.
func (t *tui) rebuildUI() {
	t.buildComponents().buildLayout().bindEvents().loadInitialData()
	t.app.SetRoot(t.root, true)
}
