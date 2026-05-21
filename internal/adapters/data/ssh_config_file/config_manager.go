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
