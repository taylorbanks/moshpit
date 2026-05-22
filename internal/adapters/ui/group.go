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
	"sort"

	"github.com/taylorbanks/moshpit/internal/core/domain"
)

// listEntry represents either a group header or a server in the grouped list view.
type listEntry struct {
	isHeader bool
	tag      string         // group name (for headers)
	server   *domain.Server // nil for headers
}

// groupServersByTag organizes servers into tag-based groups.
// Each server appears under every tag it belongs to.
// Untagged servers appear under "Ungrouped" at the bottom.
// Servers within each group are sorted according to sortMode.
func groupServersByTag(servers []domain.Server, sortMode SortMode) []listEntry {
	// Collect unique tags
	tagSet := make(map[string]bool)
	for i := range servers {
		for _, tag := range servers[i].Tags {
			tagSet[tag] = true
		}
	}

	// Sort tags alphabetically
	tags := make([]string, 0, len(tagSet))
	for tag := range tagSet {
		tags = append(tags, tag)
	}
	sort.Strings(tags)

	entries := make([]listEntry, 0, len(servers))

	// Group servers by each tag
	for _, tag := range tags {
		var group []domain.Server
		for i := range servers {
			for _, sTag := range servers[i].Tags {
				if sTag == tag {
					group = append(group, servers[i])
					break
				}
			}
		}
		if len(group) == 0 {
			continue
		}
		sortServersForUI(group, sortMode)
		entries = append(entries, listEntry{isHeader: true, tag: tag})
		for i := range group {
			entries = append(entries, listEntry{server: &group[i]})
		}
	}

	// Ungrouped servers (no tags)
	var ungrouped []domain.Server
	for i := range servers {
		if len(servers[i].Tags) == 0 {
			ungrouped = append(ungrouped, servers[i])
		}
	}
	if len(ungrouped) > 0 {
		sortServersForUI(ungrouped, sortMode)
		entries = append(entries, listEntry{isHeader: true, tag: "Ungrouped"})
		for i := range ungrouped {
			entries = append(entries, listEntry{server: &ungrouped[i]})
		}
	}

	return entries
}

// formatGroupHeader renders a styled group header string for the list.
func formatGroupHeader(tag string, count int) string {
	color := Hex(ActiveTheme.Mauve)
	return fmt.Sprintf("[%s::b]▼ %s (%d)[-::-]", color, tag, count)
}
