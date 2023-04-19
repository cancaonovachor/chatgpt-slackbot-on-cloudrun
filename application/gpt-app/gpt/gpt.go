package gpt

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

// client is a client for accessing the OpenAI GPT API.
type Client struct {
	// APIキー
	apiKey string
	// APIエンドポイントのURL
	apiEndpoint string
}

// NewClient returns a new Client.
func NewClient() *Client {
	return &Client{
		apiKey:      os.Getenv("OPENAI_API_KEY"),
		apiEndpoint: "https://api.openai.com/v1/completions",
	}
}

type ChatRequest struct {
	Model            string  `json:"model"`
	Prompt           string  `json:"prompt"`
	Temperature      float64 `json:"temperature"`
	MaxTokens        int     `json:"max_tokens"`
	TopP             float64 `json:"top_p"`
	FrequencyPenalty float64 `json:"frequency_penalty"`
	PresencePenalty  float64 `json:"presence_penalty"`
}

type ChatResponse struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Created int    `json:"created"`
	Model   string `json:"model"`
	Choices []struct {
		Text         string      `json:"text"`
		Index        int         `json:"index"`
		Logprobs     interface{} `json:"logprobs"`
		FinishReason string      `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
}

// GenerateText generates text using the GPT model.
func (c *Client) GenerateText(prompt string, maxTokens int) (string, error) {
	// リクエストを作成する
	chatRequest := ChatRequest{
		Model:            "text-davinci-003",
		Prompt:           prompt,
		Temperature:      0.7,
		MaxTokens:        maxTokens,
		TopP:             1,
		FrequencyPenalty: 0.1,
		PresencePenalty:  0.2,
	}

	jsonData, _ := json.Marshal(&chatRequest)

	req, err := http.NewRequest("POST", c.apiEndpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}

	// リクエストにAPIキーを設定する
	req.Header.Add("Authorization", "Bearer "+c.apiKey)
	req.Header.Add("Content-Type", "application/json")

	// リクエストを送信する
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	log.Println("gpt response: ", resp)
	defer resp.Body.Close()

	r := &ChatResponse{}
	if err := json.NewDecoder(resp.Body).Decode(r); err != nil {
		return "", err
	}
	log.Println("gpt response: ", r.Choices)
	text := ""
	for _, v := range r.Choices {
		text += fmt.Sprintf("%s\n", v.Text)
	}

	return text, err
}
