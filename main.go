package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/user/llmi/internal/config"
	"github.com/user/llmi/internal/context"
	"github.com/user/llmi/internal/llm"
)

func main() {
	queryFlag := flag.String("query", "", "The natural language query to translate")
	flag.Parse()

	if *queryFlag == "" {
		// If no flag, check args
		if flag.NArg() > 0 {
			*queryFlag = flag.Arg(0)
		} else {
			fmt.Fprintf(os.Stderr, "Error: No query provided. Use --query \"...\" or llmi \"...\"\n")
			os.Exit(1)
		}
	}

	cfg := config.LoadConfig()
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

	command, err := client.GenerateCommand(*queryFlag, ctx.History, ctx.Aliases, ctx.CustomRules)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error generating command: %v\n", err)
		os.Exit(1)
	}

	fmt.Print(command)
}