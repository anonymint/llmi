package context

import (
	"os"
	"path/filepath"
	"testing"
)

func TestGetHistory(t *testing.T) {
	// Mock history file
	tempDir, _ := os.MkdirTemp("", "llmi_test")
	defer os.RemoveAll(tempDir)
	
histPath := filepath.Join(tempDir, ".zsh_history")
	content := ": 1700000000:0;ls -la\n: 1700000001:0;cd /tmp\n: 1700000002:0;cat file.txt\n"
	os.WriteFile(histPath, []byte(content), 0644)

	// Update HOME for the test
	oldHome := os.Getenv("HOME")
	os.Setenv("HOME", tempDir)
	defer os.Setenv("HOME", oldHome)

	history := getHistory(2)
	if len(history) != 2 {
		t.Errorf("Expected 2 history items, got %d", len(history))
	}
	if history[0] != "cd /tmp" {
		t.Errorf("Expected 'cd /tmp', got '%s'", history[0])
	}
}

