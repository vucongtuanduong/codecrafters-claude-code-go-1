package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/codecrafters-io/claude-code-starter-go/app/constant"
)

type WriteToolArguments struct {
	FilePath string `json:"file_path"`
	Content  string `json:"content"`
}
type WriteTool struct {
	BaseTool
}

func NewWriteTool() *WriteTool {
	return &WriteTool{
		BaseTool: BaseTool{
			Name:        constant.WriteToolName,
			Description: constant.WriteToolDescription,
			Params: map[string]interface{}{
				constant.OpenaiParamKeyType:     constant.OpenaiParamValueObject,
				constant.OpenaiParamKeyRequired: []string{constant.OpenaiParamKeyProperties.Filepath, constant.OpenaiParamKeyProperties.Content},
				constant.OpenaiParamKeyProperties.PropertyName: map[string]any{
					constant.OpenaiParamKeyProperties.Filepath: map[string]any{
						constant.OpenaiParamKeyType:        constant.OpenaiParamValueStringType,
						constant.OpenaiParamKeyDescription: constant.WriteToolFilePathParamDescription,
					},
					constant.OpenaiParamKeyProperties.Content: map[string]any{
						constant.OpenaiParamKeyType:        constant.OpenaiParamValueStringType,
						constant.OpenaiParamKeyDescription: constant.WriteToolContentParamDescription,
					},
				},
			},
		},
	}
}
func (w *WriteTool) Execute(ctx context.Context, args map[string]any) (string, error) {
	var writeArgs WriteToolArguments
	argsJson, err := json.Marshal(args)
	if err != nil {
		return "", fmt.Errorf("failed to marshal arguments: %w", err)
	}
	if err := json.Unmarshal(argsJson, &writeArgs); err != nil {
		return "", fmt.Errorf("failed to parse arguments: %w", err)
	}
	//create parent directories if they don't exists
	dir := filepath.Dir(writeArgs.FilePath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return "", fmt.Errorf("failed to create directory: %w", err)
	}
	err = os.WriteFile(writeArgs.FilePath, []byte(writeArgs.Content), 0644)
	if err != nil {
		return "", fmt.Errorf("failed to write file: %w", err)
	}
	return fmt.Sprintf("Successfully wrote to %s", writeArgs.FilePath), nil
}
