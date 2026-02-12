package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
)

type ReadToolArguments struct {
	FilePath string `json:"file_path"`
}
type ReadTool struct {
	BaseTool
}

func (r *ReadTool) Execute(ctx context.Context, args map[string]any) (string, error) {
	var readArgs ReadToolArguments
	argsJSON, err := json.Marshal(args)
	if err != nil {
		return "", fmt.Errorf("failed to marshal arguments: %w", err)
	}
	if err := json.Unmarshal(argsJSON, &readArgs); err != nil {
		return "", fmt.Errorf("failed to parse arguments: %w", err)
	}

	content, err := os.ReadFile(readArgs.FilePath)
	if err != nil {
		return "", fmt.Errorf("failed to read file: %w", err)
	}

	return string(content), nil
}
func NewReadTool() *ReadTool {
	return &ReadTool{
		BaseTool: BaseTool{
			Name:        ReadToolName,
			Description: ReadToolDescription,
			Params: map[string]interface{}{
				"type": "object",
				"properties": map[string]any{
					"file_path": map[string]any{
						"type":        "string",
						"description": "The path to the file to read",
					},
				},
				"required": []string{"file_path"},
			},
		},
	}
}
