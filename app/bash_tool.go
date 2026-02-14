package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os/exec"

	"github.com/codecrafters-io/claude-code-starter-go/app/constant"
)

type BashToolArguments struct {
	Command string `json:"command"`
}

var BashToolParamConstant = BashToolArguments{
	Command: "command",
}

type BashTool struct {
	BaseTool
}

func NewBashTool() *BashTool {
	return &BashTool{
		BaseTool: BaseTool{
			Name:        constant.BashToolName,
			Description: constant.BashToolDescription,
			Params: map[string]interface{}{
				constant.OpenaiParamKeyType:     constant.OpenaiParamValueObjectType,
				constant.OpenaiParamKeyRequired: []string{BashToolParamConstant.Command},
				constant.OpenaiParamKeyProperties: map[string]any{
					BashToolParamConstant.Command: map[string]any{
						constant.OpenaiParamKeyType:        constant.OpenaiParamValueStringType,
						constant.OpenaiParamKeyDescription: constant.BashToolCommandParamDescription,
					},
				},
			},
		},
	}
}
func (r *BashTool) Execute(ctx context.Context, args map[string]any) (string, error) {
	var bashArgs BashToolArguments
	argsJSON, err := json.Marshal(args)
	if err != nil {
		return "", fmt.Errorf("failed to marshal arguments: %w", err)
	}
	if err := json.Unmarshal(argsJSON, &bashArgs); err != nil {
		return "", fmt.Errorf("failed to parse arguments: %w", err)
	}

	command := bashArgs.Command
	cmd := exec.Command("sh", "-c", command)
	out, _ := cmd.CombinedOutput()
	return string(out), nil
}
