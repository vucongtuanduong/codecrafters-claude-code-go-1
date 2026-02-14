package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/codecrafters-io/claude-code-starter-go/app/constant"
	"github.com/joho/godotenv"
	"github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/option"
)

func main() {
	var prompt string
	flag.StringVar(&prompt, "p", "", "Prompt to send to LLM")
	flag.Parse()
	_ = godotenv.Load()
	var tools_box map[string]Tool
	tools_box = make(map[string]Tool)
	tools_box[constant.ReadToolName] = NewReadTool()
	tools_box[constant.WriteToolName] = NewWriteTool()
	tools_box[constant.BashToolName] = NewBashTool()
	if prompt == "" {
		panic("Prompt must not be empty")
	}

	apiKey := os.Getenv("OPENROUTER_API_KEY")
	baseUrl := os.Getenv("OPENROUTER_BASE_URL")
	model := os.Getenv("MODEL_NAME")
	if model == "" {
		model = "anthropic/claude-haiku-4.5"
	}
	if baseUrl == "" {
		baseUrl = "https://openrouter.ai/api/v1"
	}

	if apiKey == "" {
		panic("Env variable OPENROUTER_API_KEY not found")
	}

	client := openai.NewClient(option.WithAPIKey(apiKey), option.WithBaseURL(baseUrl))
	var messages []openai.ChatCompletionMessageParamUnion
	messages = append(messages, openai.ChatCompletionMessageParamUnion{
		OfUser: &openai.ChatCompletionUserMessageParam{
			Content: openai.ChatCompletionUserMessageParamContentUnion{
				OfString: openai.String(prompt),
			},
		},
	})
	var tools []openai.ChatCompletionToolUnionParam
	tools = append(tools, tools_box[constant.ReadToolName].GetDefinition())
	tools = append(tools, tools_box[constant.WriteToolName].GetDefinition())
	for {
		params := openai.ChatCompletionNewParams{
			Model:    model,
			Messages: messages,
			Tools:    tools,
		}
		response, err := client.Chat.Completions.New(context.Background(), params)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
		if len(response.Choices) == 0 {
			panic("No choices in response")
		}
		choice := response.Choices[0]
		messages = append(messages, choice.Message.ToParam())
		if response.Choices[0].Message.ToolCalls != nil {
			for _, toolCall := range response.Choices[0].Message.ToolCalls {
				var argsMap map[string]any
				err := json.Unmarshal([]byte(toolCall.Function.Arguments), &argsMap)
				if err != nil {
					fmt.Fprintf(os.Stderr, "error: %v\n", err)
					os.Exit(1)
				}
				result, _ := tools_box[toolCall.Function.Name].Execute(context.Background(), argsMap)
				messages = append(messages, openai.ToolMessage(result, toolCall.ID))
			}
		} else {
			fmt.Print(response.Choices[0].Message.Content)
			break
		}
	}

}
