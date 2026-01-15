package context

import (
	"bufio"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"unicode/utf8"
)

type Context struct {
	History []string
	Aliases []string
	CustomRules string
}

func sanitize(s string) string {
	if utf8.ValidString(s) {
		return s
	}
	v := make([]rune, 0, len(s))
	for i, r := range s {
		if r == utf8.RuneError {
			_, size := utf8.DecodeRuneInString(s[i:])
			if size == 1 {
				continue
			}
		}
		v = append(v, r)
	}
	return string(v)
}

func GetContext(histCount int, customRulesPath string) *Context {
	ctx := &Context{}
	ctx.History = getHistory(histCount)
	ctx.Aliases = getAliases()
	
	// Sanitize everything
	for i, h := range ctx.History {
		ctx.History[i] = sanitize(h)
	}
	for i, a := range ctx.Aliases {
		ctx.Aliases[i] = sanitize(a)
	}

	if customRulesPath != "" {
		ctx.CustomRules = sanitize(getCustomRules(customRulesPath))
	}
	return ctx
}

func getHistory(count int) []string {
	home, _ := os.UserHomeDir()
	// Common zsh history file location
	histPath := filepath.Join(home, ".zsh_history")
	
	file, err := os.Open(histPath)
	if err != nil {
		return nil
	}
	defer file.Close()

	var history []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// Zsh history often has metadata like ": 1700000000:0;command"
		if strings.Contains(line, ";") {
			parts := strings.SplitN(line, ";", 2)
			if len(parts) > 1 {
				line = parts[1]
			}
		}
		history = append(history, line)
	}

	if len(history) > count {
		history = history[len(history)-count:]
	}
	return history
}

func getAliases() []string {
	// We run 'zsh -c alias' to get the current aliases
	cmd := exec.Command("zsh", "-c", "alias")
	output, err := cmd.Output()
	if err != nil {
		return nil
	}
	return strings.Split(string(output), "\n")
}

func getCustomRules(path string) string {
	// Expand ~ if present
	if strings.HasPrefix(path, "~") {
		home, _ := os.UserHomeDir()
		path = filepath.Join(home, path[1:])
	}
	data, _ := os.ReadFile(path)
	return string(data)
}
