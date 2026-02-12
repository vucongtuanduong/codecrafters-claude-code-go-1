package main

import "context"

type Tool interface {
	GetName() string
	GetDescription() string
	GetParameters() interface{}
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
