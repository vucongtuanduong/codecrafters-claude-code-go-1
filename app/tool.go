package main

import (
	"context"

	"github.com/openai/openai-go/v3"
)

type Tool interface {
	GetName() string
	GetDescription() string
	GetParameters() interface{}
	GetDefinition() openai.ChatCompletionToolUnionParam
	Execute(ctx context.Context, args map[string]any) (string, error)
}
type BaseTool struct {
	Name        string
	Description string
	Params      interface{}
}

func (b *BaseTool) GetName() string {
	return b.Name
}
func (b *BaseTool) GetDescription() string {
	return b.Description
}
func (b *BaseTool) GetParameters() interface{} {
	return b.Params
}
func (b *BaseTool) GetDefinition() openai.ChatCompletionToolUnionParam {
	return openai.ChatCompletionFunctionTool(openai.FunctionDefinitionParam{
		Name:        b.Name,
		Description: openai.String(b.Description),
		Parameters:  b.Params.(map[string]any),
	})
}
