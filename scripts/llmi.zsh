# llmi Zsh Widget

# Configuration
LLMI_BIN="/Users/mint/workspace/llmi/llmi"
LLMI_GHOST_COLOR=$'\e[38;5;242m' # Dim grey
LLMI_RESET=$'\e[0m'

_llmi_suggested_cmd=""

llmi-widget() {
    # Only trigger if the command line starts with "llmi "
    if [[ "$BUFFER" == llmi\ * ]]; then
        local query="${BUFFER#llmi }"
        
        # Show "thinking" indicator
        POSTDISPLAY=" ${LLMI_GHOST_COLOR}... thinking${LLMI_RESET}"
        zle redisplay

        # Call the binary and get the suggested command
        _llmi_suggested_cmd=$($LLMI_BIN --query "$query" 2>/dev/null)

        if [[ -n "$_llmi_suggested_cmd" ]]; then
            # Display the suggestion as ghost text
            POSTDISPLAY=" ${LLMI_GHOST_COLOR}# → $_llmi_suggested_cmd${LLMI_RESET}"
        else
            POSTDISPLAY=" ${LLMI_GHOST_COLOR}# No suggestion found${LLMI_RESET}"
        fi
        zle redisplay
    else
        # If not llmi, just behave like a normal Enter key
        zle accept-line
    fi
}

llmi-expand() {
    if [[ -n "$_llmi_suggested_cmd" ]]; then
        # Replace the buffer with the suggested command
        BUFFER="$_llmi_suggested_cmd"
        CURSOR=$#BUFFER
        POSTDISPLAY=""
        _llmi_suggested_cmd=""
        zle redisplay
    fi
}

# Define the widgets
zle -N llmi-widget
zle -N llmi-expand

# Bind Enter to our widget
# Note: This might conflict with other plugins, but it's the most direct way
bindkey '^M' llmi-widget

# Bind Tab or Ctrl-G to expand
bindkey '^G' llmi-expand
