package client

type TextGenerationRequest struct {
	Model string              `json:"model"`
	Param TextGenerationParam `json:"data"`
}

type TextGenerationParam struct {
	Prompt string  `json:"prompt"`
	TopK   int     `json:"top_k"`
	TopP   float64 `json:"top_p"`
}
