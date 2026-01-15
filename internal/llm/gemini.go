package llm

import (
	"context"
	"fmt"
	"strings"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

type Client struct {
	client *genai.Client
	model  *genai.GenerativeModel
}

func NewClient(apiKey, modelName string) (*Client, error) {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return nil, err
	}
	model := client.GenerativeModel(modelName)
	return &Client{client: client, model: model}, nil
}

func (c *Client) GenerateCommand(userPrompt string, history, aliases []string, customRules string) (string, error) {
	ctx := context.Background()

	systemPrompt := "You are a shell command generator. Translate the user's natural language prompt into a single executable shell command. "
	systemPrompt += "Return ONLY the command, with no explanation, markdown formatting, or preamble. "
	
	if len(history) > 0 {
		systemPrompt += "\n\nRecent shell history for context:\n" + strings.Join(history, "\n")
	}
	
	if len(aliases) > 0 {
		systemPrompt += "\n\nCurrent shell aliases for context:\n" + strings.Join(aliases, "\n")
	}

	if customRules != "" {
		systemPrompt += "\n\nCustom user rules:\n" + customRules
	}

	resp, err := c.model.GenerateContent(ctx, genai.Text(systemPrompt), genai.Text("User prompt: "+userPrompt))
	if err != nil {
		return "", err
	}

	if len(resp.Candidates) == 0 || len(resp.Candidates[0].Content.Parts) == 0 {
		return "", fmt.Errorf("no response from Gemini")
	}

	cmd := fmt.Sprintf("%v", resp.Candidates[0].Content.Parts[0])
	return strings.TrimSpace(cmd), nil
}
