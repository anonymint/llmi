package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/user/llmi/internal/config"
	"github.com/user/llmi/internal/context"
	"github.com/user/llmi/internal/llm"
)

func main() {
	queryFlag := flag.String("query", "", "The natural language query to translate")
	preContextFlag := flag.String("pre-context", "", "Text preceding the query (e.g., a pipe)")
	printPrefix := flag.Bool("print-prefix", false, "Print the configured trigger prefix")
	flag.Parse()

	cfg := config.LoadConfig()

	if *printPrefix {
		fmt.Print(cfg.TriggerPrefix)
		return
	}

	if *queryFlag == "" {
		// If no flag, check args
		if flag.NArg() > 0 {
			*queryFlag = flag.Arg(0)
		} else {
			fmt.Fprintf(os.Stderr, "Error: No query provided. Use --query \"...\" or llmi \"...\"\n")
			os.Exit(1)
		}
	}

	if cfg.GeminiAPIKey == "" {
		fmt.Fprintf(os.Stderr, "Error: GEMINI_API_KEY not found in env or ~/.llmi\n")
		os.Exit(1)
	}

	ctx := context.GetContext(cfg.HistCommands, cfg.CustomRulesPath)

	client, err := llm.NewClient(cfg.GeminiAPIKey, cfg.Model)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error initializing Gemini client: %v\n", err)
		os.Exit(1)
	}

	// Combine query and pre-context for better LLM understanding
	fullQuery := *queryFlag
	if *preContextFlag != "" {
		fullQuery = fmt.Sprintf("CONTEXT (Piping from or following): %s\n\nUSER QUERY: %s", *preContextFlag, *queryFlag)
	}

	command, err := client.GenerateCommand(fullQuery, ctx.History, ctx.Aliases, ctx.CustomRules)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error generating command: %v\n", err)
		os.Exit(1)
	}

	// Simple guardrail: detect destructive keywords
	destructive := []string{"rm -rf", "rm -r", "mkfs", "dd if=", "> /dev/sda", "shutdown", "reboot", "nuke"}
	isDestructive := false
	for _, word := range destructive {
		if strings.Contains(strings.ToLower(command), word) {
			isDestructive = true
			break
		}
	}

	if isDestructive {
		// Prepend a "Sabotage Prefix" that makes the command invalid until manually edited
		fmt.Printf("SAFETY_LOCK_REMOVE_THIS_PREFIX_TO_RUN %s", command)
	} else {
		fmt.Print(command)
	}
}
