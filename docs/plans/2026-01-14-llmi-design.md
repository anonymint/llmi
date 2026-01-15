# llmi Design Specification

## Overview
`llmi` is a CLI tool that translates natural language prompts into executable shell commands directly in the Zsh environment. It uses Google's Gemini models (like Flash) for fast, context-aware command generation.

## Architecture

### 1. `llmi` Binary (Go)
- **Context Engine:**
    - Reads `~/.zsh_history` (last N commands).
    - Captures current aliases and functions.
    - Reads `~/.llmi` for configuration.
    - Reads optional custom markdown rules from a user-specified path.
- **LLM Client:**
    - Uses Google Gemini SDK.
    - Targets `gemini-1.5-flash` by default.
    - Structured prompting to return *only* the shell command.
- **Config (`~/.llmi`):**
    ```env
    GEMINI_API_KEY=xxx
    MODEL=gemini-1.5-flash
    HIST_COMMANDS=100
    CUSTOM_RULES_PATH=~/.llmi_rules.md
    ```

### 2. Zsh Integration (ZLE Widget)
- **Hook:** Intercepts `Enter` when the command line starts with `llmi`.
- **Ghost Text:** Uses `POSTDISPLAY` to show the suggested command in a dimmed color.
- **Expansion:** `Tab` or a specific key replaces the NLP query with the actual command for final editing and execution.

## Data Flow
1. User types `llmi <prompt>` and hits `Enter`.
2. ZLE widget calls `llmi --query "<prompt>"`.
3. `llmi` binary gathers history, aliases, and config.
4. `llmi` binary sends prompt to Gemini API.
5. Gemini API returns the shell command.
6. `llmi` binary outputs the command to the widget.
7. Widget displays the command as ghost text.
8. User hits `Tab` to accept/edit or `Enter` to refine the prompt.

## Error Handling
- **API Failure:** Show a brief error message in the ghost text area.
- **No Context:** Fall back to generic command generation if history/aliases cannot be read.
- **Timeout:** 5-second timeout for API calls to keep the shell responsive.

## Testing Strategy
- **Unit Tests (Go):** Test history parsing, config loading, and prompt construction.
- **Integration Tests:** Mock the Gemini API to verify the binary's output.
- **Manual Verification:** Test the Zsh widget in a controlled shell session.
