# llmi Zsh Widget

# Configuration
LLMI_BIN="/Users/mint/workspace/llmi/llmi"
LLMI_GHOST_COLOR=$'\e[38;5;242m' # Dim grey
LLMI_RESET=$'\e[0m'

# Fetch prefix from binary
LLMI_PREFIX=$($LLMI_BIN --print-prefix 2>/dev/null || echo ";;")

_llmi_suggested_cmd=""

llmi-widget() {
    # Check if the prefix exists anywhere in the buffer
    if [[ "$BUFFER" == *"$LLMI_PREFIX"* ]]; then
        local prefix_index="${BUFFER%%"$LLMI_PREFIX"*}"
        local pre_context="$prefix_index"
        local query="${BUFFER#*"$LLMI_PREFIX"}"
        query="${query# }" # trim leading space

        # Call the binary
        local suggested_cmd=$($LLMI_BIN --query "$query" --pre-context "$pre_context" 2>/dev/null)

        if [[ -n "$suggested_cmd" ]]; then
            # REPLACE the buffer with the (possibly locked) suggested command
            BUFFER="$pre_context$suggested_cmd"
            
            # If it's locked, move cursor to the start of the lock for easy deletion
            if [[ "$suggested_cmd" == SAFETY_LOCK_* ]]; then
                CURSOR=$#pre_context
            else
                CURSOR=$#BUFFER
            fi
            
            _llmi_suggested_cmd=""
            zle -R
        fi
    else
        # Normal behavior: No messages, just accept line
        _llmi_suggested_cmd=""
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
        # Clear any remaining message area
        zle -M ""
        zle -R
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
