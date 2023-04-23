package gpt

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
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
		apiEndpoint: "https://api.openai.com/v1/chat/completions",
	}
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// GenerateText generates text using the GPT model.
func (c *Client) GenerateText(messages []Message) (string, error) {
	// リクエストを作成する
	requestBody, err := json.Marshal(map[string]interface{}{
		"model":    "gpt-3.5-turbo",
		"messages": messages,
	})
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", c.apiEndpoint, bytes.NewBuffer(requestBody))
	if err != nil {
		return "", err
	}

	// リクエストにAPIキーを設定する
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return "", err
	}

	if data["choices"] == nil {
		return "No response from OpenAI API", nil
	}

	choices := data["choices"].([]interface{})
	return choices[0].(map[string]interface{})["message"].(map[string]interface{})["content"].(string), nil
}
