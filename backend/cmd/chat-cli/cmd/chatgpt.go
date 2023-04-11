/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"chatcat/backend/pkg/cgpt"
	"chatcat/backend/pkg/chelp"
	"fmt"
	"github.com/sashabaranov/go-openai"

	"github.com/spf13/cobra"
)

var (
	prompt string
	model  string
	token  string
)

// chatgptCmd represents the chatgpt command
var chatgptCmd = &cobra.Command{
	Use:   "chatgpt",
	Short: "chatgpt接口测试",
	Long:  `chatgpt gpt3,gpt4 接口测试`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("chatgpt called")
		handleChat()
	},
}

func init() {
	rootCmd.AddCommand(chatgptCmd)
	chatgptCmd.PersistentFlags().StringVar(&model, "m", openai.GPT3Dot5Turbo, "模型 gpt3.5 gpt4")
	chatgptCmd.PersistentFlags().StringVar(&token, "t", chelp.GetGPTToken(), "token值")
	chatgptCmd.Flags().StringVar(&prompt, "p", "", "提示词")
	chatgptCmd.MarkFlagRequired("p")
	chatgptCmd.MarkFlagRequired("m")
	chatgptCmd.MarkFlagRequired("t")
}

func handleChat() {
	cgpt.New(token, model).WithProxy("http://127.0.0.1:7890").ChatGPTCompletionStream(prompt)
	//cgpt.New(token, model).WithProxy("http://127.0.0.1:7890").GPT3CompletionStream(prompt)
	//cgpt.New(token, model).WithProxy("http://127.0.0.1:7890").NormalChatCompletion(prompt)
	// ada
	//cgpt.New(token, model).WithProxy("http://127.0.0.1:7890").GPT3CompletionNoStream(prompt)
	//cgpt.New(token, model).WithProxy("http://127.0.0.1:7890").GetModels()

}
