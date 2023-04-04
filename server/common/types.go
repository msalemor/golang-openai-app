package common

type AppRequest struct {
	Prompt string `json:"prompt"`
}

type AppResponse struct {
	Prompt     string `json:"prompt"`
	Completion string `json:"completion"`
}

type RequestBody struct {
	Prompt      string  `json:"prompt"`
	MaxTokens   int     `json:"max_tokens"`
	N           int     `json:"n"`
	Temperature float64 `json:"temperature"`
}

type TextCompletion struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Created int    `json:"created"`
	Model   string `json:"model"`
	Choices []struct {
		Text         string      `json:"text"`
		Index        int         `json:"index"`
		FinishReason string      `json:"finish_reason"`
		Logprobs     interface{} `json:"logprobs"`
	} `json:"choices"`
	Usage struct {
		CompletionTokens int `json:"completion_tokens"`
		PromptTokens     int `json:"prompt_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
}

type App struct {
	APIKey      string
	Endpoint    string
	Temperature float64
	N           int
	MaxTokens   int
	Stream      bool
}
