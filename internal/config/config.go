package config

import (
	"bufio"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type Config struct {
	GeminiAPIKey      string
	Model             string
	HistCommands      int
	CustomRulesPath   string
}

func LoadConfig() *Config {
	cfg := &Config{
		GeminiAPIKey: os.Getenv("GEMINI_API_KEY"),
		Model:        "gemini-1.5-flash",
		HistCommands: 100,
	}

	home, _ := os.UserHomeDir()
	llmiPath := filepath.Join(home, ".llmi")

	if file, err := os.Open(llmiPath); err == nil {
		defer file.Close()
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			if strings.Contains(line, "=") {
				parts := strings.SplitN(line, "=", 2)
				key := strings.TrimSpace(parts[0])
				val := strings.TrimSpace(parts[1])
				switch key {
				case "GEMINI_API_KEY":
					if cfg.GeminiAPIKey == "" {
						cfg.GeminiAPIKey = val
					}
				case "MODEL":
					cfg.Model = val
				case "HIST_COMMANDS":
					if n, err := strconv.Atoi(val); err == nil {
						cfg.HistCommands = n
					}
				case "CUSTOM_RULES_PATH":
					cfg.CustomRulesPath = val
				}
			}
		}
	}

	return cfg
}
