# llmi

`llmi` is a Zsh-integrated CLI tool that uses Google Gemini to translate natural language into shell commands with context awareness.

## Features

- **Context Aware:** Sends your recent shell history and aliases to the LLM to provide better suggestions.
- **Piping Support:** Understands context from previous commands (e.g., `cat file.txt | ;; process this`).
- **Instant Swap:** Automatically replaces your query with the generated command directly in your prompt.
- **Safety Lock:** Prepend a safety string to destructive commands (like `rm -rf`) to prevent accidental execution.
- **Configurable:** Change the trigger prefix, model, and history depth via `~/.llmi`.

## Installation

1. **Build the binary:**
   ```bash
   make build
   ```

2. **Configure your API Key:**
   Copy the example config and add your Gemini API key:
   ```bash
   cp .llmi.example ~/.llmi
   # Then edit ~/.llmi and add your key
   ```

3. **Source the Zsh widget:**
   Add this to your `~/.zshrc`:
   ```bash
   source /Users/mint/workspace/llmi/scripts/llmi.zsh
   ```

## Usage

1. Type your configured prefix (default `;;`) followed by your query:
   ```bash
   ;; read the last 5 lines of access.log
   ```
2. Hit `Enter`. The text will immediately swap to the suggested command.
3. If the command is dangerous, it will be locked:
   `SAFETY_LOCK_REMOVE_THIS_PREFIX_TO_RUN rm -rf *`
4. Edit the command if needed (or delete the safety lock), and hit `Enter` again to execute.

## Configuration (`~/.llmi`)

- `GEMINI_API_KEY`: Your Google Gemini API key.
- `MODEL`: The model to use (default: `models/gemini-2.5-flash-lite`).
- `HIST_COMMANDS`: Number of history items to include in context (default: `100`).
- `TRIGGER_PREFIX`: The prefix to trigger the LLM (default: `;;`).
- `CUSTOM_RULES_PATH`: Path to a markdown file with custom shell rules.