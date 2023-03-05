package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/glamour"
	"github.com/common-nighthawk/go-figure"
	gpt3 "github.com/sashabaranov/go-gpt3"
)

func main() {
	// 获取 OpenAI API Key
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		fmt.Println("请设置 OPENAI_API_KEY 环境变量")
		return
	}

	// 初始化 Glamour 渲染器
	renderStyle := glamour.WithEnvironmentConfig()
	mdRenderer, err := glamour.NewTermRenderer(
		renderStyle,
	)
	if err != nil {
		fmt.Println("初始化 Markdown 渲染器失败")
		return
	}

	// 输出欢迎语
	myFigure := figure.NewFigure("ChatGPT", "", true)
	myFigure.Print()
	fmt.Println()

	// 创建 ChatGPT 客户端
	client := gpt3.NewClient(apiKey)
	if err != nil {
		fmt.Printf("创建客户端失败: %s\n", err.Error())
		return
	}

	messages := []gpt3.ChatCompletionMessage{
		{
			Role:    "system",
			Content: "你是ChatGPT, OpenAI训练的大型语言模型, 请尽可能简洁地回答我的问题",
		},
	}

	for {
		fmt.Print("👽 ")

		// 读取用户输入并交互
		inputReader := bufio.NewReader(os.Stdin)
		userInput, err := inputReader.ReadString('\n')

		if err != nil {
			fmt.Println(err)
			continue
		}

		if userInput == "" || userInput == "\n" {
			continue
		}

		if strings.HasSuffix(userInput, "\\c\n") {
			continue
		}

		messages = append(
			messages, gpt3.ChatCompletionMessage{
				Role:    "user",
				Content: userInput,
			},
		)

		// 调用 ChatGPT API 接口生成回答
		resp, err := client.CreateChatCompletion(
			context.Background(),
			gpt3.ChatCompletionRequest{
				Model:       gpt3.GPT3Dot5Turbo,
				Messages:    messages,
				MaxTokens:   1024,
				Temperature: 0,
				N:           1,
			},
		)
		if err != nil {
			fmt.Printf("ChatGPT 接口调用失败: %s\n", err.Error())
			userInput = ""
			continue
		}

		// 格式化输出结果
		output := resp.Choices[0].Message.Content
		mdOutput, err := mdRenderer.Render(output)
		if err != nil {
			fmt.Printf("Markdown 渲染失败: %s\n", err.Error())
			userInput = ""
			continue
		}
		fmt.Println("🤖 " + mdOutput)
		messages = append(
			messages, gpt3.ChatCompletionMessage{
				Role:    "assistant",
				Content: output,
			},
		)
	}
}
