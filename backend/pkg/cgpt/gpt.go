package cgpt

import (
	"context"
	"errors"
	"fmt"
	"github.com/kr/pretty"
	"github.com/sashabaranov/go-openai"
	"io"
	"net/http"
	"net/url"
	"sync"
)

/*
Model endpoint compatibility
ENDPOINT					MODEL NAME
/v1/chat/completions		gpt-4, gpt-4-0314, gpt-4-32k, gpt-4-32k-0314, gpt-3.5-turbo, gpt-3.5-turbo-0301
/v1/completions				text-davinci-003, text-davinci-002, text-curie-001, text-babbage-001, text-ada-001, davinci, curie, babbage, ada
/v1/edits					text-davinci-edit-001, code-davinci-edit-001
/v1/audio/transcriptions	whisper-1
/v1/audio/translations		whisper-1
/v1/fine-tunes				davinci, curie, babbage, ada
/v1/embeddings				text-embedding-ada-002, text-search-ada-doc-001
/v1/moderations				text-moderation-stable, text-moderation-latest
*/
var (
	once sync.Once
	gpt  *GPT
)

type GPT struct {
	Ctx                   context.Context
	Client                *openai.Client
	ChatCompletionRequest openai.ChatCompletionRequest
	CompletionRequest     openai.CompletionRequest
	Role                  string // role
	Token                 string
	Prompt                string // prompt
	Model                 string
	Suffix                string   // suffix 插入文本完成后出现的后缀。 Todo
	Temperature           int      // 介于 0 和 2 之间。较高的值（如 0.8）将使输出更加随机，而较低的值（如 0.2）将使输出更加集中和确定。建议更改这个值或者top_p，两者不能同时改。
	TopP                  int      // 核采样，其中模型考虑具有 top_p 概率质量的标记的结果。所以 0.1 意味着只考虑构成前 10% 概率质量的标记。
	N                     int      // 为每个提示生成多少完成。 因为这个参数会产生很多完成，它会很快消耗你的令牌配额。请谨慎使用并确保您对 max_tokens 和 stop 进行了合理的设置。
	Stream                bool     // 是否回流部分进度。如果设置，令牌将在可用时作为纯数据服务器发送事件发送，流由 data: [DONE] 消息终止。
	LogProbs              int      // 包括 logprobs 最有可能标记的对数概率，以及所选标记。例如，如果 logprobs 为 5，则 API 将返回 5 个最有可能的标记的列表。 API 将始终返回采样令牌的 logprob ，因此响应中最多可能有 logprobs+1 元素。logprobs 的最大值为 5
	Stop                  []string // API 将停止生成更多令牌的最多 4 个序列。返回的文本将不包含停止序列。
}

func New(token, model, prompt string, args ...interface{}) *GPT {
	once.Do(func() {
		role := openai.ChatMessageRoleUser
		if len(args) == 1 {
			role = args[0].(string)
		}
		gpt = &GPT{
			Ctx:    context.Background(),
			Client: openai.NewClient(token),
			Token:  token,
			Model:  model,
			Role:   role,
			Prompt: prompt,
		}
	})
	return gpt
}

// WithExtra
// @Description: with extra parameters
// @receiver g
// @param temp
// @param args
// @return *GPT
// @author cx
func (g *GPT) WithExtra(temp int, args ...interface{}) *GPT {
	g.Temperature = temp
	if len(args) == 1 {
		g.TopP = args[0].(int)
	}
	if len(args) == 2 {
		g.TopP = args[0].(int)
		g.N = args[1].(int)
	}
	// Todo
	return g
}

// WithProxy
// @Description: with proxy
// @receiver g
// @param rawUrl http://localhost:{port}
// @return *GPT
// @author cx
func (g *GPT) WithProxy(rawUrl string) *GPT {
	config := openai.DefaultConfig(g.Token)
	proxyUrl, err := url.Parse(rawUrl)
	if err != nil {
		panic(err)
	}
	transport := &http.Transport{
		Proxy: http.ProxyURL(proxyUrl),
	}
	config.HTTPClient = &http.Client{
		Transport: transport,
	}
	g.Client = openai.NewClientWithConfig(config)
	return g
}

func (g *GPT) WithChatCompletionRequest() *GPT {
	g.ChatCompletionRequest = openai.ChatCompletionRequest{
		Model: g.Model,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    g.Role,
				Content: g.Prompt,
			},
		},
	}
	return g
}

// ChatCompletionNoStream
// @Description: chat completion no streaming https://api.openai.com/v1/chat/completions
// @receiver g
// @param prompt
// @author cx
func (g *GPT) ChatCompletionNoStream(prompt string) (*openai.ChatCompletionResponse, error) {
	req := openai.ChatCompletionRequest{
		Model: g.Model,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    g.Role,
				Content: prompt,
			},
		},
	}
	resp, err := g.Client.CreateChatCompletion(g.Ctx, req)
	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return nil, err
	}
	//fmt.Println(resp.Choices[0].Message.Content)
	return &resp, nil
}

// ChatCompletionStream
// @Description: ChatGPT streaming completion https://api.openai.com/v1/chat/completions
// @receiver g
// @param prompt
// @author cx
func (g *GPT) ChatCompletionStream(prompt string) {
	req := openai.ChatCompletionRequest{
		Model:     g.Model,
		MaxTokens: g.getMaxTokensByModel(g.Model),
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    g.Role,
				Content: prompt,
			},
		},
		Stream: true,
	}
	stream, err := g.Client.CreateChatCompletionStream(g.Ctx, req)
	if err != nil {
		fmt.Printf("ChatCompletionStream error: %v\n", err)
		return
	}
	defer stream.Close()
	fmt.Printf("Stream response: ")
	for {
		response, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			fmt.Println("\nStream finished")
			return
		}
		if err != nil {
			fmt.Printf("\nStream error: %v\n", err)
			return
		}
		fmt.Printf(response.Choices[0].Delta.Content)
	}
}

// CompletionStream
// @Description: gpt3 completion stream https://api.openai.com/v1/completions
// @receiver g
// @param prompt
// @author cx
func (g *GPT) CompletionStream(prompt string) {
	req := openai.CompletionRequest{
		Model:     openai.GPT3Ada,
		MaxTokens: g.getMaxTokensByModel(g.Model),
		Prompt:    prompt,
		Stream:    true,
	}
	stream, err := g.Client.CreateCompletionStream(g.Ctx, req)
	if err != nil {
		fmt.Printf("\nCompletionStream error: %v\n", err)
		return
	}
	defer stream.Close()

	for {
		response, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			fmt.Println("\nStream finished")
			return
		}

		if err != nil {
			fmt.Printf("\nStream error: %v\n", err)
			return
		}
		fmt.Printf(response.Choices[0].Text)
	}
}

// CompletionNoStream
// @Description: https://api.openai.com/v1/completions 为提供的提示和参数创建完成
// Request
//
//	{
//	 "model": "text-davinci-003",
//	 "prompt": "Say this is a test",
//	 "max_tokens": 7,
//	 "temperature": 0,
//	 "top_p": 1,
//	 "n": 1,
//	 "stream": false,
//	 "logprobs": null,
//	 "stop": "\n"
//	}
//
// @receiver g
// @param prompt
// @author cx
func (g *GPT) CompletionNoStream(prompt string) {
	req := openai.CompletionRequest{
		Model:     g.Model,
		MaxTokens: g.getMaxTokensByModel(g.Model),
		Prompt:    prompt,
	}
	resp, err := g.Client.CreateCompletion(g.Ctx, req)
	if err != nil {
		fmt.Printf("Completion error: %v\n", err)
		return
	}
	pretty.Println(resp)
}

// ListModels
// @Description: get models
// @receiver g
// @return openai.ModelsList
// @return error
// @author cx
func (g *GPT) ListModels() (openai.ModelsList, error) {
	modelList, err := g.Client.ListModels(g.Ctx)
	if err != nil {
		return modelList, err
	}
	return modelList, nil
}

// getMaxTokensByModel
// @Description: https://platform.openai.com/docs/models/gpt-3-5
// @receiver g
// @param model
// @return int
// @author cx
func (g *GPT) getMaxTokensByModel(model string) int {
	tokens := map[string]int{
		openai.GPT4:                8192,  // 比任何 GPT-3.5 模型都更强大，能够执行更复杂的任务，并针对聊天进行了优化。将使用我们最新的模型迭代进行更新。截至 2021 年 9 月
		openai.GPT40314:            8192,  // gpt-4 2023 年 3 月 14 日的快照。与 gpt-4 不同，此模型不会收到更新，并且仅在 2023 年 6 月 14 日结束的三个月内提供支持。截至 2021 年 9 月
		openai.GPT432K:             32768, // 与基本 gpt-4 模式相同的功能，但上下文长度是其 4 倍。将使用我们最新的模型迭代进行更新。截至 2021 年 9 月
		openai.GPT432K0314:         32768, // gpt-4-32 2023 年 3 月 14 日的快照。与 gpt-4-32k 不同，此模型不会收到更新，并且仅在 2023 年 6 月 14 日结束的三个月内提供支持。截至 2021 年 9 月
		openai.GPT3Dot5Turbo:       4096,  // 功能最强大的 GPT-3.5 模型并针对聊天进行了优化，成本仅为 text-davinci-003 的 1/10。将使用我们最新的模型迭代进行更新。截至 2021 年 9 月
		openai.GPT3Dot5Turbo0301:   4096,  // gpt-3.5-turbo 的快照，2023 年 3 月 1 日。与 gpt-3.5-turbo 不同，此模型不会收到更新，并且仅在 2023 年 6 月 1 日结束的三个月内提供支持。截至 2021 年 9 月
		openai.GPT3TextDavinci003:  4097,  // 可以以比居里、巴贝奇或 ada 模型更好的质量、更长的输出和一致的指令遵循来完成任何语言任务。还支持在文本中插入补全。 截至 2021 年 6 月
		openai.GPT3TextDavinci002:  4097,  // 与 text-davinci-003 类似的功能，但使用监督微调而不是强化学习进行训练 截至 2021 年 6 月
		openai.CodexCodeDavinci002: 8001,  // 针对代码完成任务进行了优化 截至 2021 年 6 月
		openai.GPT3TextCurie001:    2049,  // 非常有能力，比达芬奇更快，成本更低。截至 2019 年 10 月
		openai.GPT3TextBabbage001:  2049,  // 能够执行简单的任务，速度非常快，成本更低。截至 2019 年 10 月
		openai.GPT3TextAda001:      2049,  // 能够执行非常简单的任务，通常是 GPT-3 系列中最快的型号，而且成本最低。截至 2019 年 10 月
		openai.GPT3Davinci:         2049,  // 功能最强大的 GPT-3 模型。可以完成其他模型可以完成的任何任务，而且通常质量更高。截至 2019 年 10 月
		openai.GPT3Curie:           2049,  // 非常有能力，但比达芬奇更快，成本更低。截至 2019 年 10 月
		openai.GPT3Babbage:         2049,  // 能够执行简单的任务，速度非常快，成本更低。截至 2019 年 10 月
		openai.GPT3Ada:             2049,  // 能够执行非常简单的任务，通常是 GPT-3 系列中最快的型号，而且成本最低。截至 2019 年 10 月
	}
	return tokens[model]
}
