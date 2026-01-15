# llmi Implementation Plan

> **Note:** Use `executing-plans` skill to implement this plan task-by-task.

**Goal:** Build a CLI tool `llmi` that translates natural language prompts into executable shell commands directly in the Zsh environment with ghost text support.

**Architecture:** A Go binary handles context gathering (history, aliases, config) and Gemini API calls, while a Zsh ZLE widget provides the interactive "ghost text" UI and expansion logic.

**Tech Stack:** Go 1.21+, Google Gemini Go SDK, Zsh.

---

### Task 1: Project Initialization

**Files:**
- Create: `go.mod`
- Create: `main.go`
- Create: `Makefile`

**Step 1: Initialize Go module**
```bash
go mod init github.com/user/llmi
```
**Step 2: Install Gemini SDK**
```bash
go get github.com/google/generative-ai-go/genai
go get google.golang.org/api/option
```
**Step 3: Create basic main.go**
```go
package main
import "fmt"
func main() { fmt.Println("llmi initialized") }
```
**Step 4: Create Makefile for easy building**
```makefile
build:
	go build -o llmi main.go
```
**Step 5: Verify build**
```bash
make build && ./llmi
```
**Step 6: Commit**
```bash
git add go.mod go.sum main.go Makefile && git commit -m "chore: initialize go project"
```

### Task 2: Configuration & Context Engine

**Files:**
- Create: `internal/config/config.go`
- Create: `internal/context/context.go`
- Test: `internal/context/context_test.go`

**Step 1: Implement Config loader** (Read `~/.llmi` and env vars)
**Step 2: Implement History parser** (Read `~/.zsh_history`)
**Step 3: Implement Alias/Function scraper** (Run `alias` command)
**Step 4: Write tests for context gathering**
**Step 5: Commit**
```bash
git add internal/ && git commit -m "feat: add config and context gathering"
```

### Task 3: Gemini API Integration

**Files:**
- Create: `internal/llm/gemini.go`

**Step 1: Implement Gemini client** (Initialize with API key and model)
**Step 2: Construct structured prompt** (Inject history, aliases, and user query)
**Step 3: Implement command extraction** (Ensure only the command is returned)
**Step 4: Commit**
```bash
git add internal/llm/ && git commit -m "feat: implement gemini api integration"
```

### Task 4: CLI Command Implementation

**Files:**
- Modify: `main.go`

**Step 1: Add flag parsing** (e.g., `--query`, `--config`)
**Step 2: Connect context engine and LLM client**
**Step 3: Output the raw command to stdout**
**Step 4: Verify with a real NLP query**
**Step 5: Commit**
```bash
git add main.go && git commit -m "feat: complete cli implementation"
```

### Task 5: Zsh Widget (The "Ghost Text" UI)

**Files:**
- Create: `scripts/llmi.zsh`

**Step 1: Define ZLE widget** `llmi-widget`
**Step 2: Implement logic to call `llmi` binary on Enter**
**Step 3: Implement `POSTDISPLAY` for ghost text**
**Step 4: Add keybinding for Tab/Expansion**
**Step 5: Commit**
```bash
git add scripts/ && git commit -m "feat: add zsh widget integration"
```

### Task 6: Final Verification & Installation

**Step 1: Source the zsh script**
**Step 2: Test the full flow: NLP -> Ghost Text -> Expansion -> Execution**
**Step 3: Document installation in README**
**Step 4: Commit**
```bash
git add README.md && git commit -m "docs: finalize installation instructions"
```
