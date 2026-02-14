package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/codecrafters-io/claude-code-starter-go/app/constant"
)

type ReadToolArguments struct {
	FilePath string `json:"file_path"`
}

var ReadToolParamConstant = ReadToolArguments{
	FilePath: "file_path",
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
			Name:        constant.ReadToolName,
			Description: constant.ReadToolDescription,
			Params: map[string]interface{}{
				constant.OpenaiParamKeyType: constant.OpenaiParamValueObjectType,
				constant.OpenaiParamKeyProperties: map[string]any{
					ReadToolParamConstant.FilePath: map[string]any{
						constant.OpenaiParamKeyType:        constant.OpenaiParamValueStringType,
						constant.OpenaiParamKeyDescription: constant.ReadToolFilePathParamDescription,
					},
				},
				constant.OpenaiParamKeyRequired: []string{ReadToolParamConstant.FilePath},
			},
		},
	}
}
