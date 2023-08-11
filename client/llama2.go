package client

type Llama2TextGenerationRequest struct {
	Model string                    `json:"model"`
	Param Llama2TextGenerationParam `json:"data"`
}

type Llama2TextGenerationParam struct {
	Prompt            string  `json:"prompt"`
	TopK              int     `json:"top_k"`
	TopP              float64 `json:"top_p"`
	Temp              float64 `json:"temp"`
	MaxLength         int     `json:"max_length"`
	RepetitionPenalty float64 `json:"repetition_penalty"`
	BeamSize          int     `json:"beam_size"`
}
