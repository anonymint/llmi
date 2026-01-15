# llmi

`llmi` is a Zsh-integrated CLI tool that uses Google Gemini to translate natural language into shell commands.

## Installation

1. **Build the binary:**
   ```bash
   make build
   ```

2. **Configure your API Key:**
   Add your Gemini API key to `~/.llmi` or export it as `GEMINI_API_KEY`:
   ```bash
   echo "GEMINI_API_KEY=your_key_here" >> ~/.llmi
   ```

3. **Source the Zsh widget:**
   Add this to your `~/.zshrc`:
   ```bash
   source /Users/mint/workspace/llmi/scripts/llmi.zsh
   ```

## Usage

1. Type `llmi` followed by your query:
   ```bash
   llmi read the last 5 lines of access.log
   ```
2. Hit `Enter`. You'll see the suggested command in dimmed text (ghost text).
3. Hit `Ctrl-G` to expand the suggestion into your main buffer.
4. Edit the command if needed, and hit `Enter` again to execute it.

## Configuration (`~/.llmi`)

- `GEMINI_API_KEY`: Your Google Gemini API key.
- `MODEL`: The model to use (default: `gemini-1.5-flash`).
- `HIST_COMMANDS`: Number of history items to include in context (default: `100`).
- `CUSTOM_RULES_PATH`: Path to a markdown file with custom shell rules.
