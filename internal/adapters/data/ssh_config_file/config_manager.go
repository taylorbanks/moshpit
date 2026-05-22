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

package ssh_config_file

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"go.uber.org/zap"
)

// AppConfig holds application-level configuration persisted to disk.
type AppConfig struct {
	Theme       string `json:"theme"`
	GroupedView *bool  `json:"grouped_view,omitempty"`
	ShowSplash  *bool  `json:"show_splash,omitempty"`
}

type configManager struct {
	filePath string
	logger   *zap.SugaredLogger
}

// NewConfigManager creates a new config manager for the given file path.
func NewConfigManager(filePath string, logger *zap.SugaredLogger) *configManager {
	return &configManager{filePath: filePath, logger: logger}
}

// Load reads the config from disk. Returns a zero-value AppConfig if the file doesn't exist.
func (cm *configManager) Load() AppConfig {
	var cfg AppConfig

	if _, err := os.Stat(cm.filePath); os.IsNotExist(err) {
		return cfg
	}

	data, err := os.ReadFile(cm.filePath)
	if err != nil {
		cm.logger.Warnw("failed to read config file", "path", cm.filePath, "error", err)
		return cfg
	}

	if len(data) == 0 {
		return cfg
	}

	if err := json.Unmarshal(data, &cfg); err != nil {
		cm.logger.Warnw("failed to parse config file", "path", cm.filePath, "error", err)
		return cfg
	}

	return cfg
}

// Save writes the config to disk.
func (cm *configManager) Save(cfg AppConfig) error {
	dir := filepath.Dir(cm.filePath)
	if err := os.MkdirAll(dir, 0o750); err != nil {
		return fmt.Errorf("ensure config directory '%s': %w", dir, err)
	}

	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return fmt.Errorf("marshal config: %w", err)
	}

	if err := os.WriteFile(cm.filePath, data, 0o600); err != nil {
		return fmt.Errorf("write config '%s': %w", cm.filePath, err)
	}
	return nil
}
