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

package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/taylorbanks/moshpit/internal/adapters/data/ssh_config_file"
	"github.com/taylorbanks/moshpit/internal/logger"

	"github.com/taylorbanks/moshpit/internal/adapters/ui"
	"github.com/taylorbanks/moshpit/internal/core/services"
	"github.com/spf13/cobra"
)

var (
	version   = "develop"
	gitCommit = "unknown"
)

// getMetadataPath returns the metadata file path, migrating from .lazyssh/.moshpit-legacy to .moshpit if needed.
func getMetadataPath(home string) string {
	newPath := filepath.Join(home, ".moshpit", "metadata.json")
	lazymoshPath := filepath.Join(home, ".lazymosh", "metadata.json")
	lazysshPath := filepath.Join(home, ".lazyssh", "metadata.json")

	// If new path exists, use it
	if _, err := os.Stat(newPath); err == nil {
		return newPath
	}

	// Try to migrate from .lazymosh first (legacy path)
	if _, err := os.Stat(lazymoshPath); err == nil {
		if err := os.MkdirAll(filepath.Dir(newPath), 0o750); err == nil {
			if data, err := os.ReadFile(lazymoshPath); err == nil {
				if err := os.WriteFile(newPath, data, 0o600); err == nil {
					fmt.Fprintf(os.Stderr, "Migrated metadata from %s to %s\n", lazymoshPath, newPath)
					return newPath
				}
			}
		}
		return lazymoshPath
	}

	// Try to migrate from .lazyssh (original)
	if _, err := os.Stat(lazysshPath); err == nil {
		if err := os.MkdirAll(filepath.Dir(newPath), 0o750); err == nil {
			if data, err := os.ReadFile(lazysshPath); err == nil {
				if err := os.WriteFile(newPath, data, 0o600); err == nil {
					fmt.Fprintf(os.Stderr, "Migrated metadata from %s to %s\n", lazysshPath, newPath)
					return newPath
				}
			}
		}
		return lazysshPath
	}

	// Fresh install, use new path
	return newPath
}

func main() {
	log, err := logger.New("MOSHPIT")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//nolint:errcheck // log.Sync may return an error which is safe to ignore here
	defer log.Sync()

	home, err := os.UserHomeDir()
	if err != nil {
		log.Errorw("failed to get user home directory", "error", err)
		//nolint:gocritic // exitAfterDefer: ensure immediate exit on unrecoverable error
		os.Exit(1)
	}
	sshConfigFile := filepath.Join(home, ".ssh", "config")
	metaDataFile := getMetadataPath(home)

	configFile := filepath.Join(home, ".moshpit", "config.json")
	configMgr := ssh_config_file.NewConfigManager(configFile, log)

	// Load saved config
	appConfig := configMgr.Load()
	if appConfig.Theme != "" {
		ui.SetActiveTheme(appConfig.Theme)
	}

	// Default grouped view to true when unset
	groupedView := true
	if appConfig.GroupedView != nil {
		groupedView = *appConfig.GroupedView
	}

	serverRepo := ssh_config_file.NewRepository(log, sshConfigFile, metaDataFile)
	serverService := services.NewServerService(log, serverRepo)
	tui := ui.NewTUI(log, serverService, version, gitCommit, func(themeName string) {
		appConfig.Theme = themeName
		if err := configMgr.Save(appConfig); err != nil {
			log.Warnw("failed to save theme preference", "error", err)
		}
	}, groupedView, func(grouped bool) {
		appConfig.GroupedView = &grouped
		if err := configMgr.Save(appConfig); err != nil {
			log.Warnw("failed to save grouped view preference", "error", err)
		}
	})

	rootCmd := &cobra.Command{
		Use:   ui.AppName,
		Short: "SSH/Mosh server manager with protocol flexibility",
		RunE: func(cmd *cobra.Command, args []string) error {
			return tui.Run()
		},
	}
	rootCmd.SilenceUsage = true

	if err := rootCmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
